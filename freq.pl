my $freq = `lscpu |grep "CPU MHz"`;

my @sres = split(/ /, $freq);
my $resnum = @sres;

chop ($sres[$resnum - 1]);

print $sres[$resunum - 1];

