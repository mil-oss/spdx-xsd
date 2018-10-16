<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>
    
    <xsl:include href="./common/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-document/xml/xsd/spdx-doc-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-doc/spdx-doc-struct_test.go
   -->
    
    <xsl:variable name="testdata" select="document('../instance/spdx-doc-instance.xml')"/>
    
    <xsl:template match="/">
        <!--<xsl:call-template name="maketests">
            <xsl:with-param name="testdata" select="document('../instance/spdx-doc-instance.xml')"/>
        </xsl:call-template>-->
        <xsl:variable name="rootname" select="//xs:schema/xs:annotation/xs:appinfo/*/@name"/>
        <xsl:variable name="roottype" select="//xs:schema/xs:annotation/xs:appinfo/*/@type"/>
        <xsl:call-template name="teststart">
            <xsl:with-param name="appname" select="$rootname"/>
        </xsl:call-template>
        <xsl:variable name="b" select="//xs:schema/xs:complexType[@name = $roottype]//@base"/>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="maketest">
            <xsl:with-param name="rootname" select="$rootname"/>
            <xsl:with-param name="testdata" select="$testdata"/>
        </xsl:apply-templates>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $roottype]//xs:element[@ref]" mode="maketest">
            <xsl:with-param name="rootname" select="$rootname"/>
            <xsl:with-param name="testdata" select="$testdata"/>
        </xsl:apply-templates>
        <xsl:value-of select="concat($in, '})', $cr)"/>
        <xsl:value-of select="concat($cr, '}')"/>
    </xsl:template>
    
    <xsl:template match="*[@ref='DescribesFile']" mode="maketest"/>
    <xsl:template match="*[@ref='DescribesPackage']" mode="maketest"/>
    <xsl:template match="*[@ref='HasExtractedLicensingInfo']" mode="maketest"/>
    
</xsl:stylesheet>
