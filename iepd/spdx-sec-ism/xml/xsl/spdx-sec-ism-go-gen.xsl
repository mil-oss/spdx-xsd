<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./common/go-gen.xsl"/>
    
    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-security/xml/xsd/spdx-security-iep.xsd
    output:${pdu}/spdx-xsd/src/golang/spdx-security/spdx-security-struct.go
   -->

    <xsl:template match="/">
        <xsl:variable name="rootname" select="xs:schema/xs:annotation/xs:appinfo/*/@name"/>
        <xsl:value-of select="concat('package spdxsecism', $cr, $cr)"/>
        <xsl:value-of select="concat('import ', $qt, 'encoding/xml', $qt, $cr, $cr)"/>
        <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]" mode="ismfunc">
            <xsl:with-param name="rootname" select="$rootname"/>
        </xsl:apply-templates>
        <xsl:apply-templates select="xs:schema/xs:element[@name = $rootname]">
            <xsl:with-param name="rootname" select="$rootname"/>
        </xsl:apply-templates>
        <xsl:apply-templates select="xs:schema/xs:element[not(@name = $rootname)]">
            <xsl:with-param name="rootname" select="$rootname"/>
            <xsl:sort select="@name"/>
        </xsl:apply-templates>
        <xsl:apply-templates select="xs:schema/xs:attributeGroup">
            <xsl:sort select="@name"/>
        </xsl:apply-templates>
    </xsl:template>
    
    <xsl:template match="xs:element[@name]" mode="ismfunc">
        <xsl:param name="rootname"/>
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:if test="//xs:schema/xs:complexType[@name = $t]//xs:element[@ref]">
            <xsl:variable name="b" select="/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
            <xsl:value-of select="concat('//New', $rootname, ' ...', $cr)"/>
            <xsl:value-of select="concat('func ', 'New', $rootname, '() ', $as, $n, $lb, $cr)"/>
            <xsl:value-of select="concat($tab, 'return ', $a, $n, $lb, $cr)"/>
            <xsl:if test="@name = $rootname">
                <xsl:value-of select="concat($tab,$tab, '// Required for the proper namespacing', $cr)"/>
                <xsl:value-of select="concat($tab,$tab, 'AttrXmlnsXsi', ':', $qt, 'http://www.w3.org/2001/XMLSchema-instance', $qt, $cm, $cr)"/>
                <xsl:value-of select="concat($tab,$tab, 'AttrXmlns', ':', $qt, 'spdx:xsd::1.0', $qt, $cm, $cr)"/>
                <xsl:value-of select="concat($tab,$tab, 'AttrXmlnsIsm',':', $qt, 'urn:us:gov:ic:ism', $qt, $cm, $cr)"/>
            </xsl:if>
            <!--<xsl:apply-templates select="/xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="makevar"/>
            <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]" mode="makevar"/>-->
            <xsl:value-of select="concat($tab, $rb, $cr)"/>
            <xsl:value-of select="concat($rb, $cr)"/>
        </xsl:if>
    </xsl:template>
    
    
