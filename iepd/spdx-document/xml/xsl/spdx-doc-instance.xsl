<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./common/xml-instance.xsl"/>
    <!-- 
    input:  ${pdu}/spdx-xsd/IEPD/spdx-security/xml/xsd/spdx-security-iep.xsd
    output: ${pdu}/spdx-xsd/IEPD/spdx-security/xml/instance/spdx-security-instance.xml
   -->

    <xsl:variable name="TestData" select="'../instance/spdx-doc-test-data.xml'"/>
    <xsl:param name="Root" select="'SpdxDocumentType'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <SpdxDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xmlns="urn:spdx-xml:1.0"
            xsi:schemaLocation="urn:spdx-xml:1.0 ../xsd/spdx-doc-iep.xsd">
           <xsl:apply-templates select="xs:schema/xs:complexType[@name = $Root]" mode="root">
               <xsl:with-param name="testData" select="document($TestData)"/>
            </xsl:apply-templates>
        </SpdxDocument>
    </xsl:template>

</xsl:stylesheet>
