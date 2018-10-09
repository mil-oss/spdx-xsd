<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet 
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" 
    xmlns:exsl="http://exslt.org/common" 
    version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="iep.xsl"/>

    <!-- <xsl:variable name="spdx_xsd" select="document('../xsd/spdx-ref.xsd')"/>-->

    <xsl:variable name="Top" select="'AnyLicenseInfoType'"/>
    <xsl:variable name="Super" select="'SimpleLicensingInfoType'"/>
    <xsl:variable name="Root" select="'LicenseType'"/>
    <xsl:variable name="RootEl" select="'License'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:variable name="xsddata">
        <xsl:apply-templates select="/xs:schema/*[@name = $Root]"/>
        <xsl:variable name="allnodes">
            <xsl:apply-templates select="/xs:schema/*[@name = $Super]"/>
            <xsl:apply-templates select="/xs:schema/*[@name = $Top]"/>
            <xsl:apply-templates select="/xs:schema/*[@name = $RootEl]"/>
            <xsl:call-template name="deDupList">
                <xsl:with-param name="list">
                    <xsl:apply-templates select="/xs:schema/*[@name = $Root]//xs:element" mode="iterate"/>
                    <xsl:apply-templates select="/xs:schema/*[@name = $Super]//xs:element" mode="iterate"/>
                </xsl:with-param>
            </xsl:call-template>
        </xsl:variable>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:simpleType">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:complexType">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:element">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
    </xsl:variable>


    <xsl:template name="main">
        <xsl:call-template name="makeXSD">
            <xsl:with-param name="xsdnodes" select="$xsddata"/>
        </xsl:call-template>
    </xsl:template>

</xsl:stylesheet>
