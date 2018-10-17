<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:ism="urn:us:gov:ic:ism" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>
    
    <!-- 
    input:  /iepd/xml/xsd/ismiep.xsd
    output: /iepd/xml/instance/test_instance-ism.xml
   -->
    
    <xsl:param name="TestData" select="'../instance/test_data.xml'"/>
   
    <xsl:template match="/">
        <xsl:call-template name="main"/>
        <!--<xsl:apply-templates select="xs:schema/xs:complexType[@name = 'SoftwareEvidenceArchiveType']" mode="root"/>-->
    </xsl:template>
    
    <xsl:template name="main">
        <!--<xsl:result-document href="{$path}">-->
        <xsl:apply-templates select="xs:schema/xs:complexType[@name = 'SoftwareEvidenceArchiveType']" mode="root"/>
        <!--</xsl:result-document>-->
    </xsl:template> 
    
    <xsl:template match="xs:schema/xs:complexType" mode="root">
        <xsl:variable name="annot" select="xs:annotation"/>
        <xsl:variable name="namevar" select="@name"/>
        <xsl:variable name="elname" select="//xs:schema/xs:element[@type = $namevar]/@name"/>
        <xsl:variable name="typevar" select="@type"/>
        <SoftwareEvidenceArchive xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns:ism="urn:us:gov:ic:ism" xmlns="urn:seva::1.0" xsi:schemaLocation="urn:seva::1.0 file:./../xsd/iep.xsd">
                <xsl:attribute name="ism:classification">
                    <xsl:text>U</xsl:text>
                </xsl:attribute>
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
        <xsl:element name="{$elnode/@name}" namespace="urn:seva::1.0">
                <xsl:attribute name="ism:classification">
                    <xsl:text>U</xsl:text>
                </xsl:attribute>
            <xsl:value-of select="$testValue"/>
            <xsl:apply-templates select="$typnode/*"/>
        </xsl:element>
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
