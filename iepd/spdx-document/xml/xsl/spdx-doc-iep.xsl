<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:exsl="http://exslt.org/common" version="1.0">

    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./common/iep.xsl"/>

    <!--<xsl:variable name="spdx_xsd" select="document('../xsd/spdx-xml-ref.xsd')"/>-->

    <xsl:variable name="Top" select="'SpdxDocumentType'"/>
    <xsl:variable name="Super" select="'SpdxElementType'"/>
    <xsl:variable name="Root" select="'SpdxDocumentType'"/>
    <xsl:variable name="RootEl" select="'SpdxDocument'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:variable name="xsddata">
        <xs:annotation>
            <xs:documentation>XML Schema for SPDX Document Information Exchange</xs:documentation>
            <xs:appinfo>
                <Root type="{$Root}" name="{$RootEl}"/>
            </xs:appinfo>
        </xs:annotation>
        <xsl:apply-templates select="//xs:schema/*[@name = $Root]"/>
        <xsl:variable name="allnodes">
            <xsl:call-template name="deDupList">
                <xsl:with-param name="list">
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Root]"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $RootEl]"/>
                        <xsl:with-param name="iteration" select="10"/>
                    </xsl:call-template>
                    <xsl:call-template name="iterateNode">
                        <xsl:with-param name="node" select="//xs:schema/*[@name = $Root]"/>
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
        <xsl:for-each select="exsl:node-set($allnodes)/xs:attribute">
            <xsl:sort select="@name"/>
            <xsl:copy-of select="."/>
        </xsl:for-each>
    </xsl:variable>
    
    <xsl:template name="main">
        <xsl:call-template name="makeXSD">
            <xsl:with-param name="xsdnodes" select="$xsddata"/>
        </xsl:call-template>
    </xsl:template>
    
    <!--<xsl:template match="xs:element[@name = 'DataLicense']" mode="xsdcopy">
       <xsl:element name="{name()}">
           <xsl:apply-templates select="@*" mode="xsdcopy"/>
           <xsl:attribute name="fixed">http://spdx.org/licenses/CC0-1.0</xsl:attribute>
           <xsl:apply-templates select="*" mode="xsdcopy"/>
       </xsl:element>
    </xsl:template>
    
    <xsl:template match="xs:element[@ref = 'RelatedSpdxElement'and ancestor::xs:complexType[@name='RelationshipType']]/@minOccurs" mode="xsdcopy">
        <xsl:attribute name="minOccurs">
            <xsl:text>0</xsl:text>
        </xsl:attribute>
    </xsl:template>
-->
</xsl:stylesheet>