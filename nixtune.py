# Copyright (C) 2016 opszero <hey@opszero.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import string

def procfs_feature(func):
    """
    procfs_feature returns a dictionary if the current system doesn't match
    the value we want it to have.
    """

    def wrapper(self, *args, **kwargs):
        output = func(self, *args, **kwargs)

        returned = {}
        for k, v in output.items():
            change = {
                'Value': v,
                'Docs': string.strip(func.__doc__),
            }
            returned = dict(returned.items() + [(k, change)])

        return returned

    return wrapper

def sysfs_feature(func):
    def wrapper(self):
        return func(self)

    return wrapper


class MentalModel(object):
    """
    MentalModel provides utilities to handle feature extraction.
    """

    def procfs_features(self):
        procfs = []
        for k, f in self.__class__.__dict__.iteritems():
            if callable(f) and k.startswith('procfs_'):
                procfs += f(self).items()

        return dict(procfs)

    def sysfs_features(self):
        sysfs = []
        for k, f in self.__class__.__dict__.iteritems():
            if callable(f) and k.startswith('sysfs_'):
                sysfs += f(self).items()

        return dict(sysfs)



class Memory(MentalModel):
    @procfs_feature
    def procfs_vm_swappiness(self):
        """
        Disable swapping and clear the file system page cache to free memory first.
        """

        return {
            "/proc/sys/vm/swappiness": "0"
        }

    @procfs_feature
    def procfs_vm_min_free_kbytes(self):
        """
        Amount of memory to keep free. Don't want to make this too high as
        Linux will spend more time trying to reclaim memory.
        """

        # TODO: Check the amount of ram on the machine and adjust this
        # number appropriately. We mostly want the ability to SSH in
        # if things hit the fan.
        return {
            "/proc/sys/vm/min_free_kbytes": "65536"
        }

    @sysfs_feature
    def sysfs_mm_transparent_hugepages(self):
        """
        Explict huge page usage making the page size of 2 or 4 MB instead
        of 4kb. Should reduce CPU overhead and improve MMU page
        translation.
        """

        return {
            "/sys/kernel/mm/transparent_hugepage/enabled": "always"
        }


