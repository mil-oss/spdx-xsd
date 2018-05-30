<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:owl="http://www.w3.org/2002/07/owl#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" exclude-result-prefixes="xs" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#"
    version="2.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:variable name="rdfSrc" select="document('../../resources/SPDX.rdf')"/>
    <xsl:variable name="xsdOut" select="'../xsd/spdx_nodes.xsd'"/>
    <xsl:variable name="xmlOut" select="'../instance/elements.xml'"/>



    <xsl:variable name="spdx_names">
        <xsl:variable name="nodes">
            <xsl:for-each select="$rdfSrc//@rdf:about">
                <xsl:element name="{substring-after(parent::*/name(),':')}">
                    <xsl:attribute name="xmlname">
                        <xsl:apply-templates select="." mode="xmlname"/>
                    </xsl:attribute>
                    <xsl:attribute name="id">
                        <xsl:value-of select="."/>
                    </xsl:attribute>
                </xsl:element>
            </xsl:for-each>
            <xsl:for-each select="$rdfSrc//@rdf:resource">
                <xsl:element name="{substring-after(parent::*/name(),':')}">
                    <xsl:attribute name="xmlname">
                        <xsl:apply-templates select="." mode="xmlname"/>
                    </xsl:attribute>
                    <xsl:attribute name="id">
                        <xsl:value-of select="."/>
                    </xsl:attribute>
                </xsl:element>
            </xsl:for-each>
            <xsl:for-each select="$rdfSrc//@rdf:datatype">
                <xsl:element name="{substring-after(parent::*/name(),':')}">
                    <xsl:attribute name="xmlname">
                        <xsl:apply-templates select="." mode="xmlname"/>
                    </xsl:attribute>
                    <xsl:attribute name="id">
                        <xsl:value-of select="."/>
                    </xsl:attribute>
                </xsl:element>
            </xsl:for-each>
        </xsl:variable>
            <xsl:for-each select="$nodes/*">
                <xsl:sort select="name()"/>
                <xsl:sort select="@xmlname"/>
                <xsl:variable name="i" select="@id"/>
                <xsl:choose>
                    <xsl:when test="preceding-sibling::*/@id = $i"/>
                    <xsl:otherwise>
                        <xsl:copy-of select="."/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:for-each>
        
    </xsl:variable>

    <xsl:template name="main">
        <xsl:result-document href="{$xmlOut}">
            <xsl:copy-of select="$spdx_names" copy-namespaces="no"/>
        </xsl:result-document>
        <xsl:result-document href="{$xsdOut}">
            <xs:schema xmlns="spdx:xsd::1.0" xmlns:ct="http://release.niem.gov/niem/conformanceTargets/3.0/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                xmlns:niem-xs="http://release.niem.gov/niem/proxy/xsd/4.0/" xmlns:structures="http://release.niem.gov/niem/structures/4.0/" xmlns:appinfo="http://release.niem.gov/niem/appinfo/4.0/"
                xmlns:ism="urn:us:gov:ic:ism" xmlns:xs="http://www.w3.org/2001/XMLSchema" attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="spdx:xsd::1.0" version="1"
                xsi:schemaLocation="http://release.niem.gov/niem/appinfo/4.0/ niem/utility/appinfo/4.0/appinfo.xsd http://release.niem.gov/niem/conformanceTargets/3.0/ niem/utility/conformanceTargets/3.0/conformanceTargets.xsd"
                ct:conformanceTargets="http://reference.niem.gov/niem/specification/naming-and-design-rules/4.0/#ReferenceSchemaDocument">
                <xs:import schemaLocation="ext/niem/utility/structures/4.0/structures.xsd" namespace="http://release.niem.gov/niem/structures/4.0/"/>
                <xs:import schemaLocation="ext/niem/utility/appinfo/4.0/appinfo.xsd" namespace="http://release.niem.gov/niem/appinfo/4.0/"/>
                <xs:import schemaLocation="ext/niem/proxy/xsd/4.0/xs.xsd" namespace="http://release.niem.gov/niem/proxy/xsd/4.0/"/>
                <xs:import namespace="http://www.w3.org/2002/07/owl#" schemaLocation="ext/owl.xsd"/>
                <xs:annotation>
                    <xs:documentation>This is a NIEM Conformant Reference Schema for Software Product Documentation Exchange (SPDX) information.</xs:documentation>
                </xs:annotation>
                <xs:simpleType name="AnnotationTypeSimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for a type of annotation.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Annotation" mapvar="annotationSimpleType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:string">
                        <xs:enumeration value="Other">
                            <xs:annotation>
                                <xs:documentation>Type of annotation which does not fit in any of the pre-defined annotation types.</xs:documentation>
                            </xs:annotation>
                        </xs:enumeration>
                        <xs:enumeration value="Review">
                            <xs:annotation>
                                <xs:documentation>A Review represents an audit and signoff by an individual, organization or tool on the information for an SpdxElement.</xs:documentation>
                            </xs:annotation>
                        </xs:enumeration>
                    </xs:restriction>
                </xs:simpleType>
                <xs:simpleType name="DateTimeSimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for SPDX Date Properties.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Object Property" mapvar="objectPropertySimpleType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:dateTime"/>
                </xs:simpleType>
                <xs:simpleType name="HexBinarySimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for SPDX Hex Properties.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Object Property" mapvar="objectPropertySimpleType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:hexBinary"/>
                </xs:simpleType>
                <xs:simpleType name="LinkUrlSimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for URLs.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="LinkUrl" mapvar="linkUrl"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:anyURI"/>
                </xs:simpleType>
                <xs:simpleType name="PropertyIndicatorSimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for the boolean indication of a property existence. True if known. False if not or not known.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Property Indicator" mapvar="propertyIndicator"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:boolean"/>
                </xs:simpleType>
                <xs:simpleType name="PropertySimpleType">
                    <xs:annotation>
                        <xs:documentation>A data type for SPDX Object Properties.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Object Property" mapvar="objectPropertySimpleType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:restriction base="xs:string"/>
                </xs:simpleType>
                <xs:complexType name="AnnotationType">
                    <xs:annotation>
                        <xs:documentation>Type of the annotation.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Annotation Type" mapvar="annotationType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="AnnotationTypeSimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="ChecksumType">
                    <xs:annotation>
                        <xs:documentation>The checksum property provides a mechanism that can be used to verify that the contents of a File or Package have not changed.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Checksum Type" mapvar="checksumType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="PropertySimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="CopyrightType">
                    <xs:annotation>
                        <xs:documentation>The text of copyright declarations recited in the Package or File.</xs:documentation>
                        <xs:appinfo>
                            <SimpleType name="Annotation" mapvar="annotationSimpleType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:complexContent>
                        <xs:extension base="structures:ObjectType">
                            <xs:choice>
                                <xs:element ref="NoAssertionIndicator">
                                    <xs:annotation>
                                        <xs:documentation>Indicates that the preparer of the SPDX document is not making any assertion regarding the value of this field.</xs:documentation>
                                    </xs:annotation>
                                </xs:element>
                                <xs:element ref="NoneIndicator">
                                    <xs:annotation>
                                        <xs:documentation>When this value is used as the object of a property it indicates that the preparer of the SpdxDocument believes that there is no value for the
                                            property. This value should only be used if there is sufficient evidence to support this assertion.</xs:documentation>
                                    </xs:annotation>
                                </xs:element>
                                <xs:element ref="CopyrightText">
                                    <xs:annotation>
                                        <xs:documentation>The text of copyright declarations recited in the Package or File.</xs:documentation>
                                    </xs:annotation>
                                </xs:element>
                            </xs:choice>
                        </xs:extension>
                    </xs:complexContent>
                </xs:complexType>
                <xs:complexType name="DateTimeType">
                    <xs:annotation>
                        <xs:documentation>The date and time. This value must in UTC and have 'Z' as its timezone indicator.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Date Type" mapvar="date"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="DateTimeSimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="HexBinaryType">
                    <xs:annotation>
                        <xs:documentation>Verification code as a hex encoded value.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Hex Binary" mapvar="hexBinary"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="HexBinarySimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="LinkUrlType">
                    <xs:annotation>
                        <xs:documentation>A data type for URLs.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Link Url Type" mapvar="linkUrl"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="LinkUrlSimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="PackageChecksumType">
                    <xs:annotation>
                        <xs:documentation>The checksum property provides a mechanism that can be used to verify that the contents of a File or Package have not changed.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Checksum Type" mapvar="checksumType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="PropertySimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="PropertyIndicatorType">
                    <xs:annotation>
                        <xs:documentation>A data type for the boolean indication of a property existence. True if known. False if not or not known.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Property Indicator" mapvar="propertyIndicator"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="PropertyIndicatorSimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xs:complexType name="PropertyType">
                    <xs:annotation>
                        <xs:documentation>A data type for SPDX Object Properties.</xs:documentation>
                        <xs:appinfo>
                            <ComplexType name="Property" mapvar="propertyType"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="PropertySimpleType">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
                <xsl:variable name="datatypes">
                    <xs:element name="NoAssertionIndicator" type="PropertyIndicatorType">
                        <xs:annotation>
                            <xs:documentation>Indicates that the preparer of the SPDX document is not making any assertion regarding the value of this field.</xs:documentation>
                        </xs:annotation>
                    </xs:element>
                    <xs:element name="NoneIndicator" type="PropertyIndicatorType">
                        <xs:annotation>
                            <xs:documentation>When this value is used as the object of a property it indicates that the preparer of the SpdxDocument believes that there is no value for the property.
                                This value should only be used if there is sufficient evidence to support this assertion.</xs:documentation>
                        </xs:annotation>
                    </xs:element>
                    <xs:element name="DateTime" type="DateTimeType">
                        <xs:annotation>
                            <xs:documentation>The date and time. This value must in UTC and have 'Z' as its timezone indicator.</xs:documentation>
                        </xs:annotation>
                    </xs:element>
                    <xs:element name="FileChecksum" type="ChecksumType">
                        <xs:annotation>
                            <xs:documentation>The checksum property provides a mechanism that can be used to verify that the contents of a File or Package have not changed.</xs:documentation>
                        </xs:annotation>
                    </xs:element>
                    <xs:element name="PackageChecksum" type="ChecksumType">
                        <xs:annotation>
                            <xs:documentation>The checksum property provides a mechanism that can be used to verify that the contents of a File or Package have not changed.</xs:documentation>
                        </xs:annotation>
                    </xs:element>
                    <xsl:apply-templates select="$rdfSrc/rdf:RDF/owl:DatatypeProperty[ns:term_status != 'deprecated']" mode="datatype"/>
                </xsl:variable>
                <xsl:for-each select="$rdfSrc//owl:onProperty[contains(@rdf:resource, '#')]">
                    <xsl:variable name="n">
                        <xsl:apply-templates select="@rdf:resource" mode="xmlname"/>
                    </xsl:variable>
                    <xsl:choose>
                        <xsl:when test="$datatypes/*[@name = $n]"/>
                        <xsl:otherwise>
                            <xs:element name="{$n}" type="PropertyType">
                                <xs:annotation>
                                    <xs:documentation>The checksum property provides a mechanism that can be used to verify that the contents of a File or Package have not changed.</xs:documentation>
                                </xs:annotation>
                            </xs:element>
                        </xsl:otherwise>
                    </xsl:choose>
                </xsl:for-each>
                <xsl:for-each select="$datatypes/*">
                    <xsl:sort select="@name"/>
                    <xsl:copy-of select="."/>
                </xsl:for-each>
                <xsl:variable name="objecttypes">
                    <xsl:apply-templates select="$rdfSrc/rdf:RDF/owl:ObjectProperty[ns:term_status != 'deprecated']" mode="object"/>
                </xsl:variable>
                <xsl:for-each select="$objecttypes/*">
                    <xsl:sort select="@name"/>
                    <xsl:copy-of select="."/>
                </xsl:for-each>
            </xs:schema>
        </xsl:result-document>
    </xsl:template>

    <xsl:template match="*" mode="datatype">
        <xsl:variable name="n">
            <xsl:apply-templates select="@rdf:about" mode="xmlname"/>
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
        <xs:element name="{$n}" type="{$t}">
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
        <xs:element name="{$n}" type="{$t}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="rdfs:comment"/>
                </xs:documentation>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template name="CapWord">
        <xsl:param name="text"/>
        <xsl:value-of select="translate(substring($text, 1, 1), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ')"/>
        <xsl:value-of select="substring($text, 2, string-length($text) - 1)"/>
    </xsl:template>

    <xsl:template match="@*" mode="xmlname">
        <xsl:variable name="n">
            <xsl:choose>
                <xsl:when test="string-length(.) = 0"/>
                <xsl:when test="contains(., '#')">
                    <xsl:call-template name="CapWord">
                        <xsl:with-param name="text">
                            <xsl:value-of select="substring-after(., '#')"/>
                        </xsl:with-param>
                    </xsl:call-template>
                </xsl:when>

                <xsl:when test="contains(., 'xsd;')">
                    <xsl:call-template name="CapWord">
                        <xsl:with-param name="text">
                            <xsl:value-of select="substring-after(., 'xsd;')"/>
                        </xsl:with-param>
                    </xsl:call-template>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:call-template name="CapWord">
                        <xsl:with-param name="text">
                            <xsl:value-of select="substring-after(name(), ':')"/>
                        </xsl:with-param>
                    </xsl:call-template>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xsl:choose>
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
                <xsl:value-of select="$n"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

</xsl:stylesheet>
