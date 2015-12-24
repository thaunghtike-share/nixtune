package main

# collect i/o load
if (open(FILE, "/proc/diskstats")) {
    while (my $line = <FILE>) {
        $line =~ s/^\s+|\s+$//g;
        my @cols = split(/\s+/, $line);
        splice(@cols, 0, 3);
        next if (scalar(@cols) != 11);
        $load{d_io_reads} += $cols[0];
        $load{d_io_read_sectors} += $cols[2];
        $load{d_io_read_time} += $cols[3];
        $load{d_io_writes} += $cols[4];
        $load{d_io_write_sectors} += $cols[6];
        $load{d_io_write_time} += $cols[7];
    }
    close FILE;
    $load{d_io_ops} = $load{d_io_reads} + $load{d_io_writes};
    $load{d_io_sectors} = $load{d_io_read_sectors} + $load{d_io_write_sectors};
    $load{d_io_time} = $load{d_io_read_time} + $load{d_io_write_time};
}

