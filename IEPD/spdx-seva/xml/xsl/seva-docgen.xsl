<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="no" xml:space="preserve"/>

    <!-- 
    input:  /iepd/xml/xsd/iep.xsd
    output: /iepd/xml/instance/test_instance.xml
   -->

    <xsl:param name="sevaref" select="'../xsd/ext/seva/xml/xsd/ref.xsd'"/>
    <xsl:param name="Root" select="'SoftwareEvidenceArchiveType'"/>
    <xsl:param name="Out" select="'../../../../doc/SEvA_XSD.html'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <xsl:result-document href="{$Out}">
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
                    <pre><div style="text-align:center">SEvA XML</div></pre>
                    <xsl:apply-templates select="xs:schema/xs:complexType[@name = $Root]" mode="root"/>
                </body>
            </html>
        </xsl:result-document>
    </xsl:template>

    <xsl:template match="*" mode="root">
        <xsl:variable name="n" select="xs:annotation/xs:appinfo/*/@name"/>
        <xsl:variable name="d" select="normalize-space(xs:annotation/xs:documentation)"/>
        <pre><div><b><span class="lvl1"><xsl:value-of select="$n"/></span></b></div><div><span class="lvl2"><xsl:value-of select="normalize-space($d)"/></span></div></pre>
        <xsl:apply-templates select="xs:complexContent/xs:extension/xs:sequence/xs:element" mode="ref">
            <xsl:with-param name="lvl" select="2"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="*" mode="ref">
        <xsl:param name="lvl"/>
        <xsl:if test="not(ends-with(@ref, 'AugmentationPoint'))">
            <xsl:variable name="r" select="@ref"/>
            <xsl:variable name="e" select="//xs:schema/xs:element[@name = $r]"/>
            <xsl:variable name="d" select="$e/xs:annotation/xs:documentation"/>
            <xsl:variable name="n" select="xs:annotation/xs:appinfo/*/@name"/>
            <xsl:variable name="t" select="$e/@type"/>
            <xsl:variable name="tn" select="//xs:schema/*[@name = $t]"/>
            <pre><div class="{concat('lvl',$lvl)}"><b><xsl:value-of select="$n"/></b></div><div><span class="{concat('lvl',$lvl+1)}"><xsl:value-of select="normalize-space($d)"/></span></div></pre>
            <xsl:apply-templates select="$tn/*/*/xs:sequence/xs:element" mode="ref">
                <xsl:with-param name="lvl" select="$lvl + 1"/>
            </xsl:apply-templates>
        </xsl:if>
    </xsl:template>

</xsl:stylesheet>
