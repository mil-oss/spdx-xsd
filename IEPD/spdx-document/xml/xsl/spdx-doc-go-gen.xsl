<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./common/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-document/xml/xsd/spdx-doc-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-doc/spdx-doc-struct.go
   -->


    <xsl:template match="/">
        <xsl:variable name="rootname" select="xs:schema/xs:annotation/xs:appinfo/*/@name"/>
        <xsl:value-of select="concat('package main', $cr, $cr)"/>
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
    
    <xsl:template match="xs:element[@ref = 'RelatedSpdxElement']" mode="makevar">
        <xsl:value-of select="concat($tab, 'RelatedSpdxElement', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'RelatedSpdxElement', $cm, $omitempty, $qt, ' ', $json, $qt, 'RelatedSpdxElement', $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref = 'DescribesPackage']" mode="makevar">
        <xsl:value-of select="concat($tab, 'DescribesPackage', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'DescribesPackage', $cm, $omitempty, $qt, ' ', $json, $qt, 'DescribesPackage', $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref = 'DescribesFile']" mode="makevar">
        <xsl:value-of select="concat($tab, 'DescribesFile', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'DescribesFile', $cm, $omitempty, $qt, ' ', $json, $qt, 'DescribesFile', $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref = 'HasExtractedLicensingInfo']" mode="makevar">
        <xsl:value-of select="concat($tab, 'HasExtractedLicensingInfo', $tab, 'bool', $tab, $tab, $bq, 'xml:', $qt, 'HasExtractedLicensingInfo', $cm, $omitempty, $qt, ' ', $json, $qt, 'HasExtractedLicensingInfo', $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>

</xsl:stylesheet>
