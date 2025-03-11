set -xe

# This script updates RelaxNGSimplify, run simplify.sh to use

wget http://www.kohsuke.org/relaxng/rng2srng/rng2srng-20020831.zip
unzip rng2srng-20020831.zip
rm rng2srng-20020831.zip
rm -rf ./RelaxNGSimplify
mv ./rng2srng-20020831 ./RelaxNGSimplify
