<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common"
    xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>

<!--    <xsl:include href="./common/xml-instance.xsl"/>-->
    <!-- 
    input:  ../xsd/spdx-doc.xsd
    output: ../instance/spdx-doc-test-instance.xml
   -->
    <xsl:param name="TestData" select="'../instance/spdx-doc-test-data.xml'"/>
    <xsl:param name="Template" select="'../instance/spdx-doc-template.xml'"/>
    <xsl:param name="Root" select="'SpdxDocumentType'"/>

    <xsl:variable name="testData" select="document($TestData)"/>
    
    <xsl:variable name="Xsd" select="/xs:schema"/>
    
    <xsl:variable name="nsp" select="'urn:spdx-seva::1.0'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <SpdxDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns="urn:spdx-seva::1.0"
            xsi:schemaLocation="urn:spdx-seva::1.0 ../xsd/spdx-doc-iep.xsd">
            <!--<xsl:copy-of select="/xs:schema"/>-->
            <xsl:apply-templates select="document($Template)/*/*" mode="tmplt"/>
            <!--<xsl:apply-templates select="//xs:schema/xs:complexType[@name = $Root]" mode="map"/>-->
        </SpdxDocument>
    </xsl:template>

    <xsl:template match="*" mode="tmplt">
        <xsl:variable name="elref" select="name()"/>
        <xsl:variable name="elnode" select="exsl:node-set($Xsd)/xs:element[@name = $elref]"/>
        <xsl:variable name="typname" select="$elnode/@type"/>
        <xsl:variable name="typnode" select="$Xsd/*[@name = $typname]"/>
        <xsl:variable name="typbase" select="$Xsd/*[@name = $typnode/*/xs:extension/@base]"/>
        <xsl:variable name="simplebase" select="$typnode/*/xs:extension/@base[1]"/>
        <xsl:variable name="base" select="$typbase/*/@base"/>
        <xsl:variable name="testValue">
            <xsl:choose>
                <xsl:when test="$testData//*[name() = $elref]/*[@valid = 'true']">
                    <xsl:value-of select="$testData//*[name() = $elref]/*[@valid = 'true'][1]"/>
                </xsl:when>
                <xsl:when test="$simplebase = 'xs:boolean'">
                    <xsl:value-of select="$testData//*[name() = 'Boolean']/*[@valid = 'true'][1]"/>
                </xsl:when>
                <xsl:when test="$simplebase = 'xs:string'">
                    <xsl:value-of select="$testData//*[name() = 'String']/*[@valid = 'true'][1]"/>
                </xsl:when>
                <xsl:when test="$simplebase = 'xs:dateTime'">
                    <xsl:value-of select="$testData//*[name() = 'DateTime']/*[@valid = 'true'][1]"/>
                </xsl:when>
                <xsl:when test="$simplebase = 'xs:anyURI'">
                    <xsl:value-of select="$testData//*[name() = 'AnyURI']/*[@valid = 'true'][1]"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="$testData//*[name() = $typbase/@name]/*[@valid = 'true'][1]"/>
                    <xsl:value-of select="$testData//*[name() = $typname]/*[@valid = 'true'][1]"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="$elref != ''">
                <xsl:element name="{$elref}" namespace="{$nsp}">
                    <xsl:value-of select="$testValue"/>
                    <xsl:apply-templates select="*" mode="tmplt"/>
                </xsl:element>
            </xsl:when>
            <xsl:otherwise>
                <xsl:element name="{$elref}" namespace="{$nsp}">
                    <xsl:value-of select="$testValue"/>
                </xsl:element>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

</xsl:stylesheet>
