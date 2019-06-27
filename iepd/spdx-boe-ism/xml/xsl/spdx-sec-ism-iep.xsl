<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns="urn:spdx-xml:1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:exsl="http://exslt.org/common" version="1.0">

    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./common/iep.xsl"/>

    <!--<xsl:variable name="spdx_xsd" select="document('./../../../../xml/xsd/spdx-xml-ref.xsd')"/>-->

    <xsl:variable name="Super" select="'SpdxElementType'"/>
    <xsl:variable name="Root" select="'SoftwareEvidenceArchiveType'"/>
    <xsl:variable name="RootEl" select="'SoftwareEvidenceArchive'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:variable name="xsddata">
        <xs:annotation>
            <xs:documentation>XML Schema for Software Evidence Archive Information Exchange</xs:documentation>
            <xs:appinfo>
                <Root type="{$Root}" name="{$RootEl}"/>
            </xs:appinfo>
        </xs:annotation>
        <xsl:apply-templates select="//xs:schema/*[@name = $Root]"/>
        <xsl:apply-templates select="/xs:schema/*[@name = $RootEl]"/>
        <xsl:variable name="allnodes">
            <!--<xsl:apply-templates select="//xs:schema/*[@name = $Super]"/>-->
            <xsl:call-template name="deDupList">
                <xsl:with-param name="list">
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Root]"/>
                        <xsl:with-param name="iteration" select="15"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Super]"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'SpdxElement']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'CreationInfoType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'AnnotationType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'AnnotationTypeCodeType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'ExternalDocumentRefType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'ChecksumType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/xs:element[@name = 'RelationshipType']"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                </xsl:with-param>
            </xsl:call-template>
        </xsl:variable>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:simpleType">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:complexType[not(@name = $Root)]">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
        <xsl:for-each select="exsl:node-set($allnodes)/xs:element">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
    </xsl:variable>

    <xsl:template match="xs:complexType/xs:complexContent[xs:extension[xs:sequence]]" mode="xsdcopy">
        <xs:complexContent>
            <xs:extension base="structures:ObjectType">
                <xsl:apply-templates select="xs:extension/*" mode="xsdcopy"/>
                <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
            </xs:extension>
        </xs:complexContent>
    </xsl:template>

    <xsl:template match="xs:complexContent[xs:extension[not(xs:sequence)]]" mode="xsdcopy">
        <xs:complexContent>
            <xsl:apply-templates select="*" mode="xsdcopy"/>
            <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
        </xs:complexContent>
    </xsl:template>

    <xsl:template match="xs:complexContent[xs:extension[not(xs:sequence)]]">
        <xs:complexContent>
            <xsl:apply-templates select="*"/>
            <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
        </xs:complexContent>
    </xsl:template>

    <xsl:template match="xs:complexContent">
        <xs:complexContent>
            <xsl:apply-templates select="*"/>
            <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
        </xs:complexContent>
    </xsl:template>

    <xsl:template match="xs:complexContent" mode="xsdcopy">
        <xs:complexContent>
            <xsl:apply-templates select="*" mode="xsdcopy"/>
            <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
        </xs:complexContent>
    </xsl:template>

    <xsl:template match="xs:simpleContent/xs:extension" mode="xsdcopy">
        <xs:extension base="{@base}">
            <xs:attributeGroup ref="ism:SecurityAttributesOptionGroup"/>
        </xs:extension>
    </xsl:template>

    <xsl:template name="main">
        <xs:schema xmlns="urn:spdx-xml:1.0" attributeFormDefault="unqualified" elementFormDefault="qualified" targetNamespace="urn:spdx-xml:1.0" version="1" xmlns:xs="http://www.w3.org/2001/XMLSchema"
            xmlns:structures="http://release.niem.gov/niem/structures/4.0/" xmlns:ism="urn:us:gov:ic:ism">
            <xs:import namespace="urn:us:gov:ic:ism" schemaLocation="ext/ic-xml/ic-ism.xsd"/>
            <xs:import schemaLocation="ext/niem/utility/structures/4.0/structures.xsd" namespace="http://release.niem.gov/niem/structures/4.0/"/>
            <xsl:apply-templates select="exsl:node-set($xsddata)/*" mode="xsdcopy"/>
        </xs:schema>
    </xsl:template>

</xsl:stylesheet>
