<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>
    
    <xsl:include href="./common/go-gen.xsl"/>
    
    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-security/xml/xsd/spdx-security-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-security/spdx-security-struct_test.go
   -->
    
    <xsl:template match="/">
        <xsl:call-template name="maketests">
            <xsl:with-param name="testdata" select="document('../instance/spdx-sec-ism-instance.xml')"/>
        </xsl:call-template>
    </xsl:template>
    
    
</xsl:stylesheet>
