<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    exclude-result-prefixes="xs" version="1.0">

    <!-- 
    input: /iepd/xml/xsd/ic-ism/IC-ISM-v2.1.xsd
    output:/iepd/src/goxsd/ic-ism_struct.go
   -->

    <xsl:output method="text" indent="yes"/>

    <xsl:variable name="a">
        <xsl:text>&amp;</xsl:text>
    </xsl:variable>
    <xsl:variable name="lb">
        <xsl:text>{</xsl:text>
    </xsl:variable>
    <xsl:variable name="rb">
        <xsl:text>}</xsl:text>
    </xsl:variable>
    <xsl:variable name="cm">
        <xsl:text>,</xsl:text>
    </xsl:variable>
    <xsl:variable name="cr">
        <xsl:text>&#10;</xsl:text>
    </xsl:variable>
    <xsl:variable name="lf">
        <xsl:text>&#10;</xsl:text>
    </xsl:variable>
    <xsl:variable name="qt">
        <xsl:text>&quot;</xsl:text>
    </xsl:variable>
    <xsl:variable name="bq">
        <xsl:text>&#96;</xsl:text>
    </xsl:variable>
    <xsl:variable name="sp">
        <xsl:text>&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="in">
        <xsl:value-of select="concat($sp, $sp, $sp, $sp)"/>
    </xsl:variable>
    <xsl:variable name="tab" select="25"/>
    <xsl:variable name="lowers" select='"abcdefghijklmnopqrstuvwxyz"'/>
    <xsl:variable name="uppers" select='"ABCDEFGHIJKLMNOPQRSTUVWXYZ"'/>
    <xsl:variable name="json" select="' json:'"/>
    <xsl:variable name="omitempty" select="'omitempty'"/>
    <xsl:variable name="ns" select="xsd:schema/@targetNamespace"/>

    <xsl:template match="/">
        <xsl:value-of select="concat('package main', $cr, $cr)"/>
        <xsl:value-of select="concat('import ', $qt, 'encoding/xml', $qt, $cr, $cr)"/>
        <xsl:apply-templates select="xsd:schema/xsd:simpleType[@name = 'ClassificationType']"/>
        <xsl:apply-templates select="xsd:schema/xsd:attributeGroup">
            <xsl:sort select="@name"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="*">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="*" mode="def">
        <xsl:apply-templates select="*" mode="def"/>
    </xsl:template>
    <xsl:template match="text()"/>

    <xsl:template name="mkspc">
        <xsl:param name="spc"/>
        <xsl:value-of select="$sp"/>
        <xsl:if test="number($spc) &gt; 0">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($spc) - 1"/>
            </xsl:call-template>
        </xsl:if>
    </xsl:template>

<!--    <xsl:template match="xsd:simpleType"> 
        <xsl:value-of select="concat('//', @name, ' ... ', normalize-space(xsd:annotation/xsd:documentation), $lf)"/>
        <xsl:value-of select="concat('type ', 'ClassificationEnum', ' int',$cr)"/>
        <xsl:value-of select="concat('const (',$cr)"/>
        <xsl:variable name="c1" select="xsd:restriction/xsd:enumeration[1]/@value"/>
        <xsl:value-of select="concat('//', $c1, ' ... ', $lf)"/>
        <xsl:value-of select="concat($in,$c1,' ClassificationEnum = 1+ iota',$cr)"/>
        <xsl:for-each select="xsd:restriction/xsd:enumeration[position()&gt;1]">
            <xsl:variable name="c" select="@value"/>
            <xsl:variable name="nspcs">
                <xsl:call-template name="mkspc">
                    <xsl:with-param name="spc" select="10- string-length(@value)"/>
                </xsl:call-template>
            </xsl:variable>
            <xsl:value-of select="concat('//', $c, ' ...', $lf)"/>
            <xsl:value-of select="concat($in,$c,$cr)"/>
        </xsl:for-each> 
        <xsl:value-of select="concat(')',$cr,$rb,$cr)"/>
        <xsl:value-of select="concat('var Classifications = [...]string{',$cr)"/>
        <xsl:for-each select="xsd:restriction/xsd:enumeration">
            <xsl:variable name="c" select="@value"/>
            <xsl:value-of select="concat($sp,$qt,$c,$qt, $cm,$cr,$rb,$cr)"/>
        </xsl:for-each>       
    </xsl:template>-->

    <xsl:template match="xsd:attributeGroup">
        <xsl:value-of select="concat('//', @name, ' ... ', normalize-space(xsd:annotation/xsd:documentation), $lf)"/>
        <xsl:value-of select="concat('type ', @name, ' struct ', $lb, $cr)"/>
        <xsl:for-each select="xsd:attribute">
            <xsl:variable name="r" select="@ref"/>
            <xsl:variable name="cap" select="concat(translate(substring(@ref, 1, 1), $lowers, $uppers), substring(@ref, 2))"/>
            <xsl:variable name="rspcs">
                <xsl:call-template name="mkspc">
                    <xsl:with-param name="spc" select="number($tab) - string-length($r)"/>
                </xsl:call-template>
            </xsl:variable>
            <xsl:variable name="strtype">
                <xsl:choose>
                    <xsl:when test="/xsd:schema/xsd:attribute[@name = $r]/xsd:simpleType/xsd:restriction/@base = 'xsd:NMTOKENS'">
                        <xsl:text>[]string</xsl:text>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:text>string</xsl:text>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:variable>
            <xsl:variable name="strspc">
                <xsl:call-template name="mkspc">
                    <xsl:with-param name="spc" select="number($tab) - string-length($strtype)-5"/>
                </xsl:call-template>
            </xsl:variable>
            <xsl:value-of select="concat($in, $in, $cap, $rspcs, $strtype, $strspc, $bq, 'xml:', $qt, 'ism:', @ref, $cm, 'attr',$cm,$omitempty,$qt, ' ', $json, $qt, @ref, $cm, $omitempty, $qt, $bq, $cr)"/>
        </xsl:for-each>
        <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>

    <xsl:template name="dotpaths">
        <xsl:param name="elname"/>
        <xsl:param name="nextname"/>
        <xsl:variable name="path" select="concat($elname, '.', $nextname)"/>
        <xsl:variable name="tpe" select="/xsd:schema/xsd:element[@name = $nextname]/@type"/>
        <xsl:choose>
            <xsl:when test="/xsd:schema/xsd:complexType[@name = $tpe]//xsd:element[@ref]">
                <xsl:for-each select="/xsd:schema/xsd:complexType[@name = $tpe]//xsd:element[@ref]">
                    <xsl:variable name="nr" select="@ref"/>
                    <xsl:call-template name="dotpaths">
                        <xsl:with-param name="elname" select="$path"/>
                        <xsl:with-param name="nextname" select="$nr"/>
                    </xsl:call-template>
                </xsl:for-each>
            </xsl:when>
            <xsl:otherwise>
                <xsl:element name="p">
                    <xsl:value-of select="$path"/>
                </xsl:element>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="lastel">
        <xsl:param name="str"/>
        <xsl:choose>
            <xsl:when test="contains($str, '.')">
                <xsl:call-template name="lastel">
                    <xsl:with-param name="str" select="substring-after($str, '.')"/>
                </xsl:call-template>
            </xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="$str"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

</xsl:stylesheet>
