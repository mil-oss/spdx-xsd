<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:spdx="http://spdx.org/rdf/terms#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:exsl="http://exslt.org/common" exclude-result-prefixes="xs rdf exsl spdx" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:param name="spdxxsd" select="document('../xsd/spdx-license.xsd')"/>
    <xsl:variable name="simpleLicenseType" select="$spdxxsd/xs:schema/xs:complexType[@name = 'SimpleLicensingInfoType']"/>
    <xsl:variable name="licenseType" select="$spdxxsd/xs:schema/xs:complexType[@name = 'LicenseType']"/>
    <!--<xsl:variable name="licenseList" select="document('../../resources/licenses.rdf')"/>-->

    <xsl:template match="/">
        <xsl:apply-templates select="rdf:RDF/spdx:License"/>
    </xsl:template>

    <xsl:template match="spdx:License">
        <xsl:variable name="lic" select="."/>
        <License xmlns="spdx:xsd::1.0">
            <xsl:for-each select="$simpleLicenseType/xs:complexContent/xs:extension/xs:sequence/xs:element">
                <xsl:variable name="r" select="@ref"/>
                <xsl:variable name="t" select="$spdxxsd/xs:schema/xs:element[@name = $r]/@type"/>
                <xsl:variable name="b" select="$spdxxsd/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>

                <xsl:variable name="rdfn" select="substring-after($spdxxsd/xs:schema/xs:element[@name = $r]/xs:annotation/xs:appinfo/*/@rdf, '#')"/>
                <xsl:choose>
                    <xsl:when test="$b = 'xs:boolean' and not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$r}" xmlns="spdx:xsd::1.0">false</xsl:element>
                    </xsl:when>
                    <xsl:when test="not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$r}" xmlns="spdx:xsd::1.0"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:apply-templates select="$lic/*[substring-after(name(), ':') = $rdfn]" mode="readrdf"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:for-each>
            <xsl:for-each select="$licenseType/xs:complexContent/xs:extension/xs:sequence/xs:element">
                <xsl:variable name="rr" select="@ref"/>
                <xsl:variable name="t" select="$spdxxsd/xs:schema/xs:element[@name = $rr]/@type"/>
                <xsl:variable name="b" select="$spdxxsd/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
                <xsl:variable name="rdfn" select="substring-after($spdxxsd/xs:schema/xs:element[@name = $rr]/xs:annotation/xs:appinfo/*/@rdf, '#')"/>
                <xsl:choose>
                    <xsl:when test="$b = 'xs:boolean' and not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$rr}" xmlns="spdx:xsd::1.0">false</xsl:element>
                    </xsl:when>
                    <xsl:when test="not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$rr}" xmlns="spdx:xsd::1.0"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:apply-templates select="$lic/*[substring-after(name(), ':') = $rdfn]" mode="readrdf"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:for-each>
        </License>
    </xsl:template>

    <xsl:template match="*" mode="readrdf">
        <xsl:variable name="n">
            <xsl:value-of select="substring-after(name(), ':')"/>
        </xsl:variable>
        <xsl:variable name="ns" select="namespace-uri(.)"/>
        <xsl:variable name="rdfid" select="concat($ns, $n)"/>
        <xsl:variable name="match" select="$spdxxsd/xs:schema/xs:element[xs:annotation/xs:appinfo/*/@rdf = $rdfid][1]"/>
        <xsl:choose>
            <xsl:when test="$match/@name">
                <xsl:element name="{$match/@name}" xmlns="spdx:xsd::1.0">
                    <xsl:apply-templates select="text()" mode="readrdf"/>
                    <xsl:apply-templates select="*" mode="readrdf"/>
                </xsl:element>
            </xsl:when>
            <xsl:otherwise>
                <ERROR/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="text()" mode="readrdf">
        <xsl:value-of select="."/>
    </xsl:template>
</xsl:stylesheet>
