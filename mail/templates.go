package mail

type SmtpTemplateData struct {
	Subject string
	Body    string
}

const EmailTemplate = `MIME-version: 1.0;
Content-Type: text/html;
charset: "UTF-8";
Subject: {{.Subject}}

<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"
>
<head>
    <!--[if gte mso 9]>
    <xml>
        <o:OfficeDocumentSettings>
            <o:AllowPNG/>
            <o:PixelsPerInch>96</o:PixelsPerInch>
        </o:OfficeDocumentSettings>
    </xml><![endif]-->
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width">
    <!--[if !mso]><!-->
    <meta http-equiv="X-UA-Compatible" content="IE=edge"><!--<![endif]-->
    <title>Template Base</title>
    <!--[if !mso]><!-- -->
    <link href="https://fonts.googleapis.com/css?family=Lato" rel="stylesheet" type="text/css">
    <!--<![endif]-->

    <style type="text/css" id="media-query">
        body {
            margin: 0;
            padding: 0;
        }

        table, tr, td {
            vertical-align: top;
            border-collapse: collapse;
        }

        .ie-browser table, .mso-container table {
            table-layout: fixed;
        }

        * {
            line-height: inherit;
        }

        a[x-apple-data-detectors=true] {
            color: inherit !important;
            text-decoration: none !important;
        }

        [owa] .img-container div, [owa] .img-container button {
            display: block !important;
        }

        [owa] .fullwidth button {
            width: 100% !important;
        }

        [owa] .block-grid .col {
            display: table-cell;
            float: none !important;
            vertical-align: top;
        }

        .ie-browser .num12, .ie-browser .block-grid, [owa] .num12, [owa] .block-grid {
            width: 500px !important;
        }

        .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div {
            line-height: 100%;
        }

        .ie-browser .mixed-two-up .num4, [owa] .mixed-two-up .num4 {
            width: 164px !important;
        }

        .ie-browser .mixed-two-up .num8, [owa] .mixed-two-up .num8 {
            width: 328px !important;
        }

        .ie-browser .block-grid.two-up .col, [owa] .block-grid.two-up .col {
            width: 250px !important;
        }

        .ie-browser .block-grid.three-up .col, [owa] .block-grid.three-up .col {
            width: 166px !important;
        }

        .ie-browser .block-grid.four-up .col, [owa] .block-grid.four-up .col {
            width: 125px !important;
        }

        .ie-browser .block-grid.five-up .col, [owa] .block-grid.five-up .col {
            width: 100px !important;
        }

        .ie-browser .block-grid.six-up .col, [owa] .block-grid.six-up .col {
            width: 83px !important;
        }

        .ie-browser .block-grid.seven-up .col, [owa] .block-grid.seven-up .col {
            width: 71px !important;
        }

        .ie-browser .block-grid.eight-up .col, [owa] .block-grid.eight-up .col {
            width: 62px !important;
        }

        .ie-browser .block-grid.nine-up .col, [owa] .block-grid.nine-up .col {
            width: 55px !important;
        }

        .ie-browser .block-grid.ten-up .col, [owa] .block-grid.ten-up .col {
            width: 50px !important;
        }

        .ie-browser .block-grid.eleven-up .col, [owa] .block-grid.eleven-up .col {
            width: 45px !important;
        }

        .ie-browser .block-grid.twelve-up .col, [owa] .block-grid.twelve-up .col {
            width: 41px !important;
        }

        @media only screen and (min-width: 520px) {
            .block-grid {
                width: 500px !important;
            }

            .block-grid .col {
                display: table-cell;
                Float: none !important;
                vertical-align: top;
            }

            .block-grid .col.num12 {
                width: 500px !important;
            }

            .block-grid.mixed-two-up .col.num4 {
                width: 164px !important;
            }

            .block-grid.mixed-two-up .col.num8 {
                width: 328px !important;
            }

            .block-grid.two-up .col {
                width: 250px !important;
            }

            .block-grid.three-up .col {
                width: 166px !important;
            }

            .block-grid.four-up .col {
                width: 125px !important;
            }

            .block-grid.five-up .col {
                width: 100px !important;
            }

            .block-grid.six-up .col {
                width: 83px !important;
            }

            .block-grid.seven-up .col {
                width: 71px !important;
            }

            .block-grid.eight-up .col {
                width: 62px !important;
            }

            .block-grid.nine-up .col {
                width: 55px !important;
            }

            .block-grid.ten-up .col {
                width: 50px !important;
            }

            .block-grid.eleven-up .col {
                width: 45px !important;
            }

            .block-grid.twelve-up .col {
                width: 41px !important;
            }
        }

        @media (max-width: 520px) {
            .block-grid, .col {
                min-width: 320px !important;
                max-width: 100% !important;
            }

            .block-grid {
                width: calc(100% - 40px) !important;
            }

            .col {
                width: 100% !important;
            }

            .col > div {
                margin: 0 auto;
            }

            img.fullwidth {
                max-width: 100% !important;
            }
        }

    </style>
</head>
<body class="clean-body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #E6E6E6">

<!--[if IE]>
<div class="ie-browser"><![endif]-->
<!--[if mso]>
<div class="mso-container"><![endif]-->
<div class="nl-container" style="min-width: 320px;Margin: 0 auto;background-color: #E6E6E6">
    <!--[if (mso)|(IE)]>
    <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
            <td align="center" style="background-color: #E6E6E6;"><![endif]-->

    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="500"
                    style=" width:500px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 500px;width: calc(18000% - 89500px);background-color: transparent;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div style="padding-right: 20px; padding-left: 20px; padding-top: 20px; padding-bottom: 20px;">
                                <!--[if (mso)]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 20px;padding-left: 20px; padding-top: 20px; padding-bottom: 20px;">
                                            <table width="100%" align="center" cellpadding="0" cellspacing="0"
                                                   border="0">
                                                <tr>
                                                    <td><![endif]-->
                                <div align="center">
                                    <div style="border-top: 0px solid transparent; width:100%; line-height:0px; height:0px; font-size:0px;">
                                        &#160;
                                    </div>
                                </div>
                                <!--[if (mso)]></td></tr></table></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="500"
                    style=" width:500px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 500px;width: calc(18000% - 89500px);background-color: transparent;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div align="center" class="img-container center fullwidth"
                                 style="padding-right: 0px;  padding-left: 0px;">
                                <!--[if mso]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 0px; padding-left: 0px;" align="center"><![endif]-->
                                <img class="center fullwidth" align="center" border="0"
                                     src="https://firebasestorage.googleapis.com/v0/b/wwm-itm.appspot.com/o/borderup.png?alt=media&token=547fbc16-578b-43ad-a708-d32185077042"
                                     alt="Image" title="Image"
                                     style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: block !important;border: 0;height: auto;float: none;width: 100%;max-width: 500px"
                                     width="500">
                                <!--[if mso]></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: #FFFFFF;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:#FFFFFF;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="498"
                    style=" width:498px; padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 498px;width: calc(18000% - 89500px);background-color: #FFFFFF;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div align="center" class="img-container center"
                                 style="padding-right: 0px;  padding-left: 0px;">
                                <!--[if mso]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 0px; padding-left: 0px;" align="center"><![endif]-->
                                <img class="center" align="center" border="0"
                                     src="https://avatars1.githubusercontent.com/u/15645275?v=4&amp;s=200"
                                     alt="Weisswurst-Systems Logo" title="Weisswurst-Systems Logo"
                                     style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: block !important;border: 0;height: auto;float: none;width: 100%;max-width: 200px"
                                     width="200">
                                <!--[if mso]></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: #FFFFFF;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:#FFFFFF;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="498"
                    style=" width:498px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 498px;width: calc(18000% - 89500px);background-color: #FFFFFF;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF; padding-top:0px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <!--[if mso]>
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td style="padding-right: 15px; padding-left: 15px; padding-top: 15px; padding-bottom: 15px;">
                            <![endif]-->
                            <div style="font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;color:#555555;line-height:120%; padding-right: 15px; padding-left: 15px; padding-top: 15px; padding-bottom: 15px;">
                                <div style="font-size:12px;line-height:14px;color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;text-align:left;">
                                    <p style="margin: 0;font-size: 12px;line-height: 14px;text-align: center"><a
                                            style="color:#67CCDE;text-decoration: underline;"
                                            href="https://weisswurstsystems.github.io/WWM-ITM/" target="_blank"
                                            rel="noopener noreferrer">Weisswurstverwaltung</a>&#160; &#160; | &#160; <a
                                            style="color:#67CCDE;text-decoration: underline;"
                                            href="https://github.com/WeisswurstSystems/WWM-ITM/issues" target="_blank"
                                            rel="noopener noreferrer">Probleme? Melde sie hier!</a><br></p></div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: #FFFFFF;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:#FFFFFF;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="498"
                    style=" width:498px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px; border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 498px;width: calc(18000% - 89500px);background-color: #FFFFFF;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <!--[if mso]>
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                            <![endif]-->
                            <div style="color:#555555;line-height:120%;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif; padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                <div style="font-size:12px;line-height:14px;text-align:center;color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;">
                                    <p style="margin: 0;font-size: 12px;line-height: 14px;text-align: center"><span
                                            style="font-size: 18px; line-height: 21px;"><strong><span
                                            style="line-height: 21px; font-size: 18px;">{{.Subject}}</span></strong></span>
                                    </p></div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->


                            <!--[if mso]>
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                            <![endif]-->
                            <div style="font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;color:#555555;line-height:150%; padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                <div style="font-size:12px;line-height:18px;text-align:center;color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;">
                                    <p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center">
                                        {{.Body}}</p></div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->


                            <!--[if mso]>
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                            <![endif]-->
                            <div style="font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;color:#555555;line-height:150%; padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                <div style="font-size:12px;line-height:18px;color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;text-align:left;">
                                    <p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center"><strong>Liebe
                                        Grüße!</strong></p>
                                    <p style="margin: 0;font-size: 12px;line-height: 18px;text-align: center">von&#160;Weisswurst
                                        Systems</p></div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->


                            <div style="padding-right: 10px; padding-left: 10px; padding-top: 15px; padding-bottom: 10px;">
                                <!--[if (mso)]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 10px;padding-left: 10px; padding-top: 15px; padding-bottom: 10px;">
                                            <table width="35%" align="center" cellpadding="0" cellspacing="0"
                                                   border="0">
                                                <tr>
                                                    <td><![endif]-->
                                <div align="center">
                                    <div style="border-top: 6px dotted #E6E6E6; width:35%; line-height:6px; height:6px; font-size:6px;">
                                        &#160;
                                    </div>
                                </div>
                                <!--[if (mso)]></td></tr></table></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: #FFFFFF;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:#FFFFFF;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="498"
                    style=" width:498px; padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 498px;width: calc(18000% - 89500px);background-color: #FFFFFF;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 1px solid #E2DFDF; border-bottom: 0px solid transparent; border-right: 1px solid #E2DFDF; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <!--[if mso]>
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                            <![endif]-->
                            <div style="font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;color:#67CCDE;line-height:150%; padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                <div style="font-size:12px;line-height:18px;color:#67CCDE;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;text-align:left;">
                                    <p style="margin: 0;font-size: 11px;line-height: 18px;text-align: center">
                                <span style="font-size: 11px; line-height: 16px;">Du möchtest keine E-Mails mehr von uns empfangen? <br> Du kannst E-Mail Benachrichtigungen einfach in den Profileinstellungen <br> der Weisswurstverwaltung deaktivieren!</span>
                                    </p>
                                </div>
                            </div>
                            <!--[if mso]></td></tr></table><![endif]-->


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="500"
                    style=" width:500px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 500px;width: calc(18000% - 89500px);background-color: transparent;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div align="center" class="img-container center fullwidth"
                                 style="padding-right: 0px;  padding-left: 0px;">
                                <!--[if mso]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 0px; padding-left: 0px;" align="center"><![endif]-->
                                <img class="center fullwidth" align="center" border="0"
                                     src="https://firebasestorage.googleapis.com/v0/b/wwm-itm.appspot.com/o/borderdown.png?alt=media&token=13a27ffb-6611-4841-919f-66466c3218d0"
                                     alt="Image" title="Image"
                                     style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: block !important;border: 0;height: auto;float: none;width: 100%;max-width: 500px"
                                     width="500">
                                <!--[if mso]></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="500"
                    style=" width:500px; padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 500px;width: calc(18000% - 89500px);background-color: transparent;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div style="padding-right: 25px; padding-left: 25px; padding-top: 25px; padding-bottom: 25px;">
                                <!--[if (mso)]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 25px;padding-left: 25px; padding-top: 25px; padding-bottom: 25px;">
                                            <table width="100%" align="center" cellpadding="0" cellspacing="0"
                                                   border="0">
                                                <tr>
                                                    <td><![endif]-->
                                <div align="center">
                                    <div style="border-top: 0px solid transparent; width:100%; line-height:0px; height:0px; font-size:0px;">
                                        &#160;
                                    </div>
                                </div>
                                <!--[if (mso)]></td></tr></table></td></tr></table><![endif]-->
                            </div>


                            <div style="padding-right: 15px; padding-left: 15px; padding-top: 15px; padding-bottom: 15px;">
                                <!--[if (mso)]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 15px;padding-left: 15px; padding-top: 15px; padding-bottom: 15px;">
                                            <table width="100%" align="center" cellpadding="0" cellspacing="0"
                                                   border="0">
                                                <tr>
                                                    <td><![endif]-->
                                <div align="center">
                                    <div style="border-top: 0px solid transparent; width:100%; line-height:0px; height:0px; font-size:0px;">
                                        &#160;
                                    </div>
                                </div>
                                <!--[if (mso)]></td></tr></table></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <div style="background-color:transparent;">
        <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;width: 500px;width: calc(19000% - 98300px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;"
             class="block-grid ">
            <div style="border-collapse: collapse;display: table;width: 100%;">
                <!--[if (mso)|(IE)]>
                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                    <tr>
                        <td style="background-color:transparent;" align="center">
                            <table cellpadding="0" cellspacing="0" border="0" style="width: 500px;">
                                <tr class="layout-full-width" style="background-color:transparent;"><![endif]-->

                <!--[if (mso)|(IE)]>
                <td align="center" width="500"
                    style=" width:500px; padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;"
                    valign="top"><![endif]-->
                <div class="col num12"
                     style="min-width: 320px;max-width: 500px;width: 500px;width: calc(18000% - 89500px);background-color: transparent;">
                    <div style="background-color: transparent; width: 100% !important;">
                        <!--[if (!mso)&(!IE)]><!-->
                        <div style="border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                            <!--<![endif]-->


                            <div style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                <!--[if (mso)]>
                                <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                    <tr>
                                        <td style="padding-right: 10px;padding-left: 10px; padding-top: 10px; padding-bottom: 10px;">
                                            <table width="100%" align="center" cellpadding="0" cellspacing="0"
                                                   border="0">
                                                <tr>
                                                    <td><![endif]-->
                                <div align="center">
                                    <div style="border-top: 0px solid transparent; width:100%; line-height:0px; height:0px; font-size:0px;">
                                        &#160;
                                    </div>
                                </div>
                                <!--[if (mso)]></td></tr></table></td></tr></table><![endif]-->
                            </div>


                            <!--[if (!mso)&(!IE)]><!--></div><!--<![endif]-->
                    </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
            </div>
        </div>
    </div>
    <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
</div>
<!--[if (mso)|(IE)]></div><![endif]-->


</body>
</html>
`
