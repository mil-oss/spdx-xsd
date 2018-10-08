<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:exsl="http://exslt.org/common" xmlns:xs="http://www.w3.org/2001/XMLSchema" exclude-result-prefixes="xs" version="1.0">
    <xsl:output method="text" indent="yes"/>

    <xsl:include href="./../../../../xml/xsl/go-gen.xsl"/>

    <!-- 
    input: ${pdu}/spdx-xsd/IEPD/spdx-document/xml/xsd/spdx-doc-iep.xsd
    output:${pdu}/spdx-xsd/src/spdx-doc/spdx-doc-struct.go
   -->

    <xsl:variable name="rootname" select="xs:schema/xs:annotation/xs:appinfo/Root/@type"/>

    <xsl:template match="/">
        <xsl:call-template name="makego">
            <xsl:with-param name="rootname" select="xs:schema/xs:annotation/xs:appinfo/Root/@type"/>
        </xsl:call-template>
    </xsl:template>

</xsl:stylesheet>
