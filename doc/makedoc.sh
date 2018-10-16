#!/bin/bash
java -cp /home/jdn/DATA/Neushul_Solutions/Projects/XML/saxon9he.jar net.sf.saxon.Transform -s:/home/jdn/DATA/Neushul_Solutions/Projects/security/IonChannel/spdx-xsd/doc/security_SPDX.mm -xsl:/home/jdn/DATA/Neushul_Solutions/Projects/XML/MM_Naval_Ltr_Format.xsl
