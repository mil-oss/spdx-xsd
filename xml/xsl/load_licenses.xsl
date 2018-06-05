<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:spdx="http://spdx.org/rdf/terms#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:exsl="http://exslt.org/common" exclude-result-prefixes="xs" version="2.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="read_license.xsl"/>

    <xsl:variable name="licencesrdf" select="document('https://raw.githubusercontent.com/spdx/license-list-data/master/rdfxml/licenses.rdf')"/>

    <xsl:variable name="outdir" select="'../../resources/xml-licenses/'"/>

    <xsl:variable name="licensefiles">
        <Licenses>
            <xsl:apply-templates select="$licencesrdf/rdf:RDF/spdx:License" mode="readlicense"/>
        </Licenses>
    </xsl:variable>

    <xsl:template name="main">
        <xsl:copy-of select="." copy-namespaces="no"/>
        <xsl:for-each select="document($licensefiles)/Licenses/*">
            <xsl:if test="LicenseID">
                <xsl:result-document href="{concat($outdir,LicenseID)}">
                    <xsl:copy-of select="." copy-namespaces="no"/>
                </xsl:result-document>
            </xsl:if>
        </xsl:for-each>
    </xsl:template>

    <!--    <xsl:template match="spdx:License">
        <xsl:variable name="licid">
            <xsl:value-of select="spdx:licenseId"/>
        </xsl:variable>
        <xsl:variable name="lic" select="."/>
        <License xmlns="spdx:xsd::1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="spdx:xsd::1.0 ../../xml/xsd/spdx-ref.xsd">
            <xsl:for-each select="exsl:node-set($simpleLicenseType)/xs:complexContent/xs:extension/xs:sequence/xs:element">
                <xsl:variable name="n" select="@ref"/>
                <xsl:variable name="t" select="$spdxxsd/xs:schema/xs:element[@name = $n]/@type"/>
                <xsl:variable name="b" select="$spdxxsd/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>

                <xsl:variable name="rdfn" select="substring-after($spdxxsd/xs:schema/xs:element[@name = $n]/xs:annotation/xs:appinfo/*/@rdf, '#')"/>
                <xsl:choose>
                    <xsl:when test="$b = 'xs:boolean' and not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$n}" xmlns="spdx:xsd::1.0">false</xsl:element>
                    </xsl:when>
                    <xsl:when test="not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$n}" xmlns="spdx:xsd::1.0"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:apply-templates select="$lic/*[substring-after(name(), ':') = $rdfn]" mode="readrdf"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:for-each>
            <xsl:for-each select="$licenseType/xs:complexContent/xs:extension/xs:sequence/xs:element">
                <xsl:variable name="n" select="@ref"/>
                <xsl:variable name="t" select="$spdxxsd/xs:schema/xs:element[@name = $n]/@type"/>
                <xsl:variable name="b" select="$spdxxsd/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
                <xsl:variable name="rdfn" select="substring-after($spdxxsd/xs:schema/xs:element[@name = $n]/xs:annotation/xs:appinfo/*/@rdf, '#')"/>
                <xsl:choose>
                    <xsl:when test="$b = 'xs:boolean' and not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$n}" xmlns="spdx:xsd::1.0">false</xsl:element>
                    </xsl:when>
                    <xsl:when test="not($lic/*[substring-after(name(), ':') = $rdfn])">
                        <xsl:element name="{$n}" xmlns="spdx:xsd::1.0"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:apply-templates select="$lic/*[substring-after(name(), ':') = $rdfn]" mode="readrdf"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:for-each>
        </License>
    </xsl:template>
-->
    <!--<xsl:template match="*" mode="readrdf">
        <xsl:variable name="n">
            <xsl:value-of select="substring-after(name(), ':')"/>
        </xsl:variable>
        <xsl:variable name="ns" select="namespace-uri(.)"/>
        <xsl:variable name="rdfid" select="concat($ns, $n)"/>
        <xsl:variable name="match">
            <xsl:copy-of select="$spdxxsd/xs:schema/xs:element[xs:annotation/xs:appinfo/*/@rdf = $rdfid]"/>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="$match/*/@name">
                <xsl:element name="{$match/*/@name}" xmlns="spdx:xsd::1.0">
                    <xsl:apply-templates select="text()" mode="readrdf"/>
                    <xsl:apply-templates select="*" mode="readrdf"/>
                </xsl:element>
            </xsl:when>
            <xsl:otherwise>
                <ERROR/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
-->
    <!--    <xsl:template match="text()" mode="readrdf">
        <xsl:value-of select="."/>
    </xsl:template>-->

</xsl:stylesheet>
