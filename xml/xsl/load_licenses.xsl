<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" 
    xmlns:spdx="http://spdx.org/rdf/terms#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" exclude-result-prefixes="xs" version="2.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:variable name="spdxxsd" select="document('../xsd/spdx-ref.xsd')"/>

    <xsl:variable name="rdfdir" select="'file:/home/jdn/DATA/Neushul_Solutions/Projects/SEvA/IonChannel/spdx/license-list-data/rdfxml/'"/>

    <xsl:variable name="outdir" select="'../instance/licenses/'"/>

    <xsl:variable name="licensefiles">
        <xsl:for-each select="document(concat($rdfdir, 'licenses.rdf'))/rdf:RDF/spdx:License">
            <License path="{concat($rdfdir,spdx:licenseId,'.rdf')}"/>
        </xsl:for-each>
    </xsl:variable>
    
    <xsl:variable name="licenseType" select="$spdxxsd/xs:schema/xs:complexType[@name='LicenseType']"/>

    <xsl:template name="main">
        <xsl:for-each select="$licensefiles/*">
            <xsl:apply-templates select="document(@path)/rdf:RDF/spdx:License"/>
        </xsl:for-each>
    </xsl:template>

    <xsl:template match="spdx:License">
        <xsl:variable name="n">
            <xsl:value-of select="spdx:licenseId"/>
        </xsl:variable>
        <xsl:variable name="lic" select="."/>
        <!--<xsl:value-of select="concat($outdir,$n)"/>-->
        <xsl:result-document href="{concat($outdir,$n,'.xml')}">
            <License xmlns="spdx:xsd::1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="spdx:xsd::1.0 ../../xsd/spdx-ref.xsd">
                <xsl:for-each select="$licenseType/xs:complexContent/xs:extension/xs:sequence/xs:element">
                    <xsl:variable name="n" select="@ref"/>
                    <xsl:variable name="rdfn" select="substring-after($spdxxsd/xs:schema/xs:element[@name=$n]/xs:annotation/xs:appinfo/*/@rdf,'#')"/>
                    <xsl:apply-templates select="$lic/*[substring-after(name(),':')=$rdfn]" mode="readrdf"/>
                </xsl:for-each>
            </License>
        </xsl:result-document>
    </xsl:template>

    <xsl:template match="*" mode="readrdf">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="name() = 'spdx:standardLicenseHeaderTemplate'">
                    <xsl:text>standardLicenseHeader</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="substring-after(name(), ':')"/>
                </xsl:otherwise>
            </xsl:choose>
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

    <xsl:template match="text()" mode="readrdf">
        <xsl:value-of select="."/>
    </xsl:template>

</xsl:stylesheet>
