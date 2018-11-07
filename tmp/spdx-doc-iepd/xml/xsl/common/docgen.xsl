<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="no" xml:space="preserve"/>

    <!-- 
    input:  /iepd/xml/xsd/iep.xsd
    output: /iepd/xml/instance/test_instance.xml
   -->

    <xsl:variable name="lcase" select="'abcdefghijklmnopqrstuvwxyz'"/>
    <xsl:variable name="_2sp">
        <xsl:text>&#32;&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="Lvl2">
        <xsl:text/>
    </xsl:variable>
    <xsl:variable name="Lvl3">
        <xsl:text>&#32;&#32;&#32;&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="Lvl4">
        <xsl:text>&#32;&#32;&#32;&#32;&#32;&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="Lvl5">
        <xsl:text>&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="Lvl6">
        <xsl:text>&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;</xsl:text>
    </xsl:variable>
    <xsl:variable name="Lvl7">
        <xsl:text>&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;&#32;</xsl:text>
    </xsl:variable>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <!--<xsl:result-document href="{$Out}">-->
        <xsl:variable name="t" select="//xs:schema/xs:annotation/xs:appinfo//@type"/>
        <html>
            <head>
                <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
                <style type="text/css">
                    body {
                        width: 7in;
                        padding: .75in;
                        font-family: "Courier New", Courier, monospace;
                        font-size: .9em;
                        border: thin black solid;
                    }
                    pre {
                        white-space: pre-wrap; /* css-3 */
                        white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
                        white-space: -pre-wrap; /* Opera 4-6 */
                        white-space: -o-pre-wrap; /* Opera 7 */
                        word-wrap: break-word;
                        font-family: "Courier New", Courier, monospace;
                        font-size: .9em;
                    }
                    
                    .tab {
                        padding-right: .2in;
                    }
                    .lvl2 {
                        display: inline;
                    }
                    .lvl3 {
                        display: inline;
                        padding-left: .3in;
                    }
                    .lvl4 {
                        display: inline;
                        padding-left: .6in;
                    }
                    .lvl5 {
                        display: inline;
                        padding-left: 1in;
                    }
                    .lvl6 {
                        display: inline;
                        padding-left: 1.4in;
                    }
                    .lvl7 {
                        display: inline;
                        padding-left: 1.8in;
                    }
                    .lvl8 {
                        display: inline;
                        padding-left: 2.2in;
                    }
                    .lvl9 {
                        display: inline;
                        padding-left: 2.6in;
                    }
                    .uline {
                        text-decoration: underline;
                    }
                    .ctr_title {
                        text-align: center;
                    }
                    .center
                    {
                        margin-left: auto;
                        margin-right: auto;
                        width: 100%;
                    }
                    P.pagebreak
                    {
                        page-break-before: always
                    }</style>
            </head>
            <body>
                <pre><div style="text-align:center">SPDX XML</div></pre>
                <xsl:apply-templates select="xs:schema/xs:complexType[@name = $t]" mode="root"/>
            </body>
        </html>
        <!--</xsl:result-document>-->
    </xsl:template>

    <xsl:template match="*" mode="root">
        <xsl:variable name="n" select="xs:annotation/xs:appinfo/*/@name"/>
        <xsl:variable name="d" select="normalize-space(xs:annotation/xs:documentation)"/>
        <pre>
            <div style="text-align:center"><b><span><xsl:value-of select="$n"/></span></b></div>
            <div style="text-align:center"><span><xsl:value-of select="$d"/></span></div>
        </pre>
        <xsl:apply-templates select=".//xs:sequence/xs:element" mode="ref">
            <xsl:with-param name="lvl" select="2"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="*" mode="ref">
        <xsl:param name="lvl"/>
        <xsl:variable name="seq" select="count(preceding-sibling::*) + 1"/>
        <xsl:if test="not(substring(@ref, string-length(@ref) - string-length('AugmentationPoint') + 1) = 'AugmentationPoint')">
            <xsl:variable name="r" select="@ref"/>
            <xsl:variable name="e" select="//xs:schema/xs:element[@name = $r]"/>
            <xsl:variable name="d" select="$e/xs:annotation/xs:documentation"/>
            <xsl:variable name="n" select="$e/xs:annotation/xs:appinfo/*/@name"/>
            <xsl:variable name="xn">
                <xsl:apply-templates select="$e/xs:annotation/xs:appinfo/*/@xmlname" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="t" select="$e/@type"/>
            <xsl:variable name="ctyp" select="//xs:schema/*[@name = $t]"/>
            <xsl:variable name="b">
                <xsl:if test="$ctyp//@base != 'structures:ObjectType'">
                    <xsl:value-of select="$ctyp//@base"/>
                </xsl:if>
            </xsl:variable>
            <xsl:variable name="etyp" select="//xs:schema/*[@name = $b]"/>
            <xsl:variable name="bse">
                <xsl:value-of select="$etyp//@base"/>
            </xsl:variable>
            <xsl:variable name="spacing">
                <xsl:choose>
                    <xsl:when test="$lvl = 2">
                        <xsl:value-of select="$Lvl2"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 3">
                        <xsl:value-of select="$Lvl3"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 4">
                        <xsl:value-of select="$Lvl4"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 5">
                        <xsl:value-of select="$Lvl5"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 6">
                        <xsl:value-of select="$Lvl6"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 7">
                        <xsl:value-of select="$Lvl7"/>
                    </xsl:when>
                </xsl:choose>
            </xsl:variable>
            <xsl:variable name="bullet">
                <xsl:choose>
                    <xsl:when test="$lvl = 2">
                        <xsl:value-of select="$Lvl2"/>
                        <xsl:value-of select="concat($seq, '.', $_2sp)"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 3">
                        <xsl:value-of select="$Lvl3"/>
                        <xsl:value-of select="concat(substring($lcase, $seq, 1), '.', $_2sp)"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 4">
                        <xsl:value-of select="$Lvl4"/>
                        <xsl:value-of select="concat('(', $seq, ')', $_2sp)"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 5">
                        <xsl:value-of select="$Lvl5"/>
                        <xsl:value-of select="concat('(', substring($lcase, $seq, 1), ')', $_2sp)"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 6">
                        <xsl:value-of select="$Lvl6"/>
                        <u>
                            <xsl:value-of select="$seq"/>
                        </u>
                        <xsl:text>.</xsl:text>
                        <xsl:value-of select="$_2sp"/>
                    </xsl:when>
                    <xsl:when test="$lvl = 7">
                        <xsl:value-of select="$Lvl7"/>
                        <xsl:value-of select="concat(substring($lcase, $seq, 1), '.', $_2sp)"/>
                    </xsl:when>
                </xsl:choose>
            </xsl:variable>
            <xsl:variable name="ty">
                <xsl:apply-templates select="$e/@type" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="minoccur">
                <xsl:apply-templates select="@minOccurs" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="maxoccur">
                <xsl:apply-templates select="@maxOccurs" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="regex">
                <xsl:apply-templates select="$etyp/xs:restriction/xs:pattern/@value" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="maxlen">
                <xsl:apply-templates select="$etyp/xs:restriction/xs:maxLength/@value" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="length">
                <xsl:apply-templates select="$etyp/xs:restriction/xs:length/@value" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="rdf">
                <xsl:apply-templates select="$etyp/xs:annotation/xs:appinfo/*/@rdf" mode="txt"/>
            </xsl:variable>
            <xsl:variable name="enums">
                <xsl:if test="$etyp/xs:restriction/xs:enumeration">
                    <xsl:text>enumerations = </xsl:text>
                </xsl:if>
                <xsl:for-each select="$etyp/xs:restriction/xs:enumeration">
                    <xsl:value-of select="@value"/>
                    <xsl:if test="following-sibling::xs:enumeration">
                        <xsl:text>, </xsl:text>
                    </xsl:if>
                </xsl:for-each>
            </xsl:variable>
            <pre>
                <div><xsl:value-of select="$bullet"/><u><xsl:value-of select="$n"/></u><xsl:value-of select="concat('.', $_2sp)"/><span><xsl:value-of select="normalize-space($d)"/></span></div>
                <xsl:if test="string-length($xn) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$xn"/></div>
                </xsl:if>
                <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$ty"/></div>
                <xsl:if test="string-length($b) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:text>simpletype = </xsl:text><xsl:value-of select="$b"/></div>
                </xsl:if>
                <xsl:if test="string-length($bse) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:text>base = </xsl:text><xsl:value-of select="$bse"/></div>
                 </xsl:if>   
                <xsl:if test="string-length($minoccur) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$minoccur"/></div>
                </xsl:if>
                <xsl:if test="string-length($maxoccur) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$maxoccur"/></div>
                </xsl:if>
                 <xsl:if test="string-length($maxlen) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$maxlen"/></div>
                </xsl:if>
                  <xsl:if test="string-length($length) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$length"/></div>
                </xsl:if>
                <xsl:if test="string-length($regex) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$regex"/></div>
                </xsl:if>
                 <xsl:if test="string-length($rdf) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$rdf"/></div>
                </xsl:if>
                 <xsl:if test="string-length($enums) &gt; 0">
                    <div><xsl:value-of select="concat($spacing, $_2sp, $_2sp)"/><xsl:value-of select="$enums"/></div>
                </xsl:if>
            </pre>
            <xsl:apply-templates select="$etyp//xs:sequence/xs:element" mode="ref">
                <xsl:with-param name="lvl" select="$lvl + 1"/>
            </xsl:apply-templates>
            <xsl:apply-templates select="$ctyp//xs:sequence/xs:element" mode="ref">
                <xsl:with-param name="lvl" select="$lvl + 1"/>
            </xsl:apply-templates>
        </xsl:if>
    </xsl:template>

    <xsl:template match="xs:pattern/@*" mode="txt">
        <xsl:value-of select="concat('pattern', ' = ', ., ' ')"/>
    </xsl:template>

    <xsl:template match="xs:length/@*" mode="txt">
        <xsl:value-of select="concat('length', ' = ', ., ' ')"/>
    </xsl:template>

    <xsl:template match="xs:maxLength/@*" mode="txt">
        <xsl:value-of select="concat('maxLength', ' = ', ., ' ')"/>
    </xsl:template>

    <xsl:template match="@*" mode="txt">
        <xsl:value-of select="concat(name(), ' = ', ., ' ')"/>
    </xsl:template>



</xsl:stylesheet>
