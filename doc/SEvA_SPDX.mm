<map version="freeplane 1.3.0">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="SEvA-SPDX" ID="ID_953746578" CREATED="1533147714559" MODIFIED="1537832444339">
<font SIZE="18"/>
<hook NAME="MapStyle">
    <properties fit_to_viewport="false" edgeColorConfiguration="#808080ff,#ff0000ff,#0000ffff,#00ff00ff,#ff00ffff,#00ffffff,#7c0000ff,#00007cff,#007c00ff,#7c007cff,#007c7cff,#7c7c00ff"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node" UNIFORM_SHAPE="true" VGAP_QUANTITY="24.0 pt">
<font SIZE="24"/>
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="default" COLOR="#000000" STYLE="fork" ICON_SIZE="12.0 pt">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.attributes">
<font SIZE="9"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.note" COLOR="#000000" BACKGROUND_COLOR="#ffffff" TEXT_ALIGN="LEFT"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000" SHAPE_HORIZONTAL_MARGIN="10.0 pt" SHAPE_VERTICAL_MARGIN="10.0 pt">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font SIZE="16"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font SIZE="14"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font SIZE="12"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font SIZE="10"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,5"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,6"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,7"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,8"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,9"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,10"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,11"/>
</stylenode>
</stylenode>
</map_styles>
</hook>
<hook NAME="AutomaticEdgeColor" COUNTER="9"/>
<node TEXT="Background" POSITION="right" ID="ID_1813150118" CREATED="1533147927551" MODIFIED="1537831888675" HGAP_QUANTITY="34.99999937415125 pt" VSHIFT_QUANTITY="120.74999640136969 pt" VSHIFT="40">
<edge COLOR="#ff0000"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      This paper outlines the implementation of SPDX as XML Schema and the integration of Software Evidence Archive (SEvA) information with the Software Product Documentation Exchange (SPDX) data model using XML Schema.
    </p>
  </body>
</html>

