<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:spd="spdx:xsd::1.0/ref" xmlns:owl="http://www.w3.org/2002/07/owl#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" exclude-result-prefixes="xs owl rdf ns rdfs" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
    xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#" version="2.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="identity.xsl"/>
    <xsl:include href="spdx_map.xsl"/>

    <xsl:variable name="spdxMap">
        <xsl:call-template name="mapSpdx">
            <xsl:with-param name="rdfData" select="document('../../resources/SPDX.rdf')"/>
        </xsl:call-template>
    </xsl:variable>

    <xsl:variable name="xsdOut" select="'../xsd/spdx-ref.xsd'"/>
    <xsl:variable name="xmlOut" select="'../instance/spdx-map.xml'"/>

    <xsl:variable name="Enumerations">
        <xsl:variable name="all">
            <xsl:apply-templates select="$spdxMap/SPDX//Class/Union[Restriction/@hasvalue]" mode="enum"/>
        </xsl:variable>
        <xsl:for-each select="$all/*">
            <xsl:sort select="@name"/>
            <xsl:variable name="n" select="@name"/>
            <xsl:choose>
                <xsl:when test="preceding-sibling::*[@name = $n]"/>
                <xsl:when test="name() = 'xs:simpleType'">
                    <xsl:variable name="nn" select="@name"/>
                    <xsl:copy>
                        <xsl:copy-of select="@name"/>
                        <xsl:copy-of select="xs:annotation"/>
                        <xs:restriction base="xs:string">
                            <xsl:variable name="enums">
                                <xsl:for-each select="$all/*[@name = $nn]">
                                    <xsl:for-each select="xs:restriction/xs:enumeration">
                                        <xsl:copy-of select="."/>
                                    </xsl:for-each>
                                </xsl:for-each>
                            </xsl:variable>
                            <xsl:call-template name="deDupList">
                                <xsl:with-param name="list" select="$enums"/>
                            </xsl:call-template>
                        </xs:restriction>
                    </xsl:copy>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:copy-of select="."/>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:for-each>
    </xsl:variable>

    <xsl:variable name="Objects">
        <xsl:for-each select="$spdxMap/SPDX/Object">
            <xsl:variable name="n" select="@name"/>
            <xsl:choose>
                <xsl:when test="Class/Union/Restriction"/>
                <xsl:when test="$spdxMap/SPDX/Class[@name = $n]"/>
                <xsl:when test="$spdxMap/SPDX/Datatype[@name = $n]"/>
                <xsl:when test="$n = 'Agent'"/>
                <xsl:when test="$n = 'UsedBy'"/>
                <xsl:when test="$n = 'LicenseInfoInSnippets'"/>
                <xsl:when test="$n = 'Member'"/>
                <xsl:when test="$n = 'FileType'">
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="'FileType'"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:when test="$n = 'DataLicense'">
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="'LicenseId'"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:when test="$n = 'HasExtractedLicensingInfo'">
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="'SpdxDocument'"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:when test="@domain = 'SinglePointer'">
                    <xsl:apply-templates select="." mode="ctype">
                        <xsl:with-param name="type" select="'Pointer'"/>
                    </xsl:apply-templates>
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="'Pointer'"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:when test="@range = 'SinglePointer'">
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="'Pointer'"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:when test="@domain">
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="@domain"/>
                    </xsl:apply-templates>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:apply-templates select="." mode="element">
                        <xsl:with-param name="type" select="@range"/>
                    </xsl:apply-templates>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:for-each>
    </xsl:variable>

    <xsl:variable name="Datatypes">
        <xsl:apply-templates select="$spdxMap/SPDX//Datatype" mode="dt"/>
    </xsl:variable>

    <xsl:variable name="Classes">
        <xsl:apply-templates select="$spdxMap/SPDX//Class"/>
    </xsl:variable>

    <xsl:variable name="FsfLibre">
        <xs:element name="IsFsfLibre" type="IsFsfLibreType" nillable="true">
            <xs:annotation>
                <xs:documentation>A data item to indicate if the license is FSF Libre.</xs:documentation>
                <xs:appinfo>
                    <spd:Datatype name="IsFsfLibre" comment="Indicates if the license is is FSF Libre." rdf="http://spdx.org/rdf/terms#isFsfLibre" domain="License" range="Boolean"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
        <xs:complexType name="IsFsfLibreType">
            <xs:annotation>
                <xs:documentation>A data type to indicate if the license is is FSF Libre.</xs:documentation>
                <xs:appinfo>
                    <spd:Datatype name="IsFsfLibre" comment="Indicates if the license is is FSF Libre." rdf="http://spdx.org/rdf/terms#isFsfLibre" domain="License" range="Boolean"/>
                </xs:appinfo>
            </xs:annotation>
            <xs:simpleContent>
                <xs:extension base="PropertyIndicatorSimpleType">
                    <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                </xs:extension>
            </xs:simpleContent>
        </xs:complexType>
    </xsl:variable>

    <xsl:variable name="IsDeprecatedLicenseId">
        <xs:element name="IsDeprecatedLicenseId" type="IsDeprecatedLicenseIdType" nillable="true">
            <xs:annotation>
                <xs:documentation>A data item that indicates if the license id is Deprecated.</xs:documentation>
                <xs:appinfo>
                    <spd:Datatype name="IsDeprecatedLicenseId" comment="Indicates if the license id is Deprecated." rdf="http://spdx.org/rdf/terms#isDeprecatedLicenseId" domain="License"
                        range="Boolean"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
        <xs:complexType name="IsDeprecatedLicenseIdType">
            <xs:annotation>
                <xs:documentation>A data type that indicates if the license id is Deprecated.</xs:documentation>
                <xs:appinfo>
                    <spd:Datatype name="IsDeprecatedLicenseIdType" comment="Indicates if the license is is Deprecated." rdf="http://spdx.org/rdf/terms#isDeprecatedLicenseId" domain="License"
                        range="Boolean"/>
                </xs:appinfo>
            </xs:annotation>
            <xs:simpleContent>
                <xs:extension base="PropertyIndicatorSimpleType">
                    <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                </xs:extension>
            </xs:simpleContent>
        </xs:complexType>
    </xsl:variable>

    <xsl:template name="main">
        <xsl:result-document href="{$xsdOut}">
            <xs:schema xmlns="spdx:xsd::1.0" xmlns:spd="spdx:xsd::1.0/ref" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:owl="http://www.w3.org/2002/07/owl#"
                xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:ns="http://www.w3.org/2003/06/sw-vocab-status/ns#" xmlns:xs="http://www.w3.org/2001/XMLSchema"
                xmlns:ct="http://release.niem.gov/niem/conformanceTargets/3.0/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:niem-xs="http://release.niem.gov/niem/proxy/xsd/4.0/"
                xmlns:structures="http://release.niem.gov/niem/structures/4.0/" xmlns:appinfo="http://release.niem.gov/niem/appinfo/4.0/" xmlns:ism="urn:us:gov:ic:ism"
                attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="spdx:xsd::1.0" version="1"
                xsi:schemaLocation="http://release.niem.gov/niem/appinfo/4.0/ niem/utility/appinfo/4.0/appinfo.xsd http://release.niem.gov/niem/conformanceTargets/3.0/ ext/niem/utility/conformanceTargets/3.0/conformanceTargets.xsd"
                ct:conformanceTargets="http://reference.niem.gov/niem/specification/naming-and-design-rules/4.0/#ReferenceSchemaDocument">
                <xs:import schemaLocation="ext/niem/utility/structures/4.0/structures.xsd" namespace="http://release.niem.gov/niem/structures/4.0/"/>
                <xs:import schemaLocation="ext/niem/utility/appinfo/4.0/appinfo.xsd" namespace="http://release.niem.gov/niem/appinfo/4.0/"/>
                <xs:import schemaLocation="ext/niem/proxy/xsd/4.0/xs.xsd" namespace="http://release.niem.gov/niem/proxy/xsd/4.0/"/>
                <xs:import namespace="http://www.w3.org/2002/07/owl#" schemaLocation="ext/owl.xsd"/>
                <xsl:apply-templates select="$spdxMap/SPDX/Ontology" mode="annot"/>
                <xsl:variable name="allnodes">
                    <xs:simpleType name="PropertyIndicatorSimpleType">
                        <xs:annotation>
                            <xs:documentation>A data type for the boolean indication of a property existence. True if known. False if not or not known.</xs:documentation>
                            <xs:appinfo>
                                <spd:SimpleType name="Property Indicator" mapvar="propertyIndicator"/>
                            </xs:appinfo>
                        </xs:annotation>
                        <xs:restriction base="xs:boolean"/>
                    </xs:simpleType>
                    <xsl:copy-of select="$FsfLibre" copy-namespaces="no"/>
                    <xsl:copy-of select="$IsDeprecatedLicenseId" copy-namespaces="no"/>
                    <xsl:copy-of select="$Enumerations" copy-namespaces="no"/>
                    <xsl:copy-of select="$Objects" copy-namespaces="no"/>
                    <xsl:copy-of select="$Datatypes" copy-namespaces="no"/>
                    <xsl:copy-of select="$Classes" copy-namespaces="no"/>
                </xsl:variable>
                <xsl:for-each select="$allnodes/xs:simpleType">
                    <xsl:sort select="@name"/>
                    <xsl:copy-of select="." copy-namespaces="no"/>
                </xsl:for-each>
                <xsl:for-each select="$allnodes/xs:complexType">
                    <xsl:sort select="@name"/>
                    <xsl:copy-of select="." copy-namespaces="no"/>
                </xsl:for-each>
                <xsl:for-each select="$allnodes/xs:element">
                    <xsl:sort select="@name"/>
                    <xsl:copy-of select="." copy-namespaces="no"/>
                </xsl:for-each>
            </xs:schema>
        </xsl:result-document>
        <xsl:result-document href="{$xmlOut}">
            <xsl:copy-of select="$spdxMap" copy-namespaces="no"/>
        </xsl:result-document>
    </xsl:template>

    <xsl:template match="Ontology" mode="annot">
        <xs:annotation>
            <xs:documentation>
                <xsl:value-of select="@comment"/>
            </xs:documentation>
            <xs:appinfo>
                <xsl:element name="{concat('spd:',name())}">
                    <xsl:apply-templates select="@*" mode="identity"/>
                </xsl:element>
            </xs:appinfo>
        </xs:annotation>
    </xsl:template>

    <xsl:template match="SPDX/Class">
        <xsl:choose>
            <xsl:when test="@name = 'License'">
                <xs:complexType name="LicenseType">
                    <xs:annotation>
                        <xs:documentation>A data type for License type</xs:documentation>
                        <xs:appinfo>
                            <spd:Class name="License"
                                comment="A License represents a copyright license. The SPDX license list website is annotated with these properties (using RDFa) to allow license data published there to be easily processed. The license list is populated in accordance with the License List fields guidelines. These guidelines are not normative and may change over time. SPDX tooling should not rely on values in the license list conforming to the current guidelines."
                                rdf="http://spdx.org/rdf/terms#License"/>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:complexContent>
                        <xs:extension base="structures:ObjectType">
                            <xs:sequence>
                                <xs:element ref="IsDeprecatedLicenseId" minOccurs="0"/>
                                <xs:element ref="StandardLicenseHeader" minOccurs="0"/>
                                <xs:element ref="StandardLicenseTemplate" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="LicenseText" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="IsOsiApproved" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="IsFsfLibre" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="LicenseId" minOccurs="0"/>
                                <xs:element ref="Name" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="SeeAlso" minOccurs="0" maxOccurs="10"/>
                                <xs:element ref="Comment" minOccurs="0" maxOccurs="1"/>
                                <xs:element ref="LicenseAugmentationPoint" minOccurs="0" maxOccurs="unbounded"/>
                            </xs:sequence>
                        </xs:extension>
                    </xs:complexContent>
                </xs:complexType>
                <xs:element name="License" type="LicenseType" nillable="true">
                    <xs:annotation>
                        <xs:documentation>A data type for License type</xs:documentation>
                        <xs:appinfo>
                            <spd:Class name="License"
                                comment="A License represents a copyright license. The SPDX license list website is annotated with these properties (using RDFa) to allow license data published there to be easily processed. The license list is populated in accordance with the License List fields guidelines. These guidelines are not normative and may change over time. SPDX tooling should not rely on values in the license list conforming to the current guidelines."
                                rdf="http://spdx.org/rdf/terms#License"/>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
                <xs:element name="LicenseAugmentationPoint" abstract="true">
                    <xs:annotation>
                        <xs:documentation>An augmentation point for LicenseType</xs:documentation>
                        <xs:appinfo>
                            <spd:Element name="License Augmentation Point"/>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
            </xsl:when>
            <xsl:when test="@name = 'Pointer'">
                <xs:complexType name="{concat(@name,'Type')}">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('A data type for ', @comment)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <xsl:element name="{concat('spd:',name())}">
                                <xsl:apply-templates select="@*" mode="identity"/>
                            </xsl:element>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:simpleContent>
                        <xs:extension base="xs:string">
                            <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                        </xs:extension>
                    </xs:simpleContent>
                </xs:complexType>
            </xsl:when>
            <xsl:when test="ends-with(@name, 'Pointer')">
                <xs:element name="{@name}" type="PointerType" nillable="true">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('A data item for ', @name)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <spd:Element>
                                <xsl:apply-templates select="@*" mode="identity"/>
                            </spd:Element>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
            </xsl:when>
            <xsl:when test="@name = 'Project'">
                <xsl:apply-templates select="." mode="ctype">
                    <xsl:with-param name="type" select="'Project'"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="." mode="element">
                    <xsl:with-param name="type" select="'Project'"/>
                </xsl:apply-templates>
            </xsl:when>
            <xsl:when test="@name = 'Container'">
                <xsl:apply-templates select="." mode="ctype">
                    <xsl:with-param name="type" select="'Container'"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="." mode="element">
                    <xsl:with-param name="type" select="'Container'"/>
                </xsl:apply-templates>
            </xsl:when>
            <xsl:when test="@name = 'AnyLicenseInfo'">
                <xsl:apply-templates select="." mode="ctype">
                    <xsl:with-param name="type" select="'AnyLicenseInfo'"/>
                </xsl:apply-templates>
                <xsl:apply-templates select="." mode="element">
                    <xsl:with-param name="type" select="'AnyLicenseInfo'"/>
                </xsl:apply-templates>
            </xsl:when>
            <xsl:when test="not(SubClass/Restriction)">
                <xsl:variable name="type">
                    <xsl:choose>
                        <xsl:when test="SubClass[@rdf]/@name">
                            <xsl:value-of select="concat(SubClass[@rdf]/@name, 'Type')"/>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:value-of select="concat(@name, 'Type')"/>
                        </xsl:otherwise>
                    </xsl:choose>
                </xsl:variable>
                <xs:element name="{@name}" type="{$type}" nillable="true">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('A data item for ', @name)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <spd:Element>
                                <xsl:apply-templates select="@*" mode="identity"/>
                            </spd:Element>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
            </xsl:when>
            <xsl:otherwise>
                <xs:complexType name="{concat(@name,'Type')}">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('A data type for ', @name, ' type')"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <xsl:element name="{concat('spd:',name())}">
                                <xsl:apply-templates select="@*" mode="identity"/>
                            </xsl:element>
                        </xs:appinfo>
                    </xs:annotation>
                    <xs:complexContent>
                        <xs:extension base="structures:ObjectType">
                            <xs:sequence>
                                <xsl:apply-templates select="SubClass" mode="sclass"/>
                                <xs:element ref="{concat(@name,'AugmentationPoint')}" minOccurs="0" maxOccurs="unbounded"/>
                            </xs:sequence>
                        </xs:extension>
                    </xs:complexContent>
                </xs:complexType>
                <xs:element name="{concat(@name,'AugmentationPoint')}" abstract="true">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('An augmentation point for ', @name)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <spd:Element name="{concat(@name,' Augmentation Point')}" xmlns=""/>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
                <xs:element name="{@name}" type="{concat(@name,'Type')}" nillable="true">
                    <xs:annotation>
                        <xs:documentation>
                            <xsl:value-of select="concat('A data item for ', @name)"/>
                        </xs:documentation>
                        <xs:appinfo>
                            <spd:Element>
                                <xsl:apply-templates select="@*" mode="identity"/>
                            </spd:Element>
                        </xs:appinfo>
                    </xs:annotation>
                </xs:element>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="SubClass" mode="sclass">
        <xsl:choose>
            <xsl:when test="@name = 'Thing'"/>
            <xsl:when test="@name = 'AnyLicenseInfo'"/>
            <xsl:when test="Class/Union/Restriction">
                <xs:element ref="{concat(Class/Union/Restriction[1]/@onproperty,'Code')}"/>
            </xsl:when>
            <xsl:when test="@name">
                <xs:element ref="{@name}"/>
            </xsl:when>
            <xsl:otherwise>
                <xsl:apply-templates select="*" mode="sclass"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="*/Class/Union" mode="enum">
        <xsl:variable name="prop">
            <xsl:value-of select="Restriction[1]/@onproperty"/>
        </xsl:variable>
        <xs:simpleType name="{concat($prop,'CodeSimpleType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="concat('A data type for ', $prop, ' properties')"/>
                </xs:documentation>
                <xs:appinfo>
                    <spd:Codelist name="{concat($prop,'CodeSimpleType')}" rdf="{@rdf}"/>
                </xs:appinfo>
            </xs:annotation>
            <xs:restriction base="xs:string">
                <xsl:apply-templates select="Restriction[@hasvalue]" mode="enum"/>
            </xs:restriction>
        </xs:simpleType>
        <xs:complexType name="{concat($prop,'CodeType')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="concat('A data type for ', $prop, ' properties')"/>
                </xs:documentation>
                <xs:appinfo>
                    <spd:Codelist name="{$prop,'CodeSimpleType'}" rdf="{@rdf}"/>
                </xs:appinfo>
            </xs:annotation>
            <xs:simpleContent>
                <xs:extension base="{concat($prop,'CodeSimpleType')}">
                    <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                </xs:extension>
            </xs:simpleContent>
        </xs:complexType>
        <xs:element name="{concat($prop,'Code')}" type="{concat($prop,'CodeType')}" nillable="true">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="concat('A data item for ', $prop, ' properties')"/>
                </xs:documentation>
                <xs:appinfo>
                    <spd:Codelist name="{concat($prop,'Code')}" rdf="{@rdf}"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template match="Restriction" mode="enum">
        <xsl:variable name="v" select="@hasvalue"/>
        <xs:enumeration value="{$v}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:choose>
                        <xsl:when test="//*[@name = $v]/@comment">
                            <xsl:value-of select="//*[@name = $v]/@comment"/>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:call-template name="breakIntoWords">
                                <xsl:with-param name="string" select="@name"/>
                            </xsl:call-template>
                        </xsl:otherwise>
                    </xsl:choose>
                </xs:documentation>
                <xs:appinfo>
                    <spd:NamedIndividual rdf="{@rdf}"/>
                </xs:appinfo>
            </xs:annotation>
        </xs:enumeration>
    </xsl:template>

    <xsl:template match="Restriction" mode="sclass">
        
        <xsl:choose>
            <xsl:when test="@onproperty">
                <xs:element ref="{@onproperty}">
                    <xsl:copy-of select="@minOccurs"/>
                    <xsl:copy-of select="@maxOccurs"/>
                </xs:element>
            </xsl:when>
            <xsl:when test="@hasvalue">
                <xs:element ref="{@hasvalue}">
                    <xsl:copy-of select="@minOccurs"/>
                    <xsl:copy-of select="@maxOccurs"/>
                </xs:element>
            </xsl:when>
        </xsl:choose>

    </xsl:template>

    <xsl:template match="SPDX/Datatype" mode="dt">
        <xsl:variable name="base">
            <xsl:choose>
                <xsl:when test="@name = 'SeeAlso'">
                    <xsl:text>xs:anyURI</xsl:text>
                </xsl:when>
                <xsl:when test="@name = 'Comment'">
                    <xsl:text>xs:string</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'String'">
                    <xsl:text>xs:string</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'Literal'">
                    <xsl:text>xs:string</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'PositiveInteger'">
                    <xsl:text>xs:positiveInteger</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'HexBinary'">
                    <xsl:text>xs:hexBinary</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'AnyURI'">
                    <xsl:text>xs:anyURI</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'Boolean'">
                    <xsl:text>PropertyIndicatorSimpleType</xsl:text>
                </xsl:when>
                <xsl:when test="@range = 'DateTime'">
                    <xsl:text>xs:dateTime</xsl:text>
                </xsl:when>
                <xsl:when test="@subpropertyof = 'Date'">
                    <xsl:text>xs:dateTime</xsl:text>
                </xsl:when>
                <xsl:when test="@subpropertyof = 'Name'">
                    <xsl:text>xs:string</xsl:text>
                </xsl:when>
                <xsl:when test="@subpropertyof = 'LicenseInfoFromFiles'">
                    <xsl:text>PropertyIndicatorSimpleType</xsl:text>
                </xsl:when>
                <xsl:otherwise>
                    <xsl:text>PropertyIndicatorSimpleType</xsl:text>
                </xsl:otherwise>
            </xsl:choose>
        </xsl:variable>
        <xs:complexType name="{concat(@name,'Type')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:choose>
                        <xsl:when test="@name = 'SnippetName'">
                            <xsl:text>A data type to name specific snippet in a human convenient manner</xsl:text>
                        </xsl:when>
                        <xsl:when test="string-length(@comment) &gt; 0">
                            <xsl:value-of select="concat('A data type for ', @comment)"/>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:value-of select="concat('A data type for ', @name)"/>
                        </xsl:otherwise>
                    </xsl:choose>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:element name="{concat('spd:',name())}">
                        <xsl:apply-templates select="@*" mode="identity"/>
                    </xsl:element>
                </xs:appinfo>
            </xs:annotation>
            <xs:simpleContent>
                <xs:extension base="{$base}">
                    <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                </xs:extension>
            </xs:simpleContent>
        </xs:complexType>
        <xs:element name="{@name}" type="{concat(@name,'Type')}" nillable="true">
            <xs:annotation>
                <xs:documentation>
                    <xsl:choose>
                        <xsl:when test="@name = 'SnippetName'">
                            <xsl:text>A data item to name a specific snippet in a human convenient manner</xsl:text>
                        </xsl:when>
                        <xsl:when test="string-length(@comment) &gt; 0">
                            <xsl:value-of select="concat('A data type for ', @comment)"/>
                        </xsl:when>
                        <xsl:otherwise>
                            <xsl:value-of select="concat('A data type for ', @name)"/>
                        </xsl:otherwise>
                    </xsl:choose>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:element name="{concat('spd:',name())}">
                        <xsl:apply-templates select="@*" mode="identity"/>
                    </xsl:element>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template match="*" mode="ctype">
        <xs:complexType name="{concat(@name,'Type')}">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="concat('A data type for ', @comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:element name="{concat('spd:',name())}">
                        <xsl:apply-templates select="@*" mode="identity"/>
                    </xsl:element>
                </xs:appinfo>
            </xs:annotation>
            <xs:simpleContent>
                <xs:extension base="xs:string">
                    <xs:attributeGroup ref="structures:SimpleObjectAttributeGroup"/>
                </xs:extension>
            </xs:simpleContent>
        </xs:complexType>
    </xsl:template>

    <xsl:template match="*" mode="element">
        <xsl:param name="type"/>
        <xs:element name="{@name}" type="{concat($type,'Type')}" nillable="true">
            <xs:annotation>
                <xs:documentation>
                    <xsl:value-of select="concat('A data item for ', @comment)"/>
                </xs:documentation>
                <xs:appinfo>
                    <xsl:element name="{concat('spd:',name())}">
                        <xsl:apply-templates select="@*" mode="identity"/>
                    </xsl:element>
                </xs:appinfo>
            </xs:annotation>
        </xs:element>
    </xsl:template>

    <xsl:template match="SPDX/NamedIndividual"/>

    <xsl:template match="SPDX/Description"/>

    <xsl:template match="SPDX/Ontology"/>

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

    <xsl:template name="deDupList">
        <xsl:param name="list"/>
        <xsl:for-each select="$list/*">
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
