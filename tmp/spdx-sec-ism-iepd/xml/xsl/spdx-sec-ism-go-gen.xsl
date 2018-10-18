<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./common/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-security/xml/xsd/spdx-security-iep.xsd
    output:${pdu}/spdx-xsd/src/golang/spdx-security/spdx-security-struct.go
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
        <xsl:apply-templates select="xs:schema/xs:simpleType[@name = 'ClassificationType']"/>
        <xsl:apply-templates select="xs:schema/xs:attributeGroup">
            <xsl:sort select="@name"/>
        </xsl:apply-templates>
    </xsl:template>
    
    <xsl:template match="xs:attributeGroup">
        <xsl:value-of select="concat('//', @name, ' ... ', normalize-space(xs:annotation/xs:documentation), $lf)"/>
        <xsl:value-of select="concat('type ', @name, ' struct ', $lb, $cr)"/>
        <xsl:for-each select="xs:attribute">
            <xsl:variable name="r" select="@ref"/>
            <xsl:variable name="cap" select="concat(translate(substring(@ref, 1, 1), $lowers, $uppers), substring(@ref, 2))"/>
            <xsl:variable name="strtype">
                <xsl:choose>
                    <xsl:when test="/xs:schema/xs:attribute[@name = $r]/xs:simpleType/xs:restriction/@base = 'xsd:NMTOKENS'">
                        <xsl:text>[]string</xsl:text>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:text>string</xsl:text>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:variable>
            <xsl:value-of select="concat($tabchar, $cap, $tabchar, $strtype, $tabchar, $bq, 'xml:', $qt, 'ism:', @ref, $cm, 'attr',$cm,$omitempty,$qt, ' ', $json, $qt, @ref, $cm, $omitempty, $qt, $bq, $cr)"/>
        </xsl:for-each>
        <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>
    <xsl:variable name="tabchar">
        <xsl:text>&#09;</xsl:text>
    </xsl:variable>
    <xsl:variable name="lowers" select='"abcdefghijklmnopqrstuvwxyz"'/>
    <xsl:variable name="uppers" select='"ABCDEFGHIJKLMNOPQRSTUVWXYZ"'/>
    

</xsl:stylesheet>
