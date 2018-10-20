<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema"
    exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./common/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-document/xml/xsd/spdx-doc-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-doc/spdx-doc-struct.go
   -->


    <xsl:template match="/">
        <xsl:variable name="rootname" select="xs:schema/xs:annotation/xs:appinfo/*/@name"/>
        <xsl:value-of select="concat('package spdxdoc', $cr, $cr)"/>
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
        <xsl:value-of
            select="concat('// RelatedSpdxElement ... ', substring-before(xs:annotation/xs:documentation, '.'), $cr)"/>
        <xsl:value-of select="concat('type ', 'RelatedSpdxElement', ' struct ', $lb, $cr)"/>
        <xsl:value-of
            select="concat($tab, 'Annotation', $tab, '*Annotation', $tab, $tab, $bq, 'xml:', $qt, 'Annotation', $cm, $omitempty, $qt, ' ', $json, $qt, 'Annotation', $cm, $omitempty, $qt, $bq, $cr)"/>
        <xsl:value-of
            select="concat($tab, 'Name', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'Name', $cm, $omitempty, $qt, ' ', $json, $qt, 'Name', $cm, $omitempty, $qt, $bq, $cr)"/>
        <xsl:value-of
            select="concat($tab, 'CommentText', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'CommentText', $cm, $omitempty, $qt, ' ', $json, $qt, 'CommentText', $cm, $omitempty, $qt, $bq, $cr)"/>
        <xsl:value-of
            select="concat($tab, 'XMLName', $tab, 'xml.Name', $tab, $tab, $bq, 'xml:', $qt, 'RelatedSpdxElement', $cm, $omitempty, $qt, ' ', $json, $qt, 'RelatedSpdxElement', $cm, $omitempty, $qt, $bq, $cr)"/>
        <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>

</xsl:stylesheet>
