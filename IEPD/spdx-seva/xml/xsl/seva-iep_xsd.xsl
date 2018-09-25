<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:spd="spdx:xsd::1.0/ref" xmlns:exsl="http://exslt.org/common" version="1.0">

    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="identity.xsl"/>

    <xsl:variable name="Super" select="'SpdxElementType'"/>
    <xsl:variable name="Root" select="'SoftwareEvidenceArchiveType'"/>
    <xsl:variable name="RootEl" select="'SoftwareEvidenceArchive'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <!--<xsl:result-document href="../xsd/spdx-doc-iep.xsd">-->
        <xs:schema xmlns="spdx:xsd::1.0" attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="spdx:xsd::1.0" version="1" xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
            <xsl:apply-templates select="//xs:schema/*[@name = $Root]" mode="identity"/>
            <!--<xsl:apply-templates select="/xs:schema/*[@name = $Root]"/>-->
            <xsl:variable name="allnodes">
                <xsl:apply-templates select="//xs:schema/*[@name = $Super]"/>
                <!--<xsl:apply-templates select="//xs:schema/*[@name = $RootEl]"/>-->
                <xsl:apply-templates select="//xs:schema/*[@name = 'AlgorithmCodeSimpleType']"/>
                <xsl:apply-templates select="//xs:schema/*[@name = 'AnnotationTypeCodeSimpleType']"/>
                <xsl:apply-templates select="//xs:schema/*[@name = 'RelationshipTypeCodeSimpleType']"/>
                <xsl:call-template name="deDupList">
                    <xsl:with-param name="list">
                        <xsl:apply-templates select="//xs:schema/*[@name = $Root]//xs:element" mode="iterate"/>
                        <xsl:apply-templates select="//xs:schema/*[@name = $Super]//xs:element" mode="iterate"/>
                        <xsl:apply-templates select="//xs:schema/xs:element[@name = $RootEl]"/>
                    </xsl:with-param>
                </xsl:call-template>
            </xsl:variable>
            <xsl:for-each select="exsl:node-set($allnodes)/xs:simpleType">
                <xsl:sort select="@name"/>
                <xsl:copy-of select="."/>
            </xsl:for-each>
            <xsl:for-each select="exsl:node-set($allnodes)/xs:complexType[not(@name = $Root)]">
                <xsl:sort select="@name"/>
                <xsl:copy-of select="."/>
            </xsl:for-each>
            <xsl:for-each select="exsl:node-set($allnodes)/xs:element">
                <xsl:sort select="@name"/>
                <xsl:copy-of select="."/>
            </xsl:for-each>
        </xs:schema>
        <!--</xsl:result-document>-->
    </xsl:template>

    <xsl:template match="*" mode="iterate">
        <xsl:variable name="br">
            <xsl:choose>
                <xsl:when test="@ref">
                    <xsl:value-of select="@ref"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select=".//@base"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:apply-templates select="//xs:schema/xs:element[@name = $br]"/>
        <xsl:variable name="t" select="//xs:schema/xs:element[@name = $br]/@type"/>
        <xsl:apply-templates select="//xs:schema/*[@name = $t]"/>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $t]/*" mode="iterate"/>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $t]/xs:sequence/xs:element" mode="iterate"/>
    </xsl:template>

    <xsl:template match="xs:sequence" mode="iterate">
        <xsl:apply-templates select="*" mode="iterate"/>
    </xsl:template>

    <xsl:template match="*">
        <xsl:copy>
            <xsl:apply-templates select="@*"/>
            <xsl:apply-templates select="text()"/>
            <xsl:apply-templates select="*"/>
        </xsl:copy>
    </xsl:template>

    <xsl:template match="@mapvar" mode="identity"/>

    <xsl:template match="xs:element[substring(@ref, string-length(@ref) - string-length('Representation') + 1) = 'Representation']">
        <xsl:variable name="n" select="@ref"/>
        <xsl:element name="xs:choice">
            <xsl:for-each select="//xs:schema/xs:element[@substitutionGroup = $n]">
                <xsl:element name="xs:element">
                    <xsl:attribute name="ref">
                        <xsl:value-of select="@name"/>
                    </xsl:attribute>
                    <xsl:copy-of select="xs:annotation"/>
                </xsl:element>
            </xsl:for-each>
        </xsl:element>
    </xsl:template>

    <xsl:template match="xs:complexContent">
        <xsl:apply-templates select="*" mode="iterate"/>
    </xsl:template>

    <xsl:template match="xs:simpleContent[not(xs:restriction)]">
        <xs:simpleContent>
            <xs:extension>
                <xsl:copy-of select="xs:extension/@base"/>
            </xs:extension>
        </xs:simpleContent>
    </xsl:template>

    <xsl:template match="xs:element/xs:annotation/xs:appinfo/*">
        <!-- <xsl:variable name="xpath">
            <xsl:apply-templates select="ancestor::xs:element" mode="makeXpath"/>
        </xsl:variable>-->
        <xsl:copy>
            <xsl:apply-templates select="@*"/>
            <!--<xsl:attribute name="xpath">
                <xsl:value-of select="concat(substring-before($xpath, 'xs:complexContent/xs:extension'), substring-after($xpath, 'xs:complexContent/xs:extension'))"/>
            </xsl:attribute>-->
        </xsl:copy>
    </xsl:template>

        <xsl:template match="xs:extension">
        <xsl:variable name="b" select="@base"/>
        <xsl:choose>
            <xsl:when test="$b = 'xs:boolean'">
                <xsl:element name="xs:simpleContent">
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:value-of select="$b"/>
                        </xsl:attribute>
                    </xsl:element>
                </xsl:element>
            </xsl:when>
            <xsl:when test="$b = 'xs:float'">
                <xsl:element name="xs:simpleContent">
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:value-of select="$b"/>
                        </xsl:attribute>
                    </xsl:element>
                </xsl:element>
            </xsl:when>
            <xsl:when test="$b = 'structures:ObjectType'">
                <xsl:element name="xs:simpleContent">
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:text>SpdxElement</xsl:text>
                        </xsl:attribute>
                    </xsl:element>
                </xsl:element>
            </xsl:when>
            <xsl:when test="contains($b, ':')">
                <xsl:apply-templates select="*"/>
            </xsl:when>
            <xsl:otherwise>
                <xs:complexContent>
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:value-of select="$b"/>
                        </xsl:attribute>
                        <xsl:apply-templates select="*"/>
                    </xsl:element>
                </xs:complexContent>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="xs:import"/>

    <!-- Ends-with XSL 1.0-->
    <xsl:template match="xs:element[substring(@name, string-length(@name) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']"/>
    <xsl:template match="xs:element[substring(@ref, string-length(@ref) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']"/>
    <xsl:template match="xs:element[substring(@ref, string-length(@ref) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']" mode="identity"/>

    <xsl:template match="xs:attributeGroup[@ref = 'structures:SimpleObjectAttributeGroup']"/>

    <xsl:template match="@*">
        <xsl:copy-of select="."/>
    </xsl:template>

    <xsl:template match="text()">
        <xsl:copy-of select="normalize-space(.)"/>
    </xsl:template>

    <!-- ***************** UTILITY XSL *****************-->

    <!-- *****SPLIT CAMEL CASE ****-->
    <xsl:template name="CamelCase">
        <xsl:param name="text"/>
        <xsl:choose>
            <xsl:when test="contains($text, ' ')">
                <xsl:call-template name="CamelCaseWord">
                    <xsl:with-param name="text" select="substring-before($text, ' ')"/>
                </xsl:call-template>
                <xsl:text> </xsl:text>
                <xsl:call-template name="CamelCase">
                    <xsl:with-param name="text" select="substring-after($text, ' ')"/>
                </xsl:call-template>
            </xsl:when>
            <xsl:otherwise>
                <xsl:call-template name="CamelCaseWord">
                    <xsl:with-param name="text" select="$text"/>
                </xsl:call-template>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="CamelCaseWord">
        <xsl:param name="text"/>
        <xsl:value-of select="translate(substring($text, 1, 1), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ')"/>
        <xsl:value-of select="translate(substring($text, 2, string-length($text) - 1), 'ABCDEFGHIJKLMNOPQRSTUVWXYZ', 'abcdefghijklmnopqrstuvwxyz')"/>
    </xsl:template>

    <xsl:template name="breakIntoWords">
        <xsl:param name="string"/>
        <xsl:choose>
            <xsl:when test="string-length($string) &lt; 2">
                <xsl:value-of select="$string"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:call-template name="breakIntoWordsHelper">
                    <xsl:with-param name="string" select="$string"/>
                    <xsl:with-param name="token" select="substring($string, 1, 1)"/>
                </xsl:call-template>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="breakIntoWordsHelper">
        <xsl:param name="string" select="''"/>
        <xsl:param name="token" select="''"/>
        <xsl:choose>
            <xsl:when test="string-length($string) = 0"/>
            <xsl:when test="string-length($token) = 0"/>
            <xsl:when test="string-length($string) = string-length($token)">
                <xsl:value-of select="$token"/>
            </xsl:when>
            <xsl:when test="contains('ABCDEFGHIJKLMNOPQRSTUVWXYZ', substring($string, string-length($token) + 1, 1))">
                <xsl:value-of select="concat($token, ' ')"/>
                <xsl:call-template name="breakIntoWordsHelper">
                    <xsl:with-param name="string" select="substring-after($string, $token)"/>
                    <xsl:with-param name="token" select="substring($string, string-length($token), 1)"/>
                </xsl:call-template>
            </xsl:when>
            <xsl:otherwise>
                <xsl:call-template name="breakIntoWordsHelper">
                    <xsl:with-param name="string" select="$string"/>
                    <xsl:with-param name="token" select="substring($string, 1, string-length($token) + 1)"/>
                </xsl:call-template>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="removeStrings">
        <xsl:param name="targetStr"/>
        <xsl:param name="strings"/>
        <xsl:variable name="str">
            <xsl:choose>
                <xsl:when test="contains($strings, ',')">
                    <xsl:value-of select="normalize-space(substring-before($strings, ','))"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="$strings"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="string-length($str) &gt; 0">
                <xsl:call-template name="removeStrings">
                    <xsl:with-param name="targetStr" select="translate($targetStr, $str, '')"/>
                    <xsl:with-param name="strings" select="substring-after($strings, ',')"/>
                </xsl:call-template>
            </xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="$targetStr"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <!-- *****XPATH ****-->
    <xsl:variable name="q" select="'&quot;'"/>
    <xsl:variable name="a" select='"&apos;"'/>
    <xsl:template match="*" mode="makeXpath">
        <xsl:for-each select="ancestor-or-self::*">
            <xsl:value-of select="name()"/>
            <xsl:variable name="n" select="@name"/>
            <xsl:variable name="r" select="@ref"/>
            <xsl:if test="$n">
                <xsl:value-of select="concat('[@name=', $a, $n, $a, ']')"/>
            </xsl:if>
            <xsl:if test="$r">
                <xsl:value-of select="concat('[@ref=', $a, $r, $a, ']')"/>
            </xsl:if>
            <xsl:if test="position() != last()">
                <xsl:text>/</xsl:text>
            </xsl:if>
        </xsl:for-each>
    </xsl:template>
    <xsl:template name="deDupList">
        <xsl:param name="list"/>
        <xsl:for-each select="exsl:node-set($list)/*">
            <xsl:sort select="@name"/>
            <xsl:sort select="@value"/>
            <xsl:variable name="n" select="@name"/>
            <xsl:variable name="v" select="@value"/>
            <xsl:choose>
                <xsl:when test="preceding-sibling::*[@name = $n]"/>
                <xsl:when test="preceding-sibling::*[@value = $v]"/>
                <xsl:otherwise>
                    <xsl:copy-of select="."/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:for-each>
    </xsl:template>


</xsl:stylesheet>
