<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
    xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" exclude-result-prefixes="xs" version="2.0">

    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="identity.xsl"/>

    <xsl:variable name="rdfSrc" select="document('../../resources/SPDX.rdf')"/>
    <xsl:variable name="xsdOut" select="'../xsd/spdx.xsd'"/>

    <xsl:template name="main">
        <xsl:result-document href="{$xsdOut}">
            <xs:schema xmlns="spdx:xsd" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
                xmlns:xs="http://www.w3.org/2001/XMLSchema" elementFormDefault="qualified" targetNamespace="spdx:xsd">
                <xs:import namespace="http://www.w3.org/2002/07/owl#" schemaLocation="ext/owl.xsd"/>
                <xsl:apply-templates select="$rdfSrc/rdf:RDF/owl:Class"/>
            </xs:schema>
        </xsl:result-document>
    </xsl:template>

    <xsl:template match="owl:Class">
        <xsl:variable name="n">
            <xsl:apply-templates select="@rdf:about" mode="xmlname"/>
        </xsl:variable>
        <xsl:choose>
            <xsl:when test="ns:term_status/text() = 'deprecated'"/>
            <xsl:when test="rdfs:subClassOf[@rdf:resource]">
                <xsl:variable name="sname">
                    <xsl:apply-templates select="@rdf:resource" mode="xmlname"/>
                </xsl:variable>
                <xs:complexType name="{concat($n,'Type')}">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="normalize-space(rdfs:comment)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <xsl:copy-of select="rdf:about"/>
                            <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:complexContent>
                        <xs:extension base="{concat($sname,'Type')}">
                            <xsl:for-each select="rdfs:subClassOf[owl:Restriction]">
                                <xsl:variable name="rname">
                                    <xsl:apply-templates select="owl:onProperty/@rdf:resource" mode="xmlname"/>
                                </xsl:variable>
                                <xs:element ref="{$rname}">
                                    <xsl:choose>
                                        <xsl:when test="owl:minQualifiedCardinality">
                                            <xsl:attribute name="minOccurs">
                                                <xsl:value-of select="."/>
                                            </xsl:attribute>
                                        </xsl:when>
                                        <xsl:when test="owl:maxQualifiedCardinality">
                                            <xsl:attribute name="maxOccurs">
                                                <xsl:value-of select="."/>
                                            </xsl:attribute>
                                        </xsl:when>
                                        <xsl:when test="owl:qualifiedCardinality">
                                            <xsl:attribute name="minOccurs">
                                                <xsl:value-of select="."/>
                                            </xsl:attribute>
                                        </xsl:when>
                                    </xsl:choose>
                                </xs:element>
                            </xsl:for-each>
                        </xs:extension>
                    </xs:complexContent>
                </xs:complexType>
            </xsl:when>
            <xsl:otherwise>
                <xs:element name="{$n}" abstract="true">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="rdfs:comment"/>
                        </xs:documentation>
                    </xs:annotation>
                </xs:element>
            </xsl:otherwise>
        </xsl:choose>
        <xs:complexType name="{concat($n,'ClassType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>

    <xsl:template match="@*" mode="xmlname">
        <xsl:variable name="n" select="substring-after(., '#')"/>
        <xsl:choose>
            <xsl:when test="string-length($n) = 0"/>
            <xsl:when test="contains($n, '_')">
                <xsl:variable name="pre">
                    <xsl:call-template name="CapWord">
                        <xsl:with-param name="text" select="substring-before($n, '_')"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:variable name="suf">
                    <xsl:call-template name="CapWord">
                        <xsl:with-param name="text" select="substring-after($n, '_')"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:value-of select="concat($pre, $suf)"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:call-template name="CapWord">
                    <xsl:with-param name="text" select="$n"/>
                </xsl:call-template>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template name="CapWord">
        <xsl:param name="text"/>
        <xsl:value-of select="translate(substring($text, 1, 1), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ')"/>
        <xsl:value-of select="substring($text, 2, string-length($text) - 1)"/>
    </xsl:template>

    <xsl:template match="rdf:RDF">
        <xs:complexType name="SpdxType">
            <xsl:apply-templates select="owl:Ontology"/>
        </xs:complexType>
        <xsl:apply-templates select="*[not(name() = 'owl:Ontology')]"/>
    </xsl:template>
    <xsl:template match="owl:Ontology">
        <xs:annotation>
            <xs:documentation>
                <xsl:value-of select="normalize-space(rdfs:comment)"/>
            </xs:documentation>
            <xs:appinfo>
                <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                <xsl:apply-templates select="owl:AnnotationProperty/*"/>
            </xs:appinfo>
        </xs:annotation>
    </xsl:template>
    <xsl:template match="owl:AnnotationProperty"/>
    <xsl:template match="owl:ObjectProperty[*]">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xs:complexType name="{concat($n,'ObjectType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
    <xsl:template match="owl:ObjectProperty[not(*)]"/>
    <xsl:template match="owl:DatatypeProperty">
        <xsl:variable name="n" select="substring-after(@rdf:about, '#')"/>
        <xs:complexType name="{concat($n,'DataType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
    <xsl:template match="owl:Axiom">
        <xsl:variable name="n" select="substring-after(owl:annotatedSource/@rdf:resource, '#')"/>
        <xs:complexType name="{concat($n,'AxiomType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
    <xsl:template match="owl:NamedIndividual">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="contains(@rdf:about, '#')">
                    <xsl:value-of select="substring-after(@rdf:about, '#')"/>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>CC0</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xs:element name="{$n}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <owl:NamedIndividual>
                        <xsl:copy-of select="@rdf:about"/>
                    </owl:NamedIndividual>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
    </xsl:template>
    <xsl:template match="rdf:Description[rdf:type]">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="owl:distinctMembers/rdf:Description[@rdf:about = 'http://spdx.org/rdf/terms#noassertion']">
                    <xsl:text>NoDescriptionType</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:value-of select="substring-after(rdf:type/@rdf:resource, '#')"/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>

        <xs:complexType name="{concat($n,'DescriptionType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
    <xsl:template match="rdf:Description[@rdf:about]">
        <xsl:variable name="n">
            <xsl:value-of select="substring-after(@rdf:about, '#')"/>
        </xsl:variable>
        <xs:complexType name="{concat($n,'DescriptionType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
    <xsl:template match="rdf:Description[owl:members/rdf:Description/@rdf:about = 'http://spdx.org/rdf/terms#none']">]"> <xs:complexType name="NoneDescriptionType">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="normalize-space(rdfs:comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:copy-of select="rdf:about"/>
                    <xsl:apply-templates select="*[not(name() = 'rdfs:comment')]"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:complexType>
    </xsl:template>
</xsl:stylesheet>
