<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <!-- 
    input:  /iepd/xml/xsd/iep.xsd
    output: /iepd/xml/instance/test_instance.xml
   -->
    <xsl:variable name="nspc" select="'urn:spdx-xml:1.0'"/>

    <xsl:template match="xs:schema/xs:complexType" mode="root">
        <xsl:param name="testData"/> 
        <xsl:variable name="annot" select="xs:annotation"/>
        <xsl:variable name="namevar" select="@name"/>
        <xsl:variable name="elname" select="//xs:schema/xs:element[@type = $namevar]/@name"/>
        <xsl:variable name="typevar" select="@type"/>
       <xsl:apply-templates select="*[not(name() = 'xsd:annotation')]">
            <xsl:with-param name="testData" select="$testData"/>
            <xsl:with-param name="depth" select="1"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="xs:element[@ref]">
        <xsl:param name="testData"/>
        <xsl:param name="depth"/>
        <xsl:variable name="elref" select="@ref"/>
        <xsl:variable name="elnode" select="//xs:schema/xs:element[@name = $elref]"/>
        <xsl:variable name="typname" select="$elnode/@type"/>
        <xsl:variable name="typnode" select="//xs:schema/*[@name = $typname]"/>
        <xsl:variable name="typbase" select="//xs:schema/*[@name = $typnode/*/xs:extension/@base]"/>
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
        <xsl:element name="{$elnode/@name}" namespace="{$nspc}">
            <xsl:value-of select="$testValue"/>
            <xsl:apply-templates select="$typnode/*">
                <xsl:with-param name="testData" select="$testData"/>
                <xsl:with-param name="depth" select="$depth+1"/>
            </xsl:apply-templates>
        </xsl:element>
        <xsl:if test="@maxOccurs > 1">
            <xsl:element name="{$elnode/@name}" namespace="{$nspc}">
                <xsl:value-of select="$testValue"/>
                <xsl:apply-templates select="$typnode/*">
                    <xsl:with-param name="testData" select="$testData"/>
                    <xsl:with-param name="depth" select="$depth+1"/>
                </xsl:apply-templates>
            </xsl:element>
        </xsl:if>
    </xsl:template>

    <xsl:template match="xs:extension">
        <xsl:param name="testData"/>
        <xsl:param name="depth"/>
        <xsl:variable name="b" select="@base"/>
        <xsl:choose>
            <xsl:when test="//xs:schema/xs:complexType[@name = $b]">
                <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $b]/*">
                    <xsl:with-param name="testData" select="$testData"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="*">
                    <xsl:with-param name="testData" select="$testData"/>
                    <xsl:with-param name="depth" select="$depth"/>
                </xsl:apply-templates>
            </xsl:when>
            <xsl:otherwise>
                <xsl:apply-templates select="*">
                    <xsl:with-param name="testData" select="$testData"/>
                    <xsl:with-param name="depth" select="$depth"/>
                </xsl:apply-templates>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="xs:annotation"/>

    <xsl:template match="xs:*">
        <xsl:param name="testData"/>
        <xsl:param name="depth"/>
        <xsl:apply-templates select="*">
            <xsl:with-param name="testData" select="$testData"/>
            <xsl:with-param name="depth" select="$depth+1"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="text()">
        <xsl:value-of select="."/>
    </xsl:template>

</xsl:stylesheet>
