<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="xml" indent="yes"/>

    <xsl:include href="./common/xml-instance.xsl"/>
    <!-- 
    input:  ${pdu}/spdx-xsd/IEPD/spdx-seva/xml/xsd/spdx-seva-iep.xsd
    output: ${pdu}/spdx-xsd/IEPD/spdx-seva/xml/instance/spdx-seva-instance.xml
   -->

    <xsl:variable name="TestData" select="'../instance/spdx-seva-test_data.xml'"/>
    <xsl:param name="Root" select="'SoftwareEvidenceArchiveType'"/>

    <xsl:template match="/">
        <xsl:call-template name="main"/>
    </xsl:template>

    <xsl:template name="main">
        <SoftwareEvidenceArchive xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:spdx-seva::1.0" xsi:schemaLocation="urn:spdx-seva::1.0 ../xsd/spdx-seva-iep.xsd">
            <xsl:apply-templates select="xs:schema/xs:complexType[@name = $Root]" mode="root">
                <xsl:with-param name="testData" select="document($TestData)"/>
            </xsl:apply-templates>
        </SoftwareEvidenceArchive>
    </xsl:template>

</xsl:stylesheet>
