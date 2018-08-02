<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <!-- 
    input: /iepd/xml/xsd/iep.xsd
    output:/iepd/src/golang/xsd/xsd-test.go
   -->

    <xsl:param name="TestData" select="'../instance/test_instance.xml'"/>

    <xsl:variable name="package" select="'seva'"/>
    <xsl:variable name="rootname" select="'SoftwareEvidenceArchive'"/>

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
    <xsl:variable name="tab" select="40"/>
    <xsl:variable name="json" select="' json:'"/>
    <xsl:variable name="omitempty" select="'omitempty'"/>
    <xsl:variable name="roottype" select="xs:schema/xs:element[@name = $rootname]/@type"/>
    <xsl:variable name="ns" select="xs:schema/@targetNamespace"/>

    <xsl:template match="/">
        <xsl:call-template name="maketests"/>
    </xsl:template>

    <xsl:template match="*">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    <xsl:template match="*" mode="def">
        <xsl:apply-templates select="*" mode="def"/>
    </xsl:template>
    <xsl:template match="text()"/>

    <xsl:template name="maketests">
        <xsl:value-of select="concat('package ', 'main', $cr, $cr)"/>
        <xsl:value-of select="concat('import (', $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'encoding/xml', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'fmt', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'io/ioutil', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'testing', $qt, $cr, $cr, $in, '. ', $qt, 'github.com/franela/goblin', $qt, $cr)"/>
        <xsl:value-of select="concat($in, '. ', $qt, 'github.com/onsi/gomega', $qt, $cr)"/>
        <xsl:value-of select="concat(')', $cr, $cr)"/>

        <xsl:value-of select="concat('var testinstances = map[string]string{', $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'test_data.xml', $qt, ':', '               ', $qt, 'xml/test_data.xml', $qt, $cm, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'test_instance.xml', $qt, ':', '      ', $qt, 'xml/test_instance.xml', $qt, $cm, $cr, $rb, $cr)"/>

        <xsl:value-of select="concat('func TestSeva(t *testing.T) {', $cr, $in, 'g := Goblin(t)', $cr)"/>
        <xsl:value-of select="concat($in, 'RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })', $cr, $cr)"/>
        <xsl:value-of select="concat($in, 'xf, ferr := ioutil.ReadFile(testinstances[', $qt, 'test_instance.xml', $qt, '])', $cr)"/>
        <xsl:value-of select="concat($in, 'if ferr != nil {', $cr)"/>
        <xsl:value-of select="concat($in, $in, 'fmt.Printf(ferr.Error())', $cr, $in, '}', $cr)"/>
        <xsl:value-of select="concat($in, 'var seva = NewSoftwareEvidenceArchive()', $cr)"/>
        <xsl:value-of select="concat($in, 'err := xml.Unmarshal([]byte(xf), ', $a, 'seva)', $cr)"/>
        <xsl:value-of select="concat($in, 'if err != nil {', $cr)"/>
        <xsl:value-of select="concat($in, $in, 'fmt.Printf(err.Error())', $cr, $in, '}', $cr)"/>

        <xsl:value-of select="concat($in, 'g.Describe(', $qt, 'SEVA', $qt, $cm, 'func() {', $cr)"/>
        <xsl:for-each select="/xs:schema/xs:complexType[@name = $roottype]//xs:element[@ref]">
            <xsl:variable name="msgtext" select="concat('Must have ', xs:annotation//*/@name)"/>
            <xsl:value-of select="concat($in, $in, 'g.It(', $qt, $msgtext, $qt, $cm, 'func() {', $cr)"/>
            <!--<xsl:value-of select="concat($in, $in, $in, 'seva := New', $rootname, '()', $cr)"/>-->
            <xsl:variable name="rr" select="@ref"/>
            <xsl:variable name="ty" select="/xs:schema/xs:element[@name = $rr]/@type"/>
            <xsl:for-each select="/xs:schema/xs:complexType[@name = $ty]//xs:element[@ref]">
                <xsl:variable name="rrr" select="@ref"/>
                <xsl:variable name="dotpath">
                    <xsl:call-template name="dotpaths">
                        <xsl:with-param name="elname" select="$rr"/>
                        <xsl:with-param name="nextname" select="$rrr"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:variable name="ary">
                    <xsl:if test="@maxOccurs>1">
                        <xsl:text>[0]</xsl:text>
                    </xsl:if>
                </xsl:variable>
                <xsl:for-each select="exsl:node-set($dotpath)/*">
                    <xsl:variable name="el">
                        <xsl:call-template name="lastel">
                            <xsl:with-param name="str" select="."/>
                        </xsl:call-template>
                    </xsl:variable>
                    <xsl:variable name="testval">
                        <xsl:value-of select="document($TestData)//*[name() = $rr][1]//*[name() = $el][1]"/>
                    </xsl:variable>
                    <xsl:value-of select="concat($in, $in, $in, 'Expect(', .,$ary, ').To(Equal(', $qt, $testval, $qt, '))', $cr)"/>
                </xsl:for-each>
            </xsl:for-each>
            <xsl:value-of select="concat($in, $in, '})', $cr)"/>
        </xsl:for-each>
        <xsl:value-of select="concat($in, '})', $cr)"/>
        <xsl:value-of select="concat($cr, '}')"/>
    </xsl:template>

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
        <xsl:param name="array"/>
        <xsl:variable name="ary">
            <xsl:if test="$array > 1">
                <xsl:text>[0]</xsl:text>
            </xsl:if>
        </xsl:variable>
        <xsl:variable name="path" select="concat($elname, '.', $nextname, $ary)"/>
        <xsl:variable name="tpe" select="/xs:schema/xs:element[@name = $nextname]/@type"/>
        <xsl:choose>
            <xsl:when test="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
                <xsl:for-each select="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
                    <xsl:variable name="nr" select="@ref"/>
                    <xsl:call-template name="dotpaths">
                        <xsl:with-param name="elname" select="$path"/>
                        <xsl:with-param name="nextname" select="$nr"/>
                        <xsl:with-param name="array" select="@maxOccurs"/>
                    </xsl:call-template>
                </xsl:for-each>
            </xsl:when>
            <xsl:otherwise>
                <xsl:element name="p">
                    <xsl:value-of select="concat('seva.', $path)"/>
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
                <xsl:choose>
                    <xsl:when test="contains($str, '[0]')">
                        <xsl:value-of select="substring-before($str,'[0]')"/>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:value-of select="$str"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

</xsl:stylesheet>
