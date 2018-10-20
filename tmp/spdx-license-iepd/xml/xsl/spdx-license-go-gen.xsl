<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./common/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-license/xml/xsd/spdx-license-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-license/spdx-license-struct.go
   -->
    
        <xsl:template match="/">
            <xsl:variable name="rootname" select="xs:schema/xs:annotation/xs:appinfo/*/@name"/>
            <xsl:value-of select="concat('package spdxlic', $cr, $cr)"/>
            <xsl:value-of select="concat('import ', $qt, 'encoding/xml', $qt, $cr, $cr)"/>
            <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]" mode="func">
                <xsl:with-param name="rootname" select="$rootname"/>
            </xsl:apply-templates>
            <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]">
                <xsl:with-param name="rootname" select="$rootname"/>
            </xsl:apply-templates>
            <xsl:apply-templates select="xs:schema/xs:element[not(@name = $rootname)]">
                <xsl:with-param name="rootname" select="$rootname"/>
                <xsl:sort select="@name"/>
            </xsl:apply-templates>
        </xsl:template>
    

</xsl:stylesheet>
