<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <!-- 
    input: /iepd/xml/xsd/iep.xsd
    output:/iepd/src/golang/struct/xsd-struct.go
   -->

    <!-- <xsl:template match="/">
        <xsl:call-template name="makego">
            <xsl:with-param name="rootname" select="'License'"/>
        </xsl:call-template>
    </xsl:template>-->

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
    <xsl:variable name="tab">
        <xsl:text>&#09;</xsl:text>
    </xsl:variable>
    <xsl:variable name="json" select="' json:'"/>
    <xsl:variable name="omitempty" select="'omitempty'"/>

    <xsl:template match="xs:element[@name]" mode="func">
        <xsl:param name="rootname"/>
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:if test="//xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
            <xsl:variable name="b" select="/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
            <xsl:value-of select="concat('//New', $rootname, ' ...', $cr)"/>
            <xsl:value-of select="concat('func ', 'New', $rootname, '() ', $as, $n, $lb, $cr)"/>
            <xsl:value-of select="concat($in, 'return ', $a, $n, $lb, $cr)"/>
            <xsl:if test="@name = $rootname">
                <xsl:value-of select="concat($in, $in, '// Required for the proper namespacing', $cr)"/>
                <xsl:value-of select="concat($in, $in, 'AttrXmlnsXsi', ':', $qt, 'http://www.w3.org/2001/XMLSchema-instance', $qt, $cm, $cr)"/>
                <xsl:value-of select="concat($tab, 'XsiType', $tab,  'string', $tab, $bq, 'xml:', $qt, 'http://www.w3.org/2001/XMLSchema-instance type,attr,omitempty', $qt, $bq, $cr)"/>
                <xsl:value-of select="concat($in, $in, 'AttrXmlns', ':', $qt, 'spdx:xsd::1.0', $qt, $cm, $cr)"/>
            </xsl:if>
            <!--<xsl:apply-templates select="/xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="makevar"/>
            <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]" mode="makevar"/>-->
            <xsl:value-of select="concat($in, $rb, $cr)"/>
            <xsl:value-of select="concat($rb, $cr)"/>
        </xsl:if>
    </xsl:template>

    <xsl:template match="xs:schema/xs:element">
        <xsl:param name="rootname"/>
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:choose>
            <!--Recursion Check-->
            <xsl:when test="concat($rootname,'Type')=$t and $n!=$rootname"/>
            <xsl:when test="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
                <xsl:variable name="b" select="/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
                <xsl:value-of select="concat('//', $n, ' ... ', substring-before(xs:annotation/xs:documentation, '.'), $cr)"/>
                <xsl:value-of select="concat('type ', $n, ' struct ', $lb, $cr)"/>
                <xsl:if test="@name = $rootname">
                    <xsl:value-of select="concat($in, $in, 'AttrXmlnsXsi', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'xmlns:xsi,attr', $qt, $json, $qt, 'AttrXmlnsXsi', $cm, $omitempty, $qt, $bq, $cr)"/>
                    <xsl:value-of select="concat($tab, 'AttrXmlns', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'xmlns,attr', $qt, $json, $qt, 'AttrXmlns', $cm, $omitempty, $qt, $bq, $cr)"/>
                </xsl:if>
                <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="makevar">
                    <xsl:with-param name="rootname" select="$rootname"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]" mode="makevar">
                    <xsl:with-param name="rootname" select="$rootname"/>
                </xsl:apply-templates>
                <xsl:value-of select="concat($tab, 'XMLName', $tab, 'xml.Name', $tab, $tab, $bq, 'xml:', $qt, $n, $cm, $omitempty, $qt, ' ', $json, $qt, $n, $cm, $omitempty, $qt, $bq, $cr)"/>
                <xsl:value-of select="concat($rb, $cr)"/>
            </xsl:when>
        </xsl:choose>
    </xsl:template>
    
    <xsl:template match="*" mode="makevar">
        <xsl:param name="rootname"/>
        <xsl:variable name="r" select="@ref"/>       
        <xsl:variable name="t" select="/xs:schema/xs:element[@name = $r]/@type"/>
        <xsl:variable name="b" select="/xs:schema/xs:complexType[@name = $t]//@base"/>
        <xsl:variable name="ary">
            <xsl:if test="@maxOccurs &gt;1">
                <xsl:text>[]</xsl:text>
            </xsl:if>
        </xsl:variable>
        <xsl:variable name="ptr">
            <xsl:choose>
                <xsl:when test="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
                    <xsl:text>*</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="dt">
            <xsl:choose>
                <xsl:when test="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
                    <xsl:value-of select="concat($ary, $ptr, $r)"/>
                </xsl:when>
                <xsl:when test="/xs:schema/xs:complexType[@name = $t]//@base='xs:boolean'">
                    <xsl:value-of select="concat($ary, 'bool')"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="concat($ary, 'string')"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:value-of select="concat($tab, $r, $tab, $dt, $tab, $tab, $bq, 'xml:', $qt, @ref, $cm, $omitempty, $qt, ' ', $json, $qt, @ref, $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>

    <xsl:template name="maketests">
        <xsl:param name="testdata"/>
        <xsl:param name="pckgname"/>
        <xsl:variable name="rootname" select="//xs:schema/xs:annotation/xs:appinfo/*/@name"/>
        <xsl:variable name="roottype" select="//xs:schema/xs:annotation/xs:appinfo/*/@type"/>
        <xsl:call-template name="teststart">
            <xsl:with-param name="appname" select="$rootname"/>
            <xsl:with-param name="pckgname" select="$pckgname"/>
        </xsl:call-template>
        <xsl:variable name="b" select="//xs:schema/xs:complexType[@name = $roottype]//@base"/>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="maketest">
            <xsl:with-param name="rootname" select="$rootname"/>
            <xsl:with-param name="testdata" select="$testdata"/>
        </xsl:apply-templates>
        <xsl:apply-templates select="//xs:schema/xs:complexType[@name = $roottype]//xs:element[@ref]" mode="maketest">
            <xsl:with-param name="rootname" select="$rootname"/>
            <xsl:with-param name="testdata" select="$testdata"/>
        </xsl:apply-templates>
        <xsl:value-of select="concat($in, '})', $cr)"/>
        <xsl:value-of select="concat($cr, '}')"/>
    </xsl:template>

    <xsl:template match="*" mode="maketest">
        <xsl:param name="rootname"/>
        <xsl:param name="testdata"/>
        <xsl:variable name="r" select="@ref"/>
        <xsl:if test="exsl:node-set($testdata//*[name() = $r])">
            <xsl:variable name="node" select="/xs:schema/xs:element[@name = $r]"/>
            <xsl:variable name="msgtext" select="concat('Must have ', $node/xs:annotation/xs:appinfo/*/@name)"/>
            <xsl:value-of select="concat($tab, 'g.It(', $qt, $msgtext, $qt, $cm, 'func() {', $cr)"/>
            <xsl:variable name="ty" select="$node/@type"/>
            <xsl:variable name="b" select="//xs:schema/xs:complexType[@name = $ty]//@base"/>
            <xsl:variable name="ary">
                <xsl:if test="@maxOccurs > 1">
                    <xsl:text>[0]</xsl:text>
                </xsl:if>
            </xsl:variable>
            <xsl:if test="not(//xs:schema/xs:complexType[@name = $ty]//xs:element[@ref])">
                <xsl:variable name="testval">
                    <!--<xsl:value-of select="concat($r,', ',$rr)"/>-->
                    <xsl:choose>
                        <xsl:when test="string-length($testdata//*[name() = $rootname][1]//*[name() = $r][1])&gt;0">
                            <xsl:value-of select="$testdata//*[name() = $rootname][1]//*[name() = $r][1]"/>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:text>Test string one</xsl:text>
                        </xsl:otherwise>
                    </xsl:choose>
                </xsl:variable>
                <xsl:value-of select="concat($tab, $tab, 'Expect(', 'spdx.', $r, $ary, ').To(Equal(', $qt, $testval, $qt, '))', $cr)"/>
            </xsl:if>
            <xsl:for-each select="//xs:schema/xs:complexType[@name = $ty]//xs:element[@ref]">
                <xsl:variable name="rr" select="@ref"/>
                <xsl:variable name="dotpath">
                    <xsl:call-template name="dotpaths">
                        <xsl:with-param name="elname" select="$r"/>
                        <xsl:with-param name="nextname" select="$rr"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:variable name="ary1">
                    <xsl:if test="@maxOccurs > 1">
                        <xsl:text>[0]</xsl:text>
                    </xsl:if>
                </xsl:variable>
                <xsl:for-each select="exsl:node-set($dotpath)/*">
                    <xsl:variable name="el">
                        <xsl:call-template name="lastel">
                            <xsl:with-param name="str" select="."/>
                        </xsl:call-template>
                    </xsl:variable>
                    <xsl:variable name="testval1">
                        <!--<xsl:value-of select="concat($r,', ',$rr)"/>-->
                        <xsl:value-of select="$testdata//*[name() = $r][1]//*[name() = $el][1]"/>
                    </xsl:variable>
                    <xsl:variable name="testval2">
                        <xsl:choose>
                            <xsl:when test="string-length($testval1)&gt;0">
                                <xsl:value-of select="$testval1"/>
                            </xsl:when>
                            <xsl:otherwise>
                                <xsl:text>Test string one</xsl:text>
                            </xsl:otherwise>
                        </xsl:choose>
                    </xsl:variable>
                    <xsl:value-of select="concat($in, $in, $in, 'Expect(', 'spdx.', ., $ary1, ').To(Equal(', $qt, $testval1, $qt, '))', $cr)"/>
                </xsl:for-each>
            </xsl:for-each>
            <xsl:value-of select="concat($in, $in, '})', $cr)"/>
        </xsl:if>
    </xsl:template>
    <!--Processes Ref Elements-->
    <xsl:template match="xs:sequence/xs:element[@ref]">
        <xsl:variable name="r" select="@ref"/>
        <xsl:variable name="t" select="/xs:schema/xs:element[@name = $r]/@type"/>
        <xsl:variable name="ary">
            <xsl:choose>
                <xsl:when test="@maxOccurs">
                    <xsl:text>[]</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="ptr">
            <xsl:choose>
                <xsl:when test="ancestor::xs:complexType">
                    <xsl:text/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>*</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:variable name="dt" select="concat($ary, '*', $r)"/>
        <xsl:value-of select="concat($tab, $r, $tab, $dt, $tab, $tab, $bq, 'xml:', $qt, @ref, $cm, $omitempty, $qt, ' ', $json, $qt, @ref, $ary, $cm, $omitempty, $qt, $bq, $cr)"/>
    </xsl:template>
    
    <xsl:template match="*">
        <xsl:param name="rootname"/>
        <xsl:apply-templates select="*">
            <xsl:with-param name="rootname" select="$rootname"/>
        </xsl:apply-templates>
    </xsl:template>
   
    <xsl:template match="*" mode="def">
        <xsl:param name="rootname"/>
        <xsl:apply-templates select="*" mode="def">
            <xsl:with-param name="rootname" select="$rootname"/>
        </xsl:apply-templates>
    </xsl:template>

    <xsl:template name="teststart">
        <xsl:param name="appname"/>
        <xsl:param name="pckgname"/>
        <xsl:value-of select="concat('package ', $pckgname, $cr, $cr)"/>
        <xsl:value-of select="concat('import (', $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'encoding/xml', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'fmt', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'io/ioutil', $qt, $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'testing', $qt, $cr, $cr, $in, '. ', $qt, 'github.com/franela/goblin', $qt, $cr)"/>
        <xsl:value-of select="concat($in, '. ', $qt, 'github.com/onsi/gomega', $qt, $cr)"/>
        <xsl:value-of select="concat(')', $cr, $cr)"/>
        <xsl:value-of select="concat('var testinstances = map[string]string{', $cr)"/>
        <xsl:value-of select="concat($in, $qt, 'test_instance.xml', $qt, ':', '      ', $qt, 'xml/test_instance.xml', $qt, $cm, $cr, $rb, $cr)"/>
        <xsl:value-of select="concat('func Test', $appname, '(t *testing.T) {', $cr, $in, 'g := Goblin(t)', $cr)"/>
        <xsl:value-of select="concat($in, 'RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })', $cr, $cr)"/>
        <xsl:value-of select="concat($in, 'xf, ferr := ioutil.ReadFile(testinstances[', $qt, 'test_instance.xml', $qt, '])', $cr)"/>
        <xsl:value-of select="concat($in, 'if ferr != nil {', $cr)"/>
        <xsl:value-of select="concat($in, $in, 'fmt.Printf(ferr.Error())', $cr, $in, '}', $cr)"/>
        <xsl:value-of select="concat($in, 'var spdx = New', $appname, '()', $cr)"/>
        <xsl:value-of select="concat($in, 'err := xml.Unmarshal([]byte(xf), ', $a, 'spdx)', $cr)"/>
        <xsl:value-of select="concat($in, 'if err != nil {', $cr)"/>
        <xsl:value-of select="concat($in, $in, 'fmt.Printf(err.Error())', $cr, $in, '}', $cr)"/>
        <xsl:value-of select="concat($in, 'g.Describe(', $qt, $appname, $qt, $cm, 'func() {', $cr)"/>
    </xsl:template>

    <xsl:template match="text()"/>

    <xsl:template name="dotpaths">
        <xsl:param name="elname"/>
        <xsl:param name="nextname"/>
        <xsl:variable name="path" select="concat($elname, '.', $nextname)"/>
        <xsl:variable name="tpe" select="/xs:schema/xs:element[@name = $nextname]/@type"/>
        <xsl:choose>
            <xsl:when test="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
                <xsl:for-each select="/xs:schema/xs:complexType[@name = $tpe]//xs:element[@ref]">
                    <xsl:variable name="nr" select="@ref"/>
                    <xsl:choose>
                        <xsl:when test="contains($path, $nr)"/>
                        <xsl:otherwise>
                            <xsl:call-template name="dotpaths">
                                <xsl:with-param name="elname" select="$path"/>
                                <xsl:with-param name="nextname" select="$nr"/>
                            </xsl:call-template>
                        </xsl:otherwise>
                    </xsl:choose>
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
