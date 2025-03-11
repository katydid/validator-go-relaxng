=========================================================================


            RELAX NG full syntax to simple syntax translator


=========================================================================
                                    by Kohsuke Kawaguchi (kk@kohsuke.org)


This tool reads RELAX NG full XML syntax [1] or RELAX NG compact syntax [2]
and produces the equivalent RELAX NG simplified syntax [3].

The XML produced by this tool is primarily intended for machine
consumption, but you might find it useful for studying the spec.


LICENSE
=======

This software is licensed under the BSD license [4].


USING AS A TOOL
===============

Type the following command from a command prompt:

$ java -jar rng2srng.jar

And it will show you the usage.


USING AS A LIBRARY
==================

You can integrate this tool into your own application.
Potential benefits include

- rigorous error checks of the input grammar is free
- your tool can be simplified


The core functionality of this library is to take
org.xml.sax.InputSource object as an input and produce SAX2
events. 

The org.kohsuke.rng2srng.Translator class defines a series
of utility methods that makes integration a snap. For details,
see the source code.


CONTACT
=======

I'd appreciate any comments/questions. Please send them to kk@kohsuke.org



REFERENCES
==========

[1] http://www.oasis-open.org/committees/relax-ng/tutorial-20011203.html
    http://www.oasis-open.org/committees/relax-ng/spec-20011203.html

[2] http://www.xml.com/pub/a/2002/06/19/rng-compact.html
    http://www.oasis-open.org/committees/relax-ng/compact-20020607.html
    
[3] http://www.oasis-open.org/committees/relax-ng/spec-20011203.html#simple-syntax

[4] http://www.opensource.org/licenses/bsd-license.php