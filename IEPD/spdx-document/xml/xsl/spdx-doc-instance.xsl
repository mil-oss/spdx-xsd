<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" 
    xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./../../../../xml/xsl/xml_instance.xsl"/>
    <!-- 
    input:  ../xsd/spdx-doc.xsd
    output: ../instance/spdx-doc-test-instance.xml
   -->
    <xsl:param name="TestData" select="'../instance/spdx-doc-test-data.xml'"/>
    <xsl:param name="Root" select="'SpdxDocumentType'"/>
    
    <xsl:variable name="testData" select="document($TestData)"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <SpdxDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:spdx-seva::1.0" xsi:schemaLocation="urn:spdx-seva::1.0 ../xsd/spdx-doc-iep.xsd">
            <xsl:apply-templates select="xs:schema/xs:complexType[@name = $Root]" mode="root">
                <xsl:with-param name="testData" select="document($TestData)"/>
            </xsl:apply-templates>
        </SpdxDocument>
    </xsl:template>
    
    <!--<xsl:template match="xs:element[@ref = 'Relationship']">
        <xsl:element name="{@ref}" namespace="urn:spdx-seva::1.0"/>
    </xsl:template>-->
    <xsl:template match="xs:element[@ref = 'DescribesPackage']"/>
    <xsl:template match="xs:element[@ref = 'DescribesFile']"/>
    <xsl:template match="xs:element[@ref = 'ExternalDocumentRef']"/>
    <xsl:template match="xs:element[@ref = 'HasExtractedLicensingInfo']"/>
    
    <xsl:template match="xs:element[@ref = 'RelatedSpdxElement'and ancestor::xs:complexType[@name='RelationshipType']]"/>
   <!-- <xsl:template match="xs:complexType[@name='RelationshipType']/*/xs:element[@ref = 'RelatedSpdxElement']"/>-->
    <!--<xsl:template match="xs:element[@ref = 'DataLicense']">
        <xsl:element name="{@ref}" namespace="urn:spdx-seva::1.0"/>
    </xsl:template>-->
    
    

    

</xsl:stylesheet>
