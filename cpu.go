package main

# collect cpu and other misc load
if (open(FILE, "/proc/stat")) {
    while (my $line = <FILE>) {
        $line =~ s/^\s+|\s+$//g;
        $load{boot_time} = $1 if ($line =~ /^btime\s+(\d+)/);
        $load{d_cpu} += $1 + $2 + $3 if ($line =~ /^cpu\s+(\d+)\s+(\d+)\s+(\d+)/);
        $load{d_intrs} = $1 if ($line =~ /^intr\s+(\d+)/);
        $load{d_proc_switches} = $1 if ($line =~ /^ctxt\s+(\d+)/);
        $load{d_proc_forks} = $1 if ($line =~ /^processes\s+(\d+)/);
        $load{uptime} = $load{time} - $load{boot_time};
        $load{cpus}++ if ($line =~ /^cpu\d+/);
    }
    close FILE;
}