class Networking(MentalModel):
    """
    References:

    http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html
    https://rtcamp.com/tutorials/linux/sysctl-conf/
    https://fasterdata.es.net/host-tuning/linux/
    http://cherokee-project.com/doc/other_os_tuning.html
    https://easyengine.io/tutorials/linux/sysctl-conf/
    https://access.redhat.com/sites/default/files/attachments/20150325_network_performance_tuning.pdf
    """

    def __init__(self):
        self.vars = {
            'nfConntrackMax': 200000
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_fin_timeout(self):
        """
        Usually, the Linux kernel holds a TCP connection even after it
        is closed for around two minutes. This means that there may be
        a port exhaustion as the kernel waits to close the
        connections. By moving the fin_timeout to 15 seconds we
        drastically reduce the length of time the kernel is waiting
        for the socket to get any remaining packets.
        """

        return {
            "/proc/sys/net/ipv4/tcp_fin_timeout": "15"
        }

    @procfs_feature
    def procfs_net_ipv4_ip_local_port_range(self):
        """
        On a typical machine there are around 28,000 ports available to be
        bound to. This number can get exhausted quickly if there are many
        connections. We will increase this.
        """

        return {
            "/proc/sys/net/ipv4/ip_local_port_range": "1024 65535",
        }

    @procfs_feature
    def procfs_net_core_rmem_max(self):
        """
        The size of the receive buffer for all the sockets. 16MB per socket.
        """

        # TODO: Adjust this per the instance size.

        return {
            "/proc/sys/net/core/rmem_max": "16777216"
        }

    @procfs_feature
    def procfs_net_core_wmem_max(self):
        """
        The size of the buffer for all the sockets. 16MB per socket.
        """

        # TODO: Adjust this per the instance size.

        return {
            "/proc/sys/net/core/wmem_max": "16777216",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_rmem(self):
        """
        (min, default, max): The sizes of the receive buffer for the IP protocol.
        """


        # TODO: Adjust this per the instance size.

        return {
            "/proc/sys/net/ipv4/tcp_rmem": "4096 87380 16777216",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_wmem(self):
        """
        (min, default, max): The sizes of the write buffer for the IP protocol.
        """

        # TODO: Adjust this per the instance size.

        return {
            "/proc/sys/net/ipv4/tcp_wmem": "4096 65536 16777216",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_max_syn_backlog(self):
        """
        Increase the number syn requests allowed. Sets how many half-open connections to backlog queue
        """

        # TODO: Adjust this per the instance size.

        return {
            "/proc/sys/net/ipv4/tcp_max_syn_backlog": "20480",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_syncookies(self):
        """
        Security to prevent DDoS attacks. http://cr.yp.to/syncookies.html
        """

        return {
            "/proc/sys/net/ipv4/tcp_syncookies": "1",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_no_metrics_save(self):
        """
        TCP saves various connection metrics in the route cache when the
        connection closes so that connections established in the near future
        can use these to set initial conditions. Usually, this increases
        overall performance, but may sometimes cause performance
        degradation.
        """

        return {
            "/proc/sys/net/ipv4/tcp_no_metrics_save": "1",
        }

    @procfs_feature
    def procfs_net_core_somaxconn(self):
        """
        The maximum number of queued sockets on a connection.
        """

        return {
            "/proc/sys/net/core/somaxconn": "16096",
        }

    @procfs_feature
    def procfs_net_core_netdev_max_backlog(self):
        """
        The number of incoming connections on the backlog queue. The maximum
        number of packets queued on the INPUT side.
        """

        return {
            "/proc/sys/net/core/netdev_max_backlog": "30000",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_max_tw_buckets(self):
        """
        Increase the tcp-time-wait buckets pool size to prevent simple DOS attacks
        """

        return {
            "/proc/sys/net/ipv4/tcp_max_tw_buckets": "400000",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_syn_retries(self):
        """
        Number of times initial SYNs for a TCP connection attempt will
        be retransmitted for outgoing connections.
        """

        return {
            "/proc/sys/net/ipv4/tcp_syn_retries": "2",
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_synack_retries(self):
        """
        This setting determines the number of SYN+ACK packets sent before
        the kernel gives up on the connection
        """

        return {
            "/proc/sys/net/ipv4/tcp_synack_retries": "2",
        }

    @procfs_feature
    def procfs_net_netfilter_nf_conntrack_max(self):
        """
        The max is double the previous value.
        https://wiki.khnet.info/index.php/Conntrack_tuning
        """

        return {
            "/proc/sys/net/netfilter/nf_conntrack_max": self.vars['nfConntrackMax'],
        }

    @procfs_feature
    def procfs_net_ipv4_tcp_tw_reuse(self):
        """
        """

        return {
            "/proc/sys/net/ipv4/tcp_tw_reuse": "1",
        }

    @sysfs_feature
    def sysfs_nf_conntrack_hashsize(self):
        """
        """

        return {
            "/sys/module/nf_conntrack/parameters/hashsize": self.vars["nfConntrackMax"] / 4
        }

class IO(MentalModel):
    """
    References:
      - http://www.brendangregg.com/linuxperf.html
    """

    @sysfs_feature
    def sysfs_block_queue_rq_afinity(self):
        return {
            "/sys/block/*/queue/rq_afinity": "2"
        }

    @sysfs_feature
    def sysfs_block_queue_scheduler(self):
        return {
            "/sys/block/*/queue/scheduler": "noop"
        }

    @sysfs_feature
    def sysfs_block_queue_read_ahead_kb(self):
        return {
            "/sys/block/*/queue/read_ahead_kb": "256",
        }


# # func (f *FS) fstabNoattime() *ProfileKV {
# #       b, err := ioutil.ReadFile("/etc/fstab")
# #       if err != nil {
# #               return nil
# #       }

# #       content := string(b)

# #       var hasAttime []string

# #       for _, i := range strings.Split(content, "\n") {
# #               params := strings.Fields(i)
# #               if len(params) == 0 {
# #                       continue
# #               }

# #               // device = params[0]
# #               mountPoint := params[1]
# #               // fileSystem = params[2]
# #               attributes := params[3]

# #               if !strings.Contains(attributes, "noattime") {
# #                       hasAttime = append(hasAttime, mountPoint)
# #               }
# #       }

# #       var out bytes.Buffer
# #       if len(hasAttime) > 0 {
# #               fmt.Fprintf(&out, "\tMount the following mountpoints with attribute noattime.\n")

# #               for _, i := range hasAttime {
# #                       fmt.Fprintf(&out, "\t - %s\n", i)
# #               }
# #       }

# #       return &ProfileKV{
# #               Value:       out.String(),
# #               Description: "",
# #       }
# # }

# # func (f *FS) fstabDiscard() *ProfileKV {
# #       b, err := ioutil.ReadFile("/etc/fstab")
# #       if err != nil {
# #               return nil
# #       }

# #       content := string(b)

# #       var hasDiscard []string

# #       for _, i := range strings.Split(content, "\n") {
# #               params := strings.Fields(i)
# #               if len(params) == 0 {
# #                       continue
# #               }

# #               // device = params[0]
# #               mountPoint := params[1]
# #               // fileSystem = params[2]
# #               attributes := params[3]

# #               if f.isSSD(mountPoint) && strings.Contains(attributes, "discard") {
# #                       hasDiscard = append(hasDiscard, mountPoint)
# #               }
# #       }

# #       var out bytes.Buffer
# #       if len(hasDiscard) > 0 {
# #               fmt.Fprintf(&out, "\tDon't mount the following mountpoints with attribute discard.\n")

# #               for _, i := range hasDiscard {
# #                       fmt.Fprintf(&out, "\t - %s\n", i)
# #               }
# #       }

# #       return &ProfileKV{
# #               Value:       out.String(),
# #               Description: "Avoid having a discard mount attribute as every time a file is deleted the SSD will also do a TRIM for future writing. This will increase time it takes to delete a file. Better option is to run a daily/weekly cron.",
# #       }
# # }

# # func (f *FS) limitsNoFiles() *ProfileKV {
# #       return &ProfileKV{
# #               Value: `* soft nofile unlimited
# # * hard nofile unlimited`,
# #               Description: "Every user has unlimited file descriptors available for them upping the limit from the default 1024. This allows things like increasing the number of connections etc.",
# #       }
# # }

# # func (f *FS) files() (p map[string]*ProfileKV) {
# #       p = make(map[string]*ProfileKV)

# #       p["/etc/fstab:noattime"] = f.fstabNoattime()
# #       p["/etc/fstab:discard"] = f.fstabDiscard()
# #       p["/etc/security/limits.conf"] = f.limitsNoFiles()

# #       return
# # }

# # func (f *FS) cron() (p map[string]*ProfileKV) {
# #       p = make(map[string]*ProfileKV)
# #       p["fs-trim"] = &ProfileKV{
# #               Value: `
# # #!/bin/sh
# # #
# # # To find which FS support trim, we check that DISC-MAX (discard max bytes)
# # # is great than zero. Check discard_max_bytes documentation at
# # # https://www.kernel.org/doc/Documentation/block/queue-sysfs.txt
# # #
# # for fs in $(lsblk -o MOUNTPOINT,DISC-MAX,FSTYPE | grep -E '^/.* [1-9]+.* ' | awk '{print $1}'); do
# #   fstrim "$fs"
# # done`,
# #               Schedule:    "weekly",
# #               Description: "Instead of mounting the devices with discard which slows down delete operations we should instead have a weekly cron job that goes and clears out the SSD.",
# #       }

# #       return
# # }

# # func (f *FS) procfs() (p map[string]*ProfileKV) {
# #       p = make(map[string]*ProfileKV)

# #       p["vm.dirty_ratio"] = &ProfileKV{
# #               Value:       "80",
# #               Description: "Contains, as a percentage of total available memory that contains free pages and reclaimable pages, the number of pages at which a process which is generating disk writes will itself start writing out dirty data. This value is high but should be lowered for a database application.",
# #       }

# #       p["vm.dirty_background_ratio"] = &ProfileKV{
# #               Value:       "5",
# #               Description: "Contains, as a percentage of total available memory that contains free pages and reclaimable pages, the number of pages at which the background kernel flusher threads will start writing out dirty data.",
# #       }

# #       // Reduce this.
# #       p["vm.dirty_expire_centisecs"] = &ProfileKV{
# #               Value:       "1200",
# #               Description: "This tunable is used to define when dirty data is old enough to be eligible for writeout by the kernel flusher threads.  It is expressed in 100'ths of a second.  Data which has been dirty in-memory for longer than this interval will be written out next time a flusher thread wakes up. ",
# #       }

# #       // TODO: No need to actually do this since the kernel does a pretty good job of figuring out this number.
# #       // p["fs.file-max"] = ProfileKV{
# #       //      Value:       fmt.Sprintf("%d", Stats.System.Memory.Physical.Total*11),
# #       //      Description: "The max amount of file handlers that the Linux kernel will allocate. This is one part the other part is setting the ulimits.",
# #       // }

# #       return p
# # }

# # func (f *FS) GetProfile() *Profile {
# #       p := &Profile{
# #               Name:         "fs",
# #               Subscription: StartupSubscription,
# #               Description:  "Settings for fs optimizations",
# #               References: []string{
# #                       "https://tweaked.io/guide/kernel/",
# #                       "http://blog.neutrino.es/2013/howto-properly-activate-trim-for-your-ssd-on-linux-fstrim-lvm-and-dmcrypt/",
# #               },
# #               ProcFS: f.procfs(),
# #               Files:  f.files(),
# #               Cron:   f.cron(),
# #       }

# #       return p
# # }


def main():
    attributes = [
        Memory(),
        Networking(),
        IO(),
        # FS(),
    ]

    for a in attributes:
        for procfs, feature, in a.procfs_features().iteritems():
            print("# {}".format(feature['Docs']))
            print("echo {} > {}".format(feature['Value'], procfs))
            print("")


        #print(a.sysfs_features())
        
if __name__=='__main__':
    main()

