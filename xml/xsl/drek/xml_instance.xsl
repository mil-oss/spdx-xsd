<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>
    
    <!-- 
    input:  /iepd/xml/xsd/iep.xsd
    output: /iepd/xml/instance/test_instance.xml
   -->
    
    <xsl:param name="TestData" select="'../instance/test_data.xml'"/>
    <xsl:param name="Root" select="'LicenseType'"/>
    
    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>
    
    <xsl:template name="main">
        <xsl:apply-templates select="xs:schema/xs:complexType[@name = $Root]" mode="root"/>
    </xsl:template> 
    
    <xsl:template match="xs:schema/xs:complexType" mode="root">
        <xsl:variable name="annot" select="xs:annotation"/>
        <xsl:variable name="namevar" select="@name"/>
        <xsl:variable name="elname" select="//xs:schema/xs:element[@type = $namevar]/@name"/>
        <xsl:variable name="typevar" select="@type"/>
        <License xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns="spdx:xsd::1.0" xsi:schemaLocation="spdx:xsd::1.0  ../xsd/spdx-license.xsd">
            <xsl:apply-templates select="*[not(name() = 'xsd:annotation')]"/> 
        </License>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref]">
        <xsl:variable name="elref" select="@ref"/>
        <xsl:variable name="elnode" select="//xs:schema/xs:element[@name = $elref]"/>
        <xsl:variable name="typnode" select="//xs:schema/*[@name = $elnode/@type]"/>
        <xsl:variable name="typbase" select="//xs:schema/*[@name = $typnode/*/xs:extension/@base]"/>
        <xsl:variable name="simplebase" select="$typnode/xs:simpleContent/xs:extension/@base[1]"/>
        <xsl:variable name="base" select="$typbase/xs:restriction/@base"/>
        <xsl:variable name="testValue">
            <xsl:choose>
                <xsl:when test="$simplebase='xs:boolean'">
                    <xsl:value-of select="document($TestData)//*[name()='Boolean']/*[@valid='true'][1]"/>   
                </xsl:when>
                <xsl:when test="$simplebase='xs:string'">
                    <xsl:value-of select="document($TestData)//*[name()='String']/*[@valid='true'][1]"/>   
                </xsl:when>
                <xsl:when test="$simplebase='xs:anyURI'">
                    <xsl:value-of select="document($TestData)//*[name()='AnyURI']/*[@valid='true'][1]"/>   
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="document($TestData)//*[name()=$typbase/@name]/*[@valid='true'][1]"/>                   
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:element name="{$elnode/@name}" namespace="spdx:xsd::1.0">
            <xsl:value-of select="$testValue"/>
            <xsl:apply-templates select="$typnode/*"/>
        </xsl:element>
        <xsl:if test="@maxOccurs>1">
            <xsl:element name="{$elnode/@name}" namespace="spdx:xsd::1.0">
                <xsl:value-of select="$testValue"/>
                <xsl:apply-templates select="$typnode/*"/>
            </xsl:element>
        </xsl:if>
    </xsl:template>
    
    <xsl:template match="xs:extension">
        <xsl:variable name="b" select="@base"/>
        <xsl:choose>
            <xsl:when test="//xs:schema/xs:complexType[@name = $b]">
                <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $b]/*"/>
                <xsl:apply-templates select="*"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:apply-templates select="*"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    
    <xsl:template match="xs:annotation"/>
    <xsl:template match="xs:sequence">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="xs:choice">
        <xsl:apply-templates select="*[2]"/>
    </xsl:template>
    <xsl:template match="xs:complexType">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="xs:simpleType">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="xs:simpleContent">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="xs:complexContent">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="text()"/>
    
</xsl:stylesheet>
