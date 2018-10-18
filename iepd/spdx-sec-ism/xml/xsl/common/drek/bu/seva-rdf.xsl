<?xml version="1.0" encoding="UTF-8"?>

<!DOCTYPE rdf:RDF [
    <!ENTITY foaf "http://xmlns.com/foaf/0.1/" >
    <!ENTITY owl "http://www.w3.org/2002/07/owl#" >
    <!ENTITY xsd "http://www.w3.org/2001/XMLSchema#" >
    <!ENTITY rdfs "http://www.w3.org/2000/01/rdf-schema#" >
    <!ENTITY rdf "http://www.w3.org/1999/02/22-rdf-syntax-ns#" >
    <!ENTITY ns "http://www.w3.org/2003/06/sw-vocab-status/ns#" >
]>

<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:spd="spdx:xsd::1.0/ref" xmlns:exsl="http://exslt.org/common"
    xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:security="spdx:xsd::1.0" version="1.0">

    <xsl:output method="xml" indent="yes"/>

    <!--<xsl:include href="identity.xsl"/>
    <xsl:include href="spdx_map.xsl"/>-->

    <!-- <xsl:variable name="spdx_xsd" select="document('../xsd/spdx-ref.xsd')"/>-->

    <xsl:variable name="Top" select="'AnyLicenseInfoType'"/>
    <xsl:variable name="Super" select="'SimpleLicensingInfoType'"/>
    <xsl:variable name="Root" select="'LicenseType'"/>
    <xsl:variable name="RootEl" select="'License'"/>
    <xsl:variable name="am">
        <xsl:text disable-output-escaping="yes">&amp;</xsl:text>
    </xsl:variable>
    <xsl:variable name="q" select="'&quot;'"/>
    <xsl:variable name="a" select='"&apos;"'/>
    <xsl:variable name="gt" select='"&gt;"'/>
    <xsl:variable name="lt" select='"&lt;"'/>
    <xsl:variable name="cr" select='"&#10;"'/>
    <xsl:variable name="sp" select='"&#32;"'/>
    <xsl:variable name="in">
        <xsl:value-of select="concat($sp, $sp, $sp, $sp)"/>
    </xsl:variable>
    <xsl:variable name="dt" select="concat($cr, $cr, '&lt;!DOCTYPE rdf:RDF [', $cr)"/>
    <xsl:variable name="foaf" select="concat($in, $in, $lt, '!ENTITY foaf ', $q, 'http://xmlns.com/foaf/0.1/', $q, $gt, $cr)"/>
    <xsl:variable name="owl" select="concat($in, $in, $lt, '!ENTITY owl ', $q, 'http://www.w3.org/2002/07/owl#', $q, $gt, $cr)"/>
    <xsl:variable name="xsd" select="concat($in, $in, $lt, '!ENTITY xsd ', $q, 'http://www.w3.org/2001/XMLSchema#', $q, $gt, $cr)"/>
    <xsl:variable name="rdfs" select="concat($in, $in, $lt, '!ENTITY rdfs ', $q, 'http://xmlns.com/foaf/0.1/', $q, $gt, $cr)"/>
    <xsl:variable name="rdf" select="concat($in, $in, $lt, '!ENTITY rdf ', $q, 'http://www.w3.org/1999/02/22-rdf-syntax-ns#', $q, $gt, $cr)"/>
    <xsl:variable name="ns" select="concat($in, $in, $lt, '!ENTITY ns ', $q, 'http://www.w3.org/2003/06/sw-vocab-status/ns#', $q, $gt, $cr)"/>

    <xsl:variable name="security" select="document('../xsd/security.xsd')/xs:schema"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <xsl:value-of select="$dt" disable-output-escaping="yes"/>
        <xsl:value-of select="$foaf" disable-output-escaping="yes"/>
        <xsl:value-of select="$owl" disable-output-escaping="yes"/>
        <xsl:value-of select="$xsd" disable-output-escaping="yes"/>
        <xsl:value-of select="$rdfs" disable-output-escaping="yes"/>
        <xsl:value-of select="$rdf" disable-output-escaping="yes"/>
        <xsl:value-of select="$ns" disable-output-escaping="yes"/>
        <xsl:value-of select="concat(']', $gt)" disable-output-escaping="yes"/>
        <xsl:value-of select="$cr"/>
        <xsl:value-of select="$cr"/>

        <rdf:RDF xmlns="http://www.w3.org/2000/01/rdf-schema#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#"
            xmlns:foaf="http://xmlns.com/foaf/0.1/" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
            <xsl:apply-templates select="$security/*[name() = 'xs:complexType'][@name = 'SoftwareEvidenceArchiveType']" mode="rdf"/>
            <xsl:apply-templates select="$security/*[name() = 'xs:complexType'][not(@name = 'SoftwareEvidenceArchiveType')]" mode="rdf"/>
            <xsl:apply-templates select="$security/*[name() = 'xs:simpleType']" mode="rdf"/>
            <xsl:apply-templates select="$security/*[name() = 'xs:element']" mode="rdf"/>
        </rdf:RDF>
    </xsl:template>

    <xsl:template match="xs:simpleType" mode="rdf">
        <owl:ObjectType rdf:ID="{@name}">
            <rdfs:comment xml:lang="en">
                <xsl:value-of select="xs:annotation/xs:documentation"/>
            </rdfs:comment>
            <ns:term_status xml:lang="en">proposed</ns:term_status>
            <xsl:choose>
                <xsl:when test="xs:restriction[@base = 'xs:string']">
                    <owl:onDataRange rdf:resource="xsd:string"/>
                    <xsl:apply-templates select="xs:restriction[xs:enumeration]"/>
                    <xsl:if test="xs:restriction/xs:pattern">
                        <xs:pattern xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:pattern/@value"/>
                        </xs:pattern>
                    </xsl:if>
                    <xsl:if test="xs:restriction/xs:minLength">
                        <xs:minLength xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:minLength/@value"/>
                        </xs:minLength>
                    </xsl:if>
                    <xsl:if test="xs:restriction/xs:maxLength">
                        <xs:maxLength xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:maxLength/@value"/>
                        </xs:maxLength>
                    </xsl:if>
                    <xsl:if test="xs:restriction/xs:length">
                        <xs:length xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:length/@value"/>
                        </xs:length>
                    </xsl:if>
                </xsl:when>
                <xsl:when test="xs:restriction[@base = 'xs:boolean']">
                    <owl:onDataRange rdf:resource="xsd:boolean"/>
                </xsl:when>
                <xsl:when test="xs:restriction[@base = 'xs:integer']">
                    <owl:onDataRange rdf:resource="xsd:integer"/>
                    <xsl:if test="xs:restriction/xs:minInclusive">
                        <xs:minInclusive xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:minInclusive/@value"/>
                        </xs:minInclusive>
                    </xsl:if>
                    <xsl:if test="xs:restriction/xs:maxInclusive">
                        <xs:maxInclusive xml:lang="en">
                            <xsl:value-of select="xs:restriction/xs:maxInclusive/@value"/>
                        </xs:maxInclusive>
                    </xsl:if>
                </xsl:when>
                <xsl:when test="xs:restriction[@base = 'xs:dateTime']">
                    <owl:onDataRange rdf:resource="xsd:dateTime"/>
                </xsl:when>
            </xsl:choose>
            <rdfs:domain rdf:resource="SoftwareEvidenceArchiveType'"/>
        </owl:ObjectType>
    </xsl:template>

    <xsl:template match="xs:complexType" mode="rdf">
        <owl:Class rdf:about="{@name}">
            <xsl:apply-templates select="xs:annotation" mode="rdfcomment"/>
            <ns:term_status xml:lang="en">proposed</ns:term_status>
            <xsl:choose>
                <xsl:when test="xs:sequence">
                    <owl:intersectionOf rdf:parseType="Collection">
                        <xsl:apply-templates select="xs:sequence/*" mode="rdf"/>
                    </owl:intersectionOf>
                </xsl:when>
                <xsl:otherwise>
                    <rdfs:subclass-of rdf:resource="{*/xs:extension/@base}"/>
                </xsl:otherwise>
            </xsl:choose>
        </owl:Class>
    </xsl:template>

    <xsl:template match="xs:annotation" mode="rdfcomment">
        <rdfs:comment xml:lang="en">
            <xsl:value-of select="normalize-space(xs:documentation)"/>
        </rdfs:comment>
    </xsl:template>

    <xsl:template match="xs:element[@name]" mode="rdf">
        <owl:ObjectProperty rdf:ID="{@name}">
            <rdfs:comment xml:lang="en"><xsl:value-of select="xs:annotation/xs:documentation"/></rdfs:comment>
            <ns:term_status xml:lang="en">proposed</ns:term_status>
            <rdfs:range rdf:resource="{@type}"/>
        </owl:ObjectProperty>
    </xsl:template>

    <xsl:template match="xs:element[@ref]" mode="rdf">
        <xsl:choose>
            <xsl:when test="@minOccurs and @maxOccurs">
                <owl:Restriction>
                    <owl:onProperty rdf:resource="{@ref}"/>
                    <owl:minCardinality rdf:datatype="&xsd;nonNegativeInteger">
                        <xsl:value-of select="@minOccurs"/>
                    </owl:minCardinality>
                    <owl:maxCardinality rdf:datatype="&xsd;nonNegativeInteger">
                        <xsl:value-of select="@maxOccurs"/>
                    </owl:maxCardinality>
                </owl:Restriction>
            </xsl:when>
            <xsl:when test="@minOccurs">
                <owl:Restriction>
                    <owl:onProperty rdf:resource="{@ref}"/>
                    <owl:minCardinality rdf:datatype="&xsd;nonNegativeInteger">
                        <xsl:value-of select="@minOccurs"/>
                    </owl:minCardinality>
                </owl:Restriction>
            </xsl:when>
            <xsl:when test="@maxOccurs">
                <owl:onProperty rdf:resource="{@ref}"/>
                <owl:maxCardinality rdf:datatype="&xsd;nonNegativeInteger">
                    <xsl:value-of select="@maxOccurs"/>
                </owl:maxCardinality>
            </xsl:when>
            <xsl:otherwise>
                <rdfs:subClassOf rdf:resource="{@ref}"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="*">
        <xsl:copy>
            <xsl:apply-templates select="@*"/>
            <xsl:apply-templates select="text()"/>
            <xsl:apply-templates select="*"/>
        </xsl:copy>
    </xsl:template>

    <xsl:template match="xs:complexContent">
        <xsl:apply-templates select="*"/>
    </xsl:template>

    <xsl:template match="xs:import"/>

    <xsl:template match="xs:restriction[xs:enumeration]">
        <owl:oneOf rdf:parseType="Collection">
            <xsl:apply-templates select="xs:enumeration"/>
        </owl:oneOf>
    </xsl:template>

    <xsl:template match="xs:enumeration">
        <owl:Thing rdf:about="{@value}">
            <rdfs:comment xml:lang="en">
                <xsl:value-of select="xs:annotation/xs:documentation"/>
            </rdfs:comment>
        </owl:Thing>
    </xsl:template>

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
