set -xe

# This script is used to simplify all the correct schemas in the Relax Test Suite.
# It outputs these simplified grammars to the respective folders where the correct schema was found as a s.rng file.

find ./RelaxNGTestSuite/ -name '*c.rng'|while read fname; do
	echo $fname
	dir=$(dirname "${fname}")
	java -jar ./RelaxNGSimplify/rng2srng.jar $fname > $dir/s.rng
done