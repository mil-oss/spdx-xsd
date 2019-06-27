<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>
    
   <!-- 
    input: /iepd/xml/xsd/iep.xsd
    output:/iepd/src/golang/struct/xsd-struct.go
   -->

    <xsl:variable name="a">
        <xsl:text>&amp;</xsl:text>
    </xsl:variable>
    <xsl:variable name="as">
        <xsl:text>*</xsl:text>
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
        <xsl:text>&#13;</xsl:text>
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
    <xsl:variable name="tab" select="40"/>
    <xsl:variable name="json" select="' json:'"/>
    <xsl:variable name="omitempty" select="'omitempty'"/>
    <xsl:variable name="rootname" select="'SoftwareEvidenceArchive'"/>
    <xsl:variable name="roottype" select="xs:schema/xs:element[@name = $rootname]/@type"/>
    <xsl:variable name="ns" select="xs:schema/@targetNamespace"/>

    <xsl:template match="/">
        <xsl:value-of select="concat('package main', $cr, $cr)"/>
        <xsl:value-of select="concat('import ', $qt, 'encoding/xml', $qt, $cr, $cr)"/>
        <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]" mode="func"/>
        <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]"/>
        <xsl:apply-templates select="xs:schema/xs:element[not(@name = $rootname)]">
            <xsl:sort select="@name"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template match="xs:schema/xs:element">
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:variable name="xspcs">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length('XMLName')"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="tspcs">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length('xml.Name')"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="att1spc">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length('AttrXmlnsXsi')"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="att2spc">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length('AttrXmlns')"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="strspc">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length('string')"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:if test="//xs:complexType[@name = $t]//xs:element[@ref]">
            <xsl:value-of select="concat('//', $n, ' ... ', xs:annotation/xs:documentation, $cr)"/>
            <xsl:value-of select="concat('type ', $n, ' struct ', $lb, $cr)"/>
            <xsl:if test="@name = $rootname">
                <xsl:value-of
                    select="concat($in, $in, 'AttrXmlnsXsi', $att1spc, 'string', $strspc, $bq, 'xml:', $qt, 'xmlns:xsi,attr', $qt, $json, $qt, 'AttrXmlnsXsi', $cm, $omitempty, $qt, $bq, $cr)"/>
                <xsl:value-of select="concat($in, $in, 'AttrXmlns', $att2spc, 'string', $strspc, $bq, 'xml:', $qt, 'xmlns,attr', $qt, $json, $qt, 'AttrXmlns', $cm, $omitempty, $qt, $bq, $cr)"/>
            </xsl:if>
            <xsl:apply-templates select="//xs:complexType[@name = $t]//xs:element[@ref]"/>
            <xsl:value-of select="concat($in, $in, 'XMLName', $xspcs, 'xml.Name', $tspcs, $bq, 'xml:', $qt, $n, $cm, $omitempty, $qt, ' ', $json, $qt, $n, $cm,$omitempty, $qt, $bq, $cr)"/>
            <xsl:value-of select="concat($rb, $cr)"/>
        </xsl:if>
    </xsl:template>

    <!--Processes Ref Elements-->
    <xsl:template match="xs:sequence/xs:element[@ref]">
        <xsl:variable name="r" select="@ref"/>
        <xsl:variable name="t" select="/xs:schema/xs:element[@name = $r]/@type"/>
        <xsl:variable name="typ">
            <xsl:choose>
                <xsl:when test="@maxOccurs>1">
                    <xsl:text>[]</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text></xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="ptr">
            <xsl:choose>
                <xsl:when test="@maxOccurs>1">
                    <xsl:text></xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>*</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="dt">
            <xsl:choose>
                <xsl:when test="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
                    <xsl:value-of select="concat($typ,$ptr,$r)"/>
                    <!--<xsl:value-of select="$r"/>-->
                </xsl:when>
                <xsl:when test="@maxOccurs>1">
                    <xsl:text>[]string</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>string</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="rspcs">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length($r)"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="vspcs">
            <xsl:call-template name="mkspc">
                <xsl:with-param name="spc" select="number($tab) - string-length($dt)"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:value-of select="concat($in, $in, $r, $rspcs, $dt, $vspcs, $bq, 'xml:', $qt, @ref, $cm, $omitempty, $qt, ' ', $json, $qt, @ref,$typ,$cm,$omitempty, $qt, $bq, $cr)"/>
    </xsl:template>

    <xsl:template match="xs:element[@name]" mode="func">
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:if test="//xs:complexType[@name = $t]//xs:element[@ref]">
            <xsl:value-of select="concat('//NewSoftwareEvidenceArchive ...', $cr)"/>
            <xsl:value-of select="concat('func ', 'NewSoftwareEvidenceArchive', '()', $as, $n, $lb, $cr)"/>
            <xsl:value-of select="concat($in, 'return ', $a,$n, $lb, $cr)"/>
            <xsl:if test="@name = $rootname">
                <xsl:value-of select="concat($in, $in, '// Required for the proper namespacing', $n, $cr)"/>
                <xsl:value-of select="concat($in, $in, 'AttrXmlnsXsi', ':', $qt, 'http://www.w3.org/2001/XMLSchema-instance', $qt, $cm, $cr)"/>
                <xsl:value-of select="concat($in, $in, 'AttrXmlns', ':', $qt, 'urn:security::1.0', $qt, $cm, $cr)"/>
            </xsl:if>
            <xsl:for-each select="//xs:complexType[@name = $t]//xs:element[@ref]">
                <xsl:variable name="r" select="@ref"/>
                <xsl:variable name="fspcs">
                    <xsl:call-template name="mkspc">
                        <xsl:with-param name="spc" select="number($tab) - string-length(@ref)"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:value-of select="concat($in, $in, $r, ': ', $a,$r, $lb)"/>
                <xsl:variable name="et" select="//xs:element[@name = $r]/@type"/>
                <xsl:variable name="complexKids">
                    <xsl:for-each select="//xs:complexType[@name = $et]//xs:element[@ref]">
                        <xsl:variable name="rr" select="@ref"/>
                        <xsl:variable name="ett" select="//xs:element[@name = $rr]/@type"/>
                        <xsl:if test="/xs:schema/xs:complexType[@name = $ett]//xs:element[@ref]">
                            <ctype name="{$rr}"/>
                        </xsl:if>
                    </xsl:for-each>
                </xsl:variable>
                <xsl:choose>
                    <xsl:when test="exsl:node-set($complexKids)/*">
                        <xsl:value-of select="concat($cr, $in, $in, $in)"/>
                        <xsl:for-each select="exsl:node-set($complexKids)/*">
                            <xsl:value-of select="concat(@name, ':',$a,@name, '{}')"/>
                            <xsl:if test="following-sibling::*">
                                <xsl:value-of select="concat($cm, $cr, $in, $in, $in)"/>
                            </xsl:if>
                        </xsl:for-each>
                        <xsl:value-of select="concat($cm, $cr, $in, $in, $rb, $cm, $cr)"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:value-of select="concat($rb, $cm, $cr)"/>
                    </xsl:otherwise>
                </xsl:choose>
                <xsl:if test="following-sibling::*[@name]">
                    <xsl:value-of select="concat($cm, $cr)"/>
                </xsl:if>
            </xsl:for-each>
            <xsl:value-of select="concat($in, $rb, $cr)"/>
            <xsl:value-of select="concat($rb, $cr)"/>
        </xsl:if>
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

    <xsl:template name="dotpaths">
        <xsl:param name="elname"/>
        <xsl:param name="nextname"/>
        <xsl:variable name="path" select="concat($elname, '.', $nextname)"/>
        <xsl:variable name="tpe" select="/xs:schema/xs:element[@name = $nextname]/@type"/>
        <xsl:choose>
            <xsl:when test="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
                <xsl:for-each select="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
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
