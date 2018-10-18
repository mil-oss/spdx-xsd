<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:exsl="http://exslt.org/common" xmlns:str="http://exslt.org/strings"
    extension-element-prefixes="str exsl" exclude-result-prefixes="xs" version="1.0">
    
    <xsl:variable name="nspc" select="'urn:spdx-xml:1.0'"/>

    <xsl:template name="makeXSD">
        <xsl:param name="xsdnodes"/>       
        <xs:schema xmlns="urn:spdx-xml:1.0" attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="urn:spdx-xml:1.0" version="1" xmlns:xs="http://www.w3.org/2001/XMLSchema">
            <xsl:apply-templates select="exsl:node-set($xsdnodes)/*" mode="xsdcopy"/>
        </xs:schema>
    </xsl:template>

    <!-- ***************** Remove NIEM  *****************-->
   <xsl:template match="xs:attributeGroup[@ref = 'structures:SimpleObjectAttributeGroup']" mode="xsdcopy"/>
    <xsl:template match="xs:complexContent[xs:extension]" mode="xsdcopy">
        <xsl:apply-templates select="xs:extension" mode="xsdcopy"/>
    </xsl:template>
    <xsl:template match="xs:complexContent" mode="xsdcopy">
        <xsl:apply-templates select="*" mode="xsdcopy"/>
    </xsl:template>
    <xsl:template match="xs:extension" mode="xsdcopy">
        <xsl:variable name="b" select="@base"/>
        <xsl:choose>
            <xsl:when test="$b = 'structures:ObjectType'">
                <xsl:apply-templates select="*" mode="xsdcopy"/>
            </xsl:when>
            <xsl:when test="contains($b,':')">
                <xs:extension base="{$b}"/>
            </xsl:when>
            <xsl:when test="substring($b, string-length($b) - string-length('SimpleType') + 1) = 'SimpleType'">
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:value-of select="$b"/>
                        </xsl:attribute>
                    </xsl:element>
            </xsl:when>
            <xsl:otherwise>
                <xs:complexContent>
                    <xsl:element name="xs:extension">
                        <xsl:attribute name="base">
                            <xsl:value-of select="$b"/>
                        </xsl:attribute>
                        <xsl:apply-templates select="*" mode="xsdcopy"/>
                    </xsl:element>
                </xs:complexContent>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    
    <xsl:template match="*[substring(@name, string-length(@name) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']" mode="xsdcopy"/>
    <xsl:template match="*[substring(@ref, string-length(@ref) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']" mode="xsdcopy"/>
    
    <xsl:template match="xs:appinfo/*[@name | @rdf]" mode="xsdcopy">
        <xsl:element name="{name()}" namespace="{$nspc}">
            <xsl:apply-templates select="@*" mode="xsdcopy"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="*" mode="iterate">
        <xsl:variable name="br">
            <xsl:value-of select="@ref"/>
            <xsl:value-of select=".//@base"/>
        </xsl:variable>
        <xsl:apply-templates select="/xs:schema/*[@name = $br]"/>
        <xsl:variable name="t" select="/xs:schema/*[@name = $br]/@type"/>
        <xsl:apply-templates select="/xs:schema/*[@name = $t]"/>
        <xsl:apply-templates select="/xs:schema/*[@name = $t]//xs:element" mode="iterate"/>
    </xsl:template>

    <xsl:template name="iterateNode">
        <xsl:param name="node"/>
        <xsl:param name="iteration"/>
        <xsl:param name="nodelist"/>
        <xsl:copy-of select="$node"/>
        <xsl:if test="$iteration &gt; 0">
            <xsl:for-each select="$node//@*[name() = 'ref' or name() = 'type' or name() = 'base' or name() = 'substitutionGroup' or name() = 'abstract'][not(. = $node/@name)]">
                <xsl:variable name="n">
                    <xsl:value-of select="."/>
                </xsl:variable>
                <xsl:if test="not(contains($nodelist, concat($n, ',')))">
                    <xsl:choose>
                        <xsl:when test="$n = 'abstract' and $node/annotation/appinfo/Choice">
                            <xsl:variable name="s" select="$node/@name"/>
                            <xsl:for-each select="//xs:schema/*[@substitutionGroup = $s]">
                                <xsl:call-template name="iterateNode">
                                    <xsl:with-param name="node" select="."/>
                                    <xsl:with-param name="iteration" select="number($iteration - 1)"/>
                                    <xsl:with-param name="nodelist" select="$nodelist"/>
                                </xsl:call-template>
                            </xsl:for-each>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:call-template name="iterateNode">
                                <xsl:with-param name="node" select="//xs:schema/*[@name = $n]"/>
                                <xsl:with-param name="iteration" select="number($iteration - 1)"/>
                                <xsl:with-param name="nodelist" select="concat($nodelist, $n, ',')"/>
                            </xsl:call-template>
                        </xsl:otherwise>
                    </xsl:choose>
                </xsl:if>
            </xsl:for-each>
        </xsl:if>
    </xsl:template>

    <xsl:template match="*">
        <xsl:copy>
            <xsl:apply-templates select="@*"/>
            <xsl:apply-templates select="text()"/>
            <xsl:apply-templates select="*"/>
        </xsl:copy>
    </xsl:template>

    <xsl:template match="xs:element[substring(@ref, string-length(@ref) - string-length('Abstract') + 1) = 'Abstract']">
        <xsl:variable name="n" select="@ref"/>
        <xsl:element name="xs:choice">
            <xsl:for-each select="/xs:schema/xs:element[@substitutionGroup = $n]">
                <xsl:element name="xs:element">
                    <xsl:attribute name="ref">
                        <xsl:value-of select="@name"/>
                    </xsl:attribute>
                    <xsl:copy-of select="xs:annotation"/>
                </xsl:element>
            </xsl:for-each>
        </xsl:element>
    </xsl:template>

    <xsl:template match="xs:element[substring(@ref, string-length(@ref) - string-length('Representation') + 1) = 'Representation']">
        <xsl:variable name="n" select="@ref"/>
        <xsl:element name="xs:choice">
            <xsl:for-each select="/xs:schema/xs:element[@substitutionGroup = $n]">
                <xsl:element name="xs:element">
                    <xsl:attribute name="ref">
                        <xsl:value-of select="@name"/>
                    </xsl:attribute>
                    <xsl:copy-of select="xs:annotation"/>
                </xsl:element>
            </xsl:for-each>
        </xsl:element>
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

    <xsl:template match="xs:import"/>

    <xsl:template match="@*">
        <xsl:copy-of select="."/>
    </xsl:template>

    <xsl:template match="text()">
        <xsl:copy-of select="normalize-space(.)"/>
    </xsl:template>

    <!-- ***************** UTILITY XSL *****************-->

    <xsl:template match="*" mode="xsdcopy">
        <xsl:element name="{name()}">
            <xsl:apply-templates select="@*" mode="xsdcopy"/>
            <xsl:apply-templates select="text()"/>
            <xsl:apply-templates select="*" mode="xsdcopy"/>
        </xsl:element>
    </xsl:template>
    <xsl:template match="@*" mode="xsdcopy">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="contains(name(), ':')">
                    <xsl:value-of select="substring-after(name(), ':')"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="name()"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="ns" select="substring-before(name(), ':')"/>
        <xsl:attribute name="{$n}">
            <xsl:value-of select="."/>
        </xsl:attribute>
    </xsl:template>

<!--    <xsl:template match="text()">
        <xsl:value-of select="normalize-space(translate(., '&#34;&#xA;', ''))"/>
    </xsl:template>-->

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

    <!-- *****deDupList ****-->
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


    <!-- Ends-with XSL 1.0-->
    <xsl:template match="*[substring(@name, string-length(@name) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']"/>

    <!-- Ends-with XSL 1.0-->
    <xsl:template match="*[substring(@ref, string-length(@ref) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint']"/>

</xsl:stylesheet>
