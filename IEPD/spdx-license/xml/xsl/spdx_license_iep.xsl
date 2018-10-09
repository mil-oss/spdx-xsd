<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:exsl="http://exslt.org/common" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./../../../../xml/xsl/iep.xsl"/>

    <!-- <xsl:variable name="spdx_xsd" select="document('../xsd/spdx-seva-ref.xsd')"/>-->

    <xsl:variable name="Top" select="'AnyLicenseInfoType'"/>
    <xsl:variable name="Super" select="'SimpleLicensingInfoType'"/>
    <xsl:variable name="Root" select="'LicenseType'"/>
    <xsl:variable name="RootEl" select="'License'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>
    <xsl:variable name="xsddata">
        <xs:annotation>
            <xs:documentation>XML Schema for SPDX License Information Exchange</xs:documentation>
            <xs:appinfo>
                <Root type="{$Root}" name="{$RootEl}"/>
            </xs:appinfo>
        </xs:annotation>
        <xsl:variable name="allnodes">
            <xsl:call-template name="deDupList">
                <xsl:with-param name="list">
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $RootEl]"/>
                        <xsl:with-param name="iteration" select="15"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Top]"/>
                        <xsl:with-param name="iteration" select="15"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Root]"/>
                        <xsl:with-param name="iteration" select="15"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Super]"/>
                        <xsl:with-param name="iteration" select="15"/>
                    </xsl:call-template>
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