</richcontent>
<node TEXT="Software Product Documentation Exchange (SPDX)" ID="ID_1600736025" CREATED="1533158017636" MODIFIED="1537829839565" HGAP_QUANTITY="36.49999932944776 pt" VSHIFT_QUANTITY="6.749999798834326 pt" VSHIFT="-20"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      The SPDX Specification is available as a data model defined using the Resource Definition Framework (RDF) vocabulary.Purpose.&#160;&#160;It is primarily used for License information, but the addition of software supply chain information is proposed.&#160;
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="Software Evidence Archive (SEvA)" ID="ID_1269404110" CREATED="1537829662257" MODIFIED="1537830607711"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      SEvA provides a resource for code identification, content, dependencies, vulnerabilities, and other assets in the Software Supply Chain.&#160;&#160;It is used to create audit reports on software, which can include sensitive information&#160;&#160;which requires validation for content and format using XML Schema.
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
<node TEXT="Purpose" POSITION="right" ID="ID_950180714" CREATED="1533158017636" MODIFIED="1537831906966" HGAP_QUANTITY="36.49999932944776 pt" VSHIFT_QUANTITY="6.749999798834326 pt" HGAP="50" VSHIFT="-30">
<edge COLOR="#00ffff"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Integrate Software Evidence Archive (SEvA) Information into the Software Product Documentation Exchange (SPDX) format.
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="Challenges" POSITION="right" ID="ID_1131900390" CREATED="1533148203923" MODIFIED="1537832452696" HGAP_QUANTITY="36.499999329447775 pt" VSHIFT_QUANTITY="6.749999798834333 pt" HGAP="50" VSHIFT="-70">
<edge COLOR="#0000ff"/>
<node TEXT="Automation" ID="ID_13780163" CREATED="1533148236291" MODIFIED="1537832452695" HGAP_QUANTITY="25.99999964237214 pt" VSHIFT_QUANTITY="108.74999675899755 pt" VSHIFT="30"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      XML Schema enable validation of required data structures as part of a software supply chain data provenance testing cycle.
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="Authoritative Data" ID="ID_541477447" CREATED="1533148248860" MODIFIED="1537831945028" HGAP_QUANTITY="25.249999664723887 pt" VSHIFT_QUANTITY="5.249999843537813 pt" VSHIFT="-10"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      The core authoritative data for SEvA is contained in an XML Schema.&#160;&#160;SPDX does not have an XML Schema representation, but one has been generated and is proposed as a normative option for SPDX implementers.
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
<node TEXT="Design" POSITION="right" ID="ID_1441000239" CREATED="1533150316687" MODIFIED="1537832449433" HGAP_QUANTITY="38.74999926239255 pt" VSHIFT_QUANTITY="8.999999731779099 pt" HGAP="60" VSHIFT="-40">
<edge COLOR="#00ff00"/>
<node TEXT="National Information Exchange Model (NIEM)" ID="ID_822029341" CREATED="1533150385042" MODIFIED="1537832446945" HGAP_QUANTITY="16.249999932944778 pt" VSHIFT_QUANTITY="92.9999972283841 pt" VSHIFT="40"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      NIEM&#160;&#160;XML Schema Naming and Design rules were employed in the creation of XML Schema for SEvA and SPDX.&#160;&#160;XML Schema based implementations
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="SPDX" ID="ID_967173456" CREATED="1533152824418" MODIFIED="1537832449432" HGAP_QUANTITY="16.249999932944778 pt" VSHIFT_QUANTITY="-93.74999720603236 pt" VSHIFT="-20"><richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      XML Schema for SPDX has been generated for SPDX, and License, Document, and SEvA information exchanges have been created using the NIEM Information Exchange Product Documentation (IEPD) methodology.
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
<node TEXT="Control Markings" POSITION="right" ID="ID_178890338" CREATED="1533158758424" MODIFIED="1537832444339" HGAP_QUANTITY="37.249999307096026 pt" VSHIFT_QUANTITY="4.499999865889541 pt" VSHIFT="-40">
<edge COLOR="#7c0000"/>
</node>
<node TEXT="Recommendations" POSITION="right" ID="ID_1433935904" CREATED="1533159235193" MODIFIED="1537832432632" HGAP_QUANTITY="37.24999930709602 pt" VSHIFT_QUANTITY="-6.749999798834301 pt">
<edge COLOR="#00007c"/>
</node>
<node TEXT="Conclusion" POSITION="right" ID="ID_1336118719" CREATED="1533167043972" MODIFIED="1537832441948" HGAP_QUANTITY="40.99999919533731 pt" VSHIFT_QUANTITY="-138.74999586492788 pt" VSHIFT="40">
<edge COLOR="#007c00"/>
</node>
<node TEXT="_Fileout" POSITION="left" ID="ID_381249361" CREATED="1475710702062" MODIFIED="1537828830109">
<edge COLOR="#0000ff"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      /home/jdn/DATA/Neushul_Solutions/Projects/SEvA/IonChannel/spdx-xsd/doc/SeVA-SPDX.html
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="_Subj" POSITION="left" ID="ID_1130198740" CREATED="1475711301439" MODIFIED="1537828897629" VSHIFT_QUANTITY="-20.0 px">
<edge COLOR="#00ffff"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Integrating Software Security Information Into The SPDX Standard
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="_Date" POSITION="left" ID="ID_299078720" CREATED="1475710798573" MODIFIED="1537829043649">
<edge COLOR="#00ff00"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      24 Sept 2018
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="_Title" POSITION="left" ID="ID_1783696170" CREATED="1475710823436" MODIFIED="1537828905608">
<edge COLOR="#ff00ff"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      SEvA SPDX Integration
    </p>
  </body>
</html>

</richcontent>
</node>
<node TEXT="_bash" POSITION="left" ID="ID_1253692164" CREATED="1513800134836" MODIFIED="1537828979974">
<edge COLOR="#ff0000"/>
<richcontent TYPE="DETAILS">

<html>
  <head>
    
  </head>
  <body>
    <p>
      #!/bin/bash
    </p>
    <p>
      java -cp /home/jdn/DATA/Neushul_Solutions/Projects/XML/saxon9he.jar net.sf.saxon.Transform -s:/home/jdn/DATA/Neushul_Solutions/Projects/SEvA/IonChannel/spdx-xsd/doc/SEvA_SPDX.mm -xsl:/home/jdn/DATA/Neushul_Solutions/Projects/XML/MM_Naval_Ltr_Format.xsl
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
</map>
