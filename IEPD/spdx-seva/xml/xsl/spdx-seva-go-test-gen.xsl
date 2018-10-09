<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>
    
    <xsl:include href="./common/go-gen.xsl"/>
    
    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-seva/xml/xsd/spdx-seva-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-seva/spdx-seva-struct_test.go
   -->
    
    <xsl:template match="/">
        <xsl:call-template name="maketests">
            <xsl:with-param name="testdata" select="document('../instance/spdx-seva-instance.xml')"/>
        </xsl:call-template>
    </xsl:template>
    
    
</xsl:stylesheet>
