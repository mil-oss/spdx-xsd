<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:owl="http://www.w3.org/2002/07/owl#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" exclude-result-prefixes="xs owl rdf ns rdfs" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
    xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#" version="2.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:variable name="changes" select="document('../instance/changes.xml')/SpdxChanges"/>

    <xsl:template name="mapSpdx">
        <xsl:param name="rdfData"/>
        <xsl:element name="SPDX">
            <xsl:apply-templates select="$rdfData/rdf:RDF/*"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="owl:Ontology">
        <Ontology name="{rdfs:label}" version="{owl:versionInfo}" comment="{normalize-space(rdfs:comment)}" rdf="{@rdf:about}"/>
    </xsl:template>

    <xsl:template match="owl:ObjectProperty[contains(@rdf:about, ';member')]"/>

    <xsl:template match="owl:ObjectProperty">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xsl:variable name="xn">
            <xsl:apply-templates select="@rdf:about" mode="getname"/>
        </xsl:variable>
        <xsl:variable name="c">
            <xsl:apply-templates select="@rdf:about" mode="getcomment">
                <xsl:with-param name="comment" select="rdfs:comment"/>
            </xsl:apply-templates>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="string-length($n) = 0"/>
            <xsl:otherwise>
                <Object name="{$n}" xmlname="{$xn}" comment="{$c}" rdf="{@rdf:about}">
                    <xsl:if test="rdfs:domain/@rdf:resource">
                        <xsl:attribute name="domain">
                            <xsl:apply-templates select="rdfs:domain/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:if test="rdfs:range/@rdf:resource">
                        <xsl:attribute name="range">
                            <xsl:apply-templates select="rdfs:range/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:if test="rdf:type/@rdf:resource">
                        <xsl:attribute name="type">
                            <xsl:apply-templates select="rdf:type/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:if test="rdfs:subPropertyOf/@rdf:resource">
                        <xsl:attribute name="subpropertyof">
                            <xsl:apply-templates select="rdfs:subPropertyOf/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:apply-templates select="*"/>
                </Object>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="owl:DatatypeProperty">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xsl:variable name="xn">
            <xsl:apply-templates select="@rdf:about" mode="getname"/>
        </xsl:variable>
        <xsl:variable name="c">
            <xsl:apply-templates select="@rdf:about" mode="getcomment">
                <xsl:with-param name="comment" select="rdfs:comment"/>
            </xsl:apply-templates>
        </xsl:variable>
        <Datatype name="{$n}" xmlname="{$xn}" comment="{$c}" rdf="{@rdf:about}">
            <xsl:if test="rdfs:domain/@rdf:resource">
                <xsl:attribute name="domain">
                    <xsl:apply-templates select="rdfs:domain/@rdf:resource" mode="getname"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:if test="rdfs:range/@rdf:resource">
                <xsl:attribute name="range">
                    <xsl:apply-templates select="rdfs:range/@rdf:resource" mode="getname"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:if test="rdfs:subPropertyOf/@rdf:resource">
                <xsl:attribute name="subpropertyof">
                    <xsl:apply-templates select="rdfs:subPropertyOf/@rdf:resource" mode="getname"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:apply-templates select="*"/>
        </Datatype>
    </xsl:template>

    <xsl:template match="owl:Class">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xsl:variable name="xn">
            <xsl:apply-templates select="@rdf:about" mode="getname"/>
        </xsl:variable>
        <xsl:variable name="c">
            <xsl:apply-templates select="@rdf:about" mode="getcomment">
                <xsl:with-param name="comment" select="rdfs:comment"/>
            </xsl:apply-templates>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="string-length($n) &gt; 0">
                <Class name="{$n}" xmlname="{$xn}" comment="{$c}" rdf="{@rdf:about}">
                    <xsl:if test="rdfs:isDefinedBy">
                        <xsl:attribute name="definedby">
                            <xsl:value-of select="substring-after(rdfs:isDefinedBy, '#')"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:if test="owl:disjointWith">
                        <xsl:attribute name="disjointwith">
                            <xsl:apply-templates select="owl:disjointWith/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:if test="rdfs:subClassOf[@rdf:resource]">
                        <xsl:attribute name="subclassof">
                            <xsl:apply-templates select="rdfs:subClassOf[1]/@rdf:resource" mode="getname"/>
                        </xsl:attribute>
                    </xsl:if>
                    <xsl:apply-templates select="*"/>
                </Class>
            </xsl:when>
            <xsl:otherwise>
                <Class>
                    <xsl:apply-templates select="*"/>
                </Class>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="owl:NamedIndividual">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xsl:variable name="xn">
            <xsl:apply-templates select="@rdf:about" mode="getname"/>
        </xsl:variable>
        <xsl:variable name="c">
            <xsl:apply-templates select="@rdf:about" mode="getcomment">
                <xsl:with-param name="comment" select="rdfs:comment"/>
            </xsl:apply-templates>
        </xsl:variable>
        <NamedIndividual name="{$n}" xmlname="{$xn}" comment="{$c}" rdf="{@rdf:about}"/>
    </xsl:template>

    <xsl:template match="rdf:Description">
        <Description>
            <xsl:if test="rdf:type/@rdf:resource">
                <xsl:attribute name="type">
                    <xsl:apply-templates select="rdf:type/@rdf:resource" mode="xmlname"/>
                </xsl:attribute>
                <xsl:attribute name="rdf">
                    <xsl:value-of select="rdf:type/@rdf:resource"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:if test="@rdf:about">
                <xsl:attribute name="name">
                    <xsl:apply-templates select="@rdf:about" mode="xmlname"/>
                </xsl:attribute>
                <xsl:attribute name="rdf">
                    <xsl:value-of select="@rdf:about"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:apply-templates select="*"/>
        </Description>
    </xsl:template>

    <xsl:template match="owl:distinctMembers">
        <Choice>
            <xsl:apply-templates select="*"/>
        </Choice>
    </xsl:template>

    <xsl:template match="owl:members">
        <Sequence>
            <xsl:apply-templates select="*"/>
        </Sequence>
    </xsl:template>

    <xsl:template match="owl:unionOf">
        <Union rdf="{owl:Restriction[1]/owl:onProperty/@rdf:resource}">
            <xsl:apply-templates select="*"/>
        </Union>
    </xsl:template>

    <xsl:template match="owl:Restriction">
        <Restriction rdf="{owl:onProperty/@rdf:resource}">
            <xsl:apply-templates select="*" mode="att"/>
            <xsl:attribute name="xmlname">
                <xsl:apply-templates select="owl:onProperty/@rdf:resource" mode="getname"/>
            </xsl:attribute>
        </Restriction>
    </xsl:template>

    <xsl:template match="owl:onProperty" mode="att">
        <xsl:attribute name="onproperty">
            <xsl:apply-templates select="@rdf:resource" mode="getname"/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="owl:onClass" mode="att">
        <xsl:attribute name="onclass">
            <xsl:apply-templates select="@rdf:resource" mode="getname"/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="owl:minQualifiedCardinality" mode="att">
        <xsl:attribute name="minOccurs">
            <xsl:value-of select="."/>
        </xsl:attribute>
        <xsl:if test="number(.) &gt; 1">
            <xsl:attribute name="maxOccurs">
                <xsl:text>unbounded</xsl:text>
            </xsl:attribute>
        </xsl:if>
    </xsl:template>

    <xsl:template match="owl:maxQualifiedCardinality" mode="att">
        <xsl:attribute name="maxOccurs">
            <xsl:value-of select="."/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="owl:qualifiedCardinality" mode="att">
        <xsl:attribute name="minOccurs">
            <xsl:value-of select="."/>
        </xsl:attribute>
        <xsl:attribute name="maxOccurs">
            <xsl:value-of select="."/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="owl:onDataRange" mode="att">
        <xsl:attribute name="ondatarange">
            <xsl:value-of select="substring-after(@rdf:resource, '#')"/>
            <xsl:value-of select="substring-after(@rdf:resource, ';')"/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="owl:hasValue" mode="att">
        <xsl:attribute name="hasvalue">
            <xsl:apply-templates select="@rdf:resource" mode="getname"/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="@rdf:resource" mode="att">
        <xsl:attribute name="hasvalue">
            <xsl:apply-templates select="." mode="getname"/>
        </xsl:attribute>
    </xsl:template>

    <xsl:template match="rdfs:range">
        <xsl:apply-templates select="*"/>
    </xsl:template>

    <xsl:template match="rdfs:subClassOf">
        <SubClass>
            <xsl:if test="@rdf:resource">
                <xsl:attribute name="rdf">
                    <xsl:value-of select="@rdf:resource"/>
                </xsl:attribute>
            </xsl:if>
            <xsl:apply-templates select="@rdf:resource" mode="att"/>
            <xsl:apply-templates select="*"/>
        </SubClass>
    </xsl:template>

    <!--<xsl:template match="owl:Class">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    
    <xsl:template match="owl:unionOf">
        <xsl:apply-templates select="*"/>
    </xsl:template>
    
    <xsl:template match="owl:Restriction">
        <xsl:apply-templates select="*"/>
    </xsl:template>-->
    <xsl:template match="rdf:type"/>
    <xsl:template match="rdfs:comment"/>
    <xsl:template match="rdfs:domain"/>
    <xsl:template match="rdfs:subPropertyOf"/>
    <xsl:template match="rdfs:isDefinedBy"/>
    <xsl:template match="ns:term_status"/>
    <xsl:template match="owl:Axiom"/>
    <xsl:template match="owl:disjointWith"/>
    <xsl:template match="owl:AnnotationProperty"/>
    <xsl:template match="*[ns:term_status = 'deprecated' or owl:deprecatedProperty]"/>

    <xsl:template match="*" mode="datatype">
        <xsl:variable name="n">
            <xsl:apply-templates select="@rdf:about" mode="xmlname"/>
        </xsl:variable>
        <xsl:variable name="lcn">
            <xsl:call-template name="LCaseWord">
                <xsl:with-param name="text" select="$n"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:variable name="t">
            <xsl:choose>
                <xsl:when test="contains(*/@rdf:resource[0], 'string')">
                    <xsl:text>PropertyType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'date')">
                    <xsl:text>DateType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'hexBinary')">
                    <xsl:text>HexBinaryType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'boolean')">
                    <xsl:text>PropertyIndicatorType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'anyURI')">
                    <xsl:text>LinkUrlType</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>PropertyType</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xs:element name="{$lcn}" type="{$t}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="rdfs:comment"/>
                </xs:documentation>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template match="*" mode="object">
        <xsl:variable name="n">
            <xsl:apply-templates select="@rdf:about" mode="xmlname"/>
        </xsl:variable>
        <xsl:variable name="lcn">
            <xsl:call-template name="LCaseWord">
                <xsl:with-param name="text" select="$n"/>
            </xsl:call-template>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="rdfs:range/owl:Class"/>
            <xsl:otherwise> </xsl:otherwise>
        </xsl:choose>
        <xsl:variable name="t">
            <xsl:choose>
                <xsl:when test="contains(*/@rdf:resource[0], 'string')">
                    <xsl:text>PropertyType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'date')">
                    <xsl:text>DateType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'hexBinary')">
                    <xsl:text>HexBinaryType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'boolean')">
                    <xsl:text>PropertyIndicatorType</xsl:text>
                </xsl:when>
                <xsl:when test="contains(*/@rdf:resource[0], 'anyURI')">
                    <xsl:text>LinkUrlType</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>PropertyType</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xs:element name="{$lcn}" type="{$t}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="rdfs:comment"/>
                </xs:documentation>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template name="UCaseWord">
        <xsl:param name="text"/>
        <xsl:variable name="w">
            <xsl:value-of select="translate(substring($text, 1, 1), 'abcdefghijklmnopqrstuvwxyz','ABCDEFGHIJKLMNOPQRSTUVWXYZ')"/>
            <xsl:value-of select="substring($text, 2, string-length($text) - 1)"/>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="ends-with($w, 'Id')">
                <xsl:value-of select="concat(substring($w, 0, string-length($text) - 1), 'ID')"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="$w"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="@*" mode="getname">
        <xsl:variable name="r" select="."/>
        <xsl:choose>
            <xsl:when test="$changes/*[@rdf = $r]">
                <xsl:value-of select="$changes/*[@rdf = $r]/@xmlname"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:apply-templates select="." mode="xmlname"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="@*" mode="getcomment">
        <xsl:param name="comment"/>
        <xsl:choose>
            <xsl:when test="$changes/*[@rdf = .]">
                <xsl:value-of select="$changes/*[@rdf = .]/@comment"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="normalize-space($comment)"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="@*" mode="xmlname">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="string-length(.) = 0"/>
                <xsl:when test="contains(., '#SimpleLicenseingInfo')">
                    <xsl:text>SimpleLicensingInfo</xsl:text>
                </xsl:when>
                <xsl:when test="contains(., '#')">
                    <xsl:call-template name="UCaseWord">
                        <xsl:with-param name="text">
                            <xsl:value-of select="substring-after(., '#')"/>
                        </xsl:with-param>
                    </xsl:call-template>
                </xsl:when>
                <xsl:when test="contains(., 'xsd;')">
                    <xsl:call-template name="UCaseWord">
                        <xsl:with-param name="text">
                            <xsl:value-of select="substring-after(., 'xsd;')"/>
                        </xsl:with-param>
                    </xsl:call-template>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:call-template name="extract-lastname">
                        <xsl:with-param name="path" select="."/>
                    </xsl:call-template>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="contains($n, '_')">
                <xsl:variable name="pre">
                    <xsl:call-template name="UCaseWord">
                        <xsl:with-param name="text" select="substring-before($n, '_')"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:variable name="suf">
                    <xsl:call-template name="UCaseWord">
                        <xsl:with-param name="text" select="substring-after($n, '_')"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:value-of select="concat($pre, $suf)"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="$n"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="extract-lastname">
        <xsl:param name="path" select="."/>
        <xsl:choose>
            <xsl:when test="not(contains($path, '/'))">
                <xsl:value-of select="$path"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:call-template name="extract-lastname">
                    <xsl:with-param name="path" select="substring-after($path, '/')"/>
                </xsl:call-template>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

</xsl:stylesheet>