<!--    <xsl:template match="xs:schema/xs:element">
        <xsl:param name="rootname"/>
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
                <xsl:variable name="b" select="/xs:schema/xs:complexType[@name = $t]//xs:extension/@base"/>
                <xsl:value-of select="concat('//', $n, ' ... ', substring-before(xs:annotation/xs:documentation, '.'), $cr)"/>
                <xsl:value-of select="concat('type ', $n, ' struct ', $lb, $cr)"/>
                <xsl:if test="@name = $rootname">
                    <xsl:value-of select="concat($in, $in, 'AttrXmlnsXsi', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'xmlns:xsi,attr', $qt, $json, $qt, 'AttrXmlnsXsi', $cm, $omitempty, $qt, $bq, $cr)"/>
                    <xsl:value-of select="concat($tab, 'AttrXmlns', $tab, 'string', $tab, $tab, $bq, 'xml:', $qt, 'xmlns,attr', $qt, $json, $qt, 'AttrXmlns', $cm, $omitempty, $qt, $bq, $cr)"/>
                    <xsl:value-of select="concat($tab, 'AttrXmlnsIsm', $tab, 'string', $tab, $bq, 'xml:', $qt, 'xmlns:ism,attr', $qt, $json, $qt, 'AttrXmlnsIsm', $cm, $omitempty, $qt, $bq, $cr)"/>
                </xsl:if>
                <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $b]//xs:element[@ref]" mode="makevar">
                    <xsl:with-param name="rootname" select="$rootname"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="/xs:schema/xs:complexType[@name = $t]//xs:element[@ref]" mode="makevar">
                    <xsl:with-param name="rootname" select="$rootname"/>
                </xsl:apply-templates>
                <!-\-<xsl:apply-templates select="//xs:complexType[@name = $t]//xs:element[@ref]"/>-\->
                <xsl:if test="not(//xs:complexType[@name = $t]//xs:element[@ref])">
                    <xsl:value-of select="concat($tab, 'Value', $tab, 'string', $tab, $bq, 'xml:', $qt, $cm, 'chardata', $qt, ' ', $json, $qt, 'Value', $cm, $omitempty, $qt, $bq, $cr)"/>
                </xsl:if>
                <xsl:value-of select="concat($tab, 'SecurityAttributesOptionGroup', $tab, $tab, '//SecurityAttributesOptionGroup', $cr)"/>
                <xsl:value-of select="concat($tab, 'XMLName', $tab, 'xml.Name', $tab, $tab, $bq, 'xml:', $qt, $n, $cm, $omitempty, $qt, ' ', $json, $qt, $n, $cm, $omitempty, $qt, $bq, $cr)"/>
                <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>
    -->
  <xsl:template match="xs:schema/xs:element">
        <xsl:param name="rootname"/>
        <xsl:variable name="n" select="@name"/>
        <xsl:variable name="t" select="@type"/>
        <xsl:value-of select="concat('//', $n, ' ... ', xs:annotation/xs:documentation, $cr)"/>
        <xsl:value-of select="concat('type ', $n, ' struct ', $lb, $cr)"/>
        <xsl:if test="@name = $rootname">
            <xsl:value-of select="concat($tab, 'AttrXmlnsXsi', $tab, 'string', $tab, $bq, 'xml:', $qt, 'xmlns:xsi,attr', $qt, $json, $qt, 'xmlns xsi,attr', $cm, $omitempty, $qt, $bq, $cr)"/>
            <xsl:value-of select="concat($tab, 'AttrXmlns', $tab, 'string', $tab, $bq, 'xml:', $qt, 'xmlns,attr', $qt, $json, $qt, 'AttrXmlns', $cm, $omitempty, $qt, $bq, $cr)"/>
            <xsl:value-of select="concat($tab, 'AttrXmlnsIsm', $tab, 'string', $tab, $bq, 'xml:', $qt, 'xmlns:ism,attr', $qt, $json, $qt, 'AttrXmlnsIsm', $cm, $omitempty, $qt, $bq, $cr)"/>
        </xsl:if>
        <xsl:apply-templates select="//xs:complexType[@name = $t]//xs:element[@ref]"/>
        <xsl:if test="not(//xs:complexType[@name = $t]//xs:element[@ref])">
            <xsl:value-of select="concat($tab, 'Value', $tab, 'string', $tab, $bq, 'xml:', $qt, $cm, 'chardata', $qt, ' ', $json, $qt, 'Value', $cm, $omitempty, $qt, $bq, $cr)"/>
        </xsl:if>
        <xsl:value-of select="concat($tab, 'SecurityAttributesOptionGroup', $tab, $tab, '//SecurityAttributesOptionGroup', $cr)"/>
        <xsl:value-of select="concat($tab, 'XMLName', $tab, 'xml.Name', $tab, $bq, 'xml:', $qt, $n, $cm, $omitempty, $qt, ' ', $json, $qt, $n, $cm, $omitempty, $qt, $bq, $cr)"/>
        <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>
    
    <xsl:template match="xs:attributeGroup">
        <xsl:value-of select="concat('//', @name, ' ... ', normalize-space(xs:annotation/xs:documentation), $lf)"/>
        <xsl:value-of select="concat('type ', @name, ' struct ', $lb, $cr)"/>
        <xsl:for-each select="xs:attribute">
            <xsl:variable name="r" select="@ref"/>
            <xsl:variable name="cap" select="concat(translate(substring(@ref, 1, 1), $lowers, $uppers), substring(@ref, 2))"/>
            <xsl:variable name="strtype">
                <xsl:choose>
                    <xsl:when test="/xs:schema/xs:attribute[@name = $r]/xs:simpleType/xs:restriction/@base = 'xsd:NMTOKENS'">
                        <xsl:text>[]string</xsl:text>
                    </xsl:when>
                    <xsl:otherwise>
                        <xsl:text>string</xsl:text>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:variable>
            <xsl:value-of select="concat($tabchar, $cap, $tabchar, $strtype, $tabchar, $bq, 'xml:', $qt, 'ism:', @ref, $cm, 'attr',$cm,$omitempty,$qt, ' ', $json, $qt, @ref, $cm, $omitempty, $qt, $bq, $cr)"/>
        </xsl:for-each>
        <xsl:value-of select="concat($rb, $cr)"/>
    </xsl:template>
    <xsl:variable name="tabchar">
        <xsl:text>&#09;</xsl:text>
    </xsl:variable>
    <xsl:variable name="lowers" select='"abcdefghijklmnopqrstuvwxyz"'/>
    <xsl:variable name="uppers" select='"ABCDEFGHIJKLMNOPQRSTUVWXYZ"'/>
    

</xsl:stylesheet>
