my $temp = `sensors |grep "id"`;

my @dres = split(/\+/, $temp);
my $d1 = $dres[1];
my @dres = split(/C/, $d1);

chop($dres[0]);
chop($dres[0]);

print $dres[0];

