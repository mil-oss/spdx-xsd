<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>
    
    <!-- 
    input:  /iepd/xml/xsd/iep.xsd
    output: /iepd/xml/instance/test_instance.xml
   -->
    
    <xsl:param name="TestData" select="'../instance/test_data.xml'"/>
    <xsl:param name="Root" select="'SoftwareEvidenceArchiveType'"/>
    
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
        <SoftwareEvidenceArchive xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns="urn:security::1.0" xsi:schemaLocation="urn:security::1.0  https://security.specchain.org/iepd/iep.xsd">
            <xsl:apply-templates select="*[not(name() = 'xsd:annotation')]"/> 
        </SoftwareEvidenceArchive>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref]">
        <xsl:variable name="elref" select="@ref"/>
        <xsl:variable name="elnode" select="//xs:schema/xs:element[@name = $elref]"/>
        <xsl:variable name="typnode" select="//xs:schema/*[@name = $elnode/@type]"/>
        <xsl:variable name="typbase" select="//xs:schema/*[@name = $typnode/*/xs:extension/@base]"/>
        <xsl:variable name="base" select="$typbase/xs:restriction/@base"/>
        <xsl:variable name="testValue">
            <xsl:value-of select="document($TestData)//*[name()=$typbase/@name]/*[@valid='true'][1]"/>
        </xsl:variable>
        <xsl:element name="{$elnode/@name}" namespace="urn:security::1.0">
            <xsl:value-of select="$testValue"/>
            <xsl:apply-templates select="$typnode/*"/>
        </xsl:element>
        <xsl:if test="@maxOccurs>1">
            <xsl:element name="{$elnode/@name}" namespace="urn:security::1.0">
                <xsl:value-of select="$testValue"/>
                <xsl:apply-templates select="$typnode/*"/>
            </xsl:element>
        </xsl:if>
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
    <xsl:template match="xs:extension">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="text()"/>
    
</xsl:stylesheet>
