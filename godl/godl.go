package godl

import (
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	humanize "github.com/dustin/go-humanize"

	"time"

	pb "gopkg.in/cheggaaa/pb.v1"
)

var mimesCons = func() map[string]string {
	return map[string]string{
		"application/vnd.hzn-3d-crossword":                            ".x3d",
		"video/3gpp":                                                  ".3gp",
		"video/3gpp2":                                                 ".3g2",
		"application/vnd.mseq":                                        ".mseq",
		"application/vnd.3m.post-it-notes":                            ".pwn",
		"application/vnd.3gpp.pic-bw-large":                           ".plb",
		"application/vnd.3gpp.pic-bw-small":                           ".psb",
		"application/vnd.3gpp.pic-bw-var":                             ".pvb",
		"application/vnd.3gpp2.tcap":                                  ".tcap",
		"application/x-7z-compressed":                                 ".7z",
		"application/x-abiword":                                       ".abw",
		"application/x-ace-compressed":                                ".ace",
		"application/vnd.americandynamics.acc":                        ".acc",
		"application/vnd.acucobol":                                    ".acu",
		"application/vnd.acucorp":                                     ".atc",
		"audio/adpcm":                                                 ".adp",
		"application/x-authorware-bin":                                ".aab",
		"application/x-authorware-map":                                ".aam",
		"application/x-authorware-seg":                                ".aas",
		"application/vnd.adobe.air-application-installer-package+zip": ".air",
		"application/x-shockwave-flash":                               ".swf",
		"application/vnd.adobe.fxp":                                   ".fxp",
		"application/pdf":                                             ".pdf",
		"application/vnd.cups-ppd":                                    ".ppd",
		"application/x-director":                                      ".dir",
		"application/vnd.adobe.xdp+xml":                               ".xdp",
		"application/vnd.adobe.xfdf":                                  ".xfdf",
		"audio/x-aac":                                                 ".aac",
		"application/vnd.ahead.space":                                 ".ahead",
		"application/vnd.airzip.filesecure.azf":                       ".azf",
		"application/vnd.airzip.filesecure.azs":                       ".azs",
		"application/vnd.amazon.ebook":                                ".azw",
		"application/vnd.amiga.ami":                                   ".ami",
		"application/andrew-inset":                                    "N/A",
		"application/vnd.android.package-archive":                     ".apk",
		"application/vnd.anser-web-certificate-issue-initiation":      ".cii",
		"application/vnd.anser-web-funds-transfer-initiation":         ".fti",
		"application/vnd.antix.game-component":                        ".atx",
		"application/x-apple-diskimage":                               ".dmg",
		"application/vnd.apple.installer+xml":                         ".mpkg",
		"application/applixware":                                      ".aw",
		"application/vnd.hhe.lesson-player":                           ".les",
		"application/vnd.aristanetworks.swi":                          ".swi",
		"text/x-asm":                                                  ".s",
		"application/atomcat+xml":                                     ".atomcat",
		"application/atomsvc+xml":                                     ".atomsvc",
		"application/atom+xml":                                        ".atom, .xml",
		"application/pkix-attr-cert":                                  ".ac",
		"audio/x-aiff":                                                ".aif",
		"video/x-msvideo":                                             ".avi",
		"application/vnd.audiograph":                                  ".aep",
		"image/vnd.dxf":                                               ".dxf",
		"model/vnd.dwf":                                               ".dwf",
		"text/plain-bas":                                              ".par",
		"application/x-bcpio":                                         ".bcpio",
		"application/octet-stream":                                    ".bin",
		"image/bmp":                                                   ".bmp",
		"application/x-bittorrent":                                    ".torrent",
		"application/vnd.rim.cod":                                     ".cod",
		"application/vnd.blueice.multipass":                           ".mpm",
		"application/vnd.bmi":                                         ".bmi",
		"application/x-sh":                                            ".sh",
		"image/prs.btif":                                              ".btif",
		"application/vnd.businessobjects":                             ".rep",
		"application/x-bzip":                                          ".bz",
		"application/x-bzip2":                                         ".bz2",
		"application/x-csh":                                           ".csh",
		"text/x-c":                                                    ".c",
		"application/vnd.chemdraw+xml":                                ".cdxml",
		"text/css":                                                    ".css",
		"chemical/x-cdx":                                              ".cdx",
		"chemical/x-cml":                                              ".cml",
		"chemical/x-csml":                                             ".csml",
		"application/vnd.contact.cmsg":                                ".cdbcmsg",
		"application/vnd.claymore":                                    ".cla",
		"application/vnd.clonk.c4group":                               ".c4g",
		"image/vnd.dvb.subtitle":                                      ".sub",
		"application/cdmi-capability":                                 ".cdmia",
		"application/cdmi-container":                                  ".cdmic",
		"application/cdmi-domain":                                     ".cdmid",
		"application/cdmi-object":                                     ".cdmio",
		"application/cdmi-queue":                                      ".cdmiq",
		"application/vnd.cluetrust.cartomobile-config":                ".c11amc",
		"application/vnd.cluetrust.cartomobile-config-pkg":            ".c11amz",
		"image/x-cmu-raster":                                          ".ras",
		"model/vnd.collada+xml":                                       ".dae",
		"text/csv":                                                    ".csv",
		"application/mac-compactpro":                                  ".cpt",
		"application/vnd.wap.wmlc":                                    ".wmlc",
		"image/cgm":                                                   ".cgm",
		"x-conference/x-cooltalk":                                     ".ice",
		"image/x-cmx":                                                 ".cmx",
		"application/vnd.xara":                                        ".xar",
		"application/vnd.cosmocaller":                                 ".cmc",
		"application/x-cpio":                                          ".cpio",
		"application/vnd.crick.clicker":                               ".clkx",
		"application/vnd.crick.clicker.keyboard":                      ".clkk",
		"application/vnd.crick.clicker.palette":                       ".clkp",
		"application/vnd.crick.clicker.template":                      ".clkt",
		"application/vnd.crick.clicker.wordbank":                      ".clkw",
		"application/vnd.criticaltools.wbs+xml":                       ".wbs",
		"application/vnd.rig.cryptonote":                              ".cryptonote",
		"chemical/x-cif":                                              ".cif",
		"chemical/x-cmdf":                                             ".cmdf",
		"application/cu-seeme":                                        ".cu",
		"application/prs.cww":                                         ".cww",
		"text/vnd.curl":                                               ".curl",
		"text/vnd.curl.dcurl":                                         ".dcurl",
		"text/vnd.curl.mcurl":                                         ".mcurl",
		"text/vnd.curl.scurl":                                         ".scurl",
		"application/vnd.curl.car":                                    ".car",
		"application/vnd.curl.pcurl":                                  ".pcurl",
		"application/vnd.yellowriver-custom-menu":                     ".cmp",
		"application/dssc+der":                                        ".dssc",
		"application/dssc+xml":                                        ".xdssc",
		"application/x-debian-package":                                ".deb",
		"audio/vnd.dece.audio":                                        ".uva",
		"image/vnd.dece.graphic":                                      ".uvi",
		"video/vnd.dece.hd":                                           ".uvh",
		"video/vnd.dece.mobile":                                       ".uvm",
		"video/vnd.uvvu.mp4":                                          ".uvu",
		"video/vnd.dece.pd":                                           ".uvp",
		"video/vnd.dece.sd":                                           ".uvs",
		"video/vnd.dece.video":                                        ".uvv",
		"application/x-dvi":                                           ".dvi",
		"application/vnd.fdsn.seed":                                   ".seed",
		"application/x-dtbook+xml":                                    ".dtb",
		"application/x-dtbresource+xml":                               ".res",
		"application/vnd.dvb.ait":                                     ".ait",
		"application/vnd.dvb.service":                                 ".svc",
		"audio/vnd.digital-winds":                                     ".eol",
		"image/vnd.djvu":                                              ".djvu",
		"application/xml-dtd":                                         ".dtd",
		"application/vnd.dolby.mlp":                                   ".mlp",
		"application/x-doom":                                          ".wad",
		"application/vnd.dpgraph":                                     ".dpg",
		"audio/vnd.dra":                                               ".dra",
		"application/vnd.dreamfactory":                                ".dfac",
		"audio/vnd.dts":                                               ".dts",
		"audio/vnd.dts.hd":                                            ".dtshd",
		"image/vnd.dwg":                                               ".dwg",
		"application/vnd.dynageo":                                     ".geo",
		"application/ecmascript":                                      ".es",
		"application/vnd.ecowin.chart":                                ".mag",
		"image/vnd.fujixerox.edmics-mmr":                              ".mmr",
		"image/vnd.fujixerox.edmics-rlc":                              ".rlc",
		"application/exi":                                             ".exi",
		"application/vnd.proteus.magazine":                            ".mgz",
		"application/epub+zip":                                        ".epub",
		"message/rfc822":                                              ".eml",
		"application/vnd.enliven":                                     ".nml",
		"application/vnd.is-xpr":                                      ".xpr",
		"image/vnd.xiff":                                              ".xif",
		"application/vnd.xfdl":                                        ".xfdl",
		"application/emma+xml":                                        ".emma",
		"application/vnd.ezpix-album":                                 ".ez2",
		"application/vnd.ezpix-package":                               ".ez3",
		"image/vnd.fst":                                               ".fst",
		"video/vnd.fvt":                                               ".fvt",
		"image/vnd.fastbidsheet":                                      ".fbs",
		"application/vnd.denovo.fcselayout-link":                      ".fe_launch",
		"video/x-f4v":                                                 ".f4v",
		"video/x-flv":                                                 ".flv",
		"image/vnd.fpx":                                               ".fpx",
		"image/vnd.net-fpx":                                           ".npx",
		"text/vnd.fmi.flexstor":                                       ".flx",
		"video/x-fli":                                                 ".fli",
		"application/vnd.fluxtime.clip":                               ".ftc",
		"application/vnd.fdf":                                         ".fdf",
		"text/x-fortran":                                              ".f",
		"application/vnd.mif":                                         ".mif",
		"application/vnd.framemaker":                                  ".fm",
		"image/x-freehand":                                            ".fh",
		"application/vnd.fsc.weblaunch":                               ".fsc",
		"application/vnd.frogans.fnc":                                 ".fnc",
		"application/vnd.frogans.ltf":                                 ".ltf",
		"application/vnd.fujixerox.ddd":                               ".ddd",
		"application/vnd.fujixerox.docuworks":                         ".xdw",
		"application/vnd.fujixerox.docuworks.binder":                  ".xbd",
		"application/vnd.fujitsu.oasys":                               ".oas",
		"application/vnd.fujitsu.oasys2":                              ".oa2",
		"application/vnd.fujitsu.oasys3":                              ".oa3",
		"application/vnd.fujitsu.oasysgp":                             ".fg5",
		"application/vnd.fujitsu.oasysprs":                            ".bh2",
		"application/x-futuresplash":                                  ".spl",
		"application/vnd.fuzzysheet":                                  ".fzs",
		"image/g3fax":                                                 ".g3",
		"application/vnd.gmx":                                         ".gmx",
		"model/vnd.gtw":                                               ".gtw",
		"application/vnd.genomatix.tuxedo":                            ".txd",
		"application/vnd.geogebra.file":                               ".ggb",
		"application/vnd.geogebra.tool":                               ".ggt",
		"model/vnd.gdl":                                               ".gdl",
		"application/vnd.geometry-explorer":                           ".gex",
		"application/vnd.geonext":                                     ".gxt",
		"application/vnd.geoplan":                                     ".g2w",
		"application/vnd.geospace":                                    ".g3w",
		"application/x-font-ghostscript":                              ".gsf",
		"application/x-font-bdf":                                      ".bdf",
		"application/x-gtar":                                          ".gtar",
		"application/x-texinfo":                                       ".texinfo",
		"application/x-gnumeric":                                      ".gnumeric",
		"application/vnd.google-earth.kml+xml":                        ".kml",
		"application/vnd.google-earth.kmz":                            ".kmz",
		"application/vnd.grafeq":                                      ".gqf",
		"image/gif":                                                   ".gif",
		"text/vnd.graphviz":                                           ".gv",
		"application/vnd.groove-account":                              ".gac",
		"application/vnd.groove-help":                                 ".ghf",
		"application/vnd.groove-identity-message":                     ".gim",
		"application/vnd.groove-injector":                             ".grv",
		"application/vnd.groove-tool-message":                         ".gtm",
		"application/vnd.groove-tool-template":                        ".tpl",
		"application/vnd.groove-vcard":                                ".vcg",
		"video/h261":                                                  ".h261",
		"video/h263":                                                  ".h263",
		"video/h264":                                                  ".h264",
		"application/vnd.hp-hpid":                                     ".hpid",
		"application/vnd.hp-hps":                                      ".hps",
		"application/x-hdf":                                           ".hdf",
		"audio/vnd.rip":                                               ".rip",
		"application/vnd.hbci":                                        ".hbci",
		"application/vnd.hp-jlyt":                                     ".jlt",
		"application/vnd.hp-pcl":                                      ".pcl",
		"application/vnd.hp-hpgl":                                     ".hpgl",
		"application/vnd.yamaha.hv-script":                            ".hvs",
		"application/vnd.yamaha.hv-dic":                               ".hvd",
		"application/vnd.yamaha.hv-voice":                             ".hvp",
		"application/vnd.hydrostatix.sof-data":                        ".sfd-hdstx",
		"application/hyperstudio":                                     ".stk",
		"application/vnd.hal+xml":                                     ".hal",
		"text/html":                                                   ".html",
		"application/vnd.ibm.rights-management":                       ".irm",
		"application/vnd.ibm.secure-container":                        ".sc",
		"text/calendar":                                               ".ics",
		"application/vnd.iccprofile":                                  ".icc",
		"image/x-icon":                                                ".ico",
		"application/vnd.igloader":                                    ".igl",
		"image/ief":                                                   ".ief",
		"application/vnd.immervision-ivp":                             ".ivp",
		"application/vnd.immervision-ivu":                             ".ivu",
		"application/reginfo+xml":                                     ".rif",
		"text/vnd.in3d.3dml":                                          ".3dml",
		"text/vnd.in3d.spot":                                          ".spot",
		"model/iges":                                                  ".igs",
		"application/vnd.intergeo":                                    ".i2g",
		"application/vnd.cinderella":                                  ".cdy",
		"application/vnd.intercon.formnet":                            ".xpw",
		"application/vnd.isac.fcs":                                    ".fcs",
		"application/ipfix":                                           ".ipfix",
		"application/pkix-cert":                                       ".cer",
		"application/pkixcmp":                                         ".pki",
		"application/pkix-crl":                                        ".crl",
		"application/pkix-pkipath":                                    ".pkipath",
		"application/vnd.insors.igm":                                  ".igm",
		"application/vnd.ipunplugged.rcprofile":                       ".rcprofile",
		"application/vnd.irepository.package+xml":                     ".irp",
		"text/vnd.sun.j2me.app-descriptor":                            ".jad",
		"application/java-archive":                                    ".jar",
		"application/java-vm":                                         ".class",
		"application/x-java-jnlp-file":                                ".jnlp",
		"application/java-serialized-object":                          ".ser",
		"text/x-java-source,java":                                     ".java",
		"application/javascript":                                      ".js",
		"application/json":                                            ".json",
		"application/vnd.joost.joda-archive":                          ".joda",
		"video/jpm":                                                   ".jpm",
		"image/jpeg":                                                  ".jpg",
		"image/x-citrix-jpeg":                                         ".jpg",
		"image/pjpeg":                                                 ".pjpeg",
		"video/jpeg":                                                  ".jpgv",
		"application/vnd.kahootz":                                     ".ktz",
		"application/vnd.chipnuts.karaoke-mmd":                        ".mmd",
		"application/vnd.kde.karbon":                                  ".karbon",
		"application/vnd.kde.kchart":                                  ".chrt",
		"application/vnd.kde.kformula":                                ".kfo",
		"application/vnd.kde.kivio":                                   ".flw",
		"application/vnd.kde.kontour":                                 ".kon",
		"application/vnd.kde.kpresenter":                              ".kpr",
		"application/vnd.kde.kspread":                                 ".ksp",
		"application/vnd.kde.kword":                                   ".kwd",
		"application/vnd.kenameaapp":                                  ".htke",
		"application/vnd.kidspiration":                                ".kia",
		"application/vnd.kinar":                                       ".kne",
		"application/vnd.kodak-descriptor":                            ".sse",
		"application/vnd.las.las+xml":                                 ".lasxml",
		"application/x-latex":                                         ".latex",
		"application/vnd.llamagraphics.life-balance.desktop":          ".lbd",
		"application/vnd.llamagraphics.life-balance.exchange+xml":     ".lbe",
		"application/vnd.jam":                                         ".jam",
		"application/vnd.lotus-1-2-3":                                 ".123",
		"application/vnd.lotus-approach":                              ".apr",
		"application/vnd.lotus-freelance":                             ".pre",
		"application/vnd.lotus-notes":                                 ".nsf",
		"application/vnd.lotus-organizer":                             ".org",
		"application/vnd.lotus-screencam":                             ".scm",
		"application/vnd.lotus-wordpro":                               ".lwp",
		"audio/vnd.lucent.voice":                                      ".lvp",
		"audio/x-mpegurl":                                             ".m3u",
		"video/x-m4v":                                                 ".m4v",
		"application/mac-binhex40":                                    ".hqx",
		"application/vnd.macports.portpkg":                            ".portpkg",
		"application/vnd.osgeo.mapguide.package":                      ".mgp",
		"application/marc":                                            ".mrc",
		"application/marcxml+xml":                                     ".mrcx",
		"application/mxf":                                             ".mxf",
		"application/vnd.wolfram.player":                              ".nbp",
		"application/mathematica":                                     ".ma",
		"application/mathml+xml":                                      ".mathml",
		"application/mbox":                                            ".mbox",
		"application/vnd.medcalcdata":                                 ".mc1",
		"application/mediaservercontrol+xml":                          ".mscml",
		"application/vnd.mediastation.cdkey":                          ".cdkey",
		"application/vnd.mfer":                                        ".mwf",
		"application/vnd.mfmp":                                        ".mfm",
		"model/mesh":                                                  ".msh",
		"application/mads+xml":                                        ".mads",
		"application/mets+xml":                                        ".mets",
		"application/mods+xml":                                        ".mods",
		"application/metalink4+xml":                                   ".meta4",
		"application/vnd.mcd":                                         ".mcd",
		"application/vnd.micrografx.flo":                              ".flo",
		"application/vnd.micrografx.igx":                              ".igx",
		"application/vnd.eszigno3+xml":                                ".es3",
		"application/x-msaccess":                                      ".mdb",
		"video/x-ms-asf":                                              ".asf",
		"application/x-msdownload":                                    ".exe",
		"application/vnd.ms-artgalry":                                 ".cil",
		"application/vnd.ms-cab-compressed":                           ".cab",
		"application/vnd.ms-ims":                                      ".ims",
		"application/x-ms-application":                                ".application",
		"application/x-msclip":                                        ".clp",
		"image/vnd.ms-modi":                                           ".mdi",
		"application/vnd.ms-fontobject":                               ".eot",
		"application/vnd.ms-excel":                                    ".xls",
		"application/vnd.ms-excel.addin.macroenabled.12":              ".xlam",
		"application/vnd.ms-excel.sheet.binary.macroenabled.12":       ".xlsb",
		"application/vnd.ms-excel.template.macroenabled.12":           ".xltm",
		"application/vnd.ms-excel.sheet.macroenabled.12":              ".xlsm",
		"application/vnd.ms-htmlhelp":                                 ".chm",
		"application/x-mscardfile":                                    ".crd",
		"application/vnd.ms-lrm":                                      ".lrm",
		"application/x-msmediaview":                                   ".mvb",
		"application/x-msmoney":                                       ".mny",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
		"application/vnd.openxmlformats-officedocument.presentationml.slide":        ".sldx",
		"application/vnd.openxmlformats-officedocument.presentationml.slideshow":    ".ppsx",
		"application/vnd.openxmlformats-officedocument.presentationml.template":     ".potx",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.template":      ".xltx",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   ".docx",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.template":   ".dotx",
		"application/x-msbinder":                                     ".obd",
		"application/vnd.ms-officetheme":                             ".thmx",
		"application/onenote":                                        ".onetoc",
		"audio/vnd.ms-playready.media.pya":                           ".pya",
		"video/vnd.ms-playready.media.pyv":                           ".pyv",
		"application/vnd.ms-powerpoint":                              ".ppt",
		"application/vnd.ms-powerpoint.addin.macroenabled.12":        ".ppam",
		"application/vnd.ms-powerpoint.slide.macroenabled.12":        ".sldm",
		"application/vnd.ms-powerpoint.presentation.macroenabled.12": ".pptm",
		"application/vnd.ms-powerpoint.slideshow.macroenabled.12":    ".ppsm",
		"application/vnd.ms-powerpoint.template.macroenabled.12":     ".potm",
		"application/vnd.ms-project":                                 ".mpp",
		"application/x-mspublisher":                                  ".pub",
		"application/x-msschedule":                                   ".scd",
		"application/x-silverlight-app":                              ".xap",
		"application/vnd.ms-pki.stl":                                 ".stl",
		"application/vnd.ms-pki.seccat":                              ".cat",
		"application/vnd.visio":                                      ".vsd",
		"application/vnd.visio2013":                                  ".vsdx",
		"video/x-ms-wm":                                              ".wm",
		"audio/x-ms-wma":                                             ".wma",
		"audio/x-ms-wax":                                             ".wax",
		"video/x-ms-wmx":                                             ".wmx",
		"application/x-ms-wmd":                                       ".wmd",
		"application/vnd.ms-wpl":                                     ".wpl",
		"application/x-ms-wmz":                                       ".wmz",
		"video/x-ms-wmv":                                             ".wmv",
		"video/x-ms-wvx":                                             ".wvx",
		"application/x-msmetafile":                                   ".wmf",
		"application/x-msterminal":                                   ".trm",
		"application/msword":                                         ".doc",
		"application/vnd.ms-word.document.macroenabled.12":           ".docm",
		"application/vnd.ms-word.template.macroenabled.12":           ".dotm",
		"application/x-mswrite":                                      ".wri",
		"application/vnd.ms-works":                                   ".wps",
		"application/x-ms-xbap":                                      ".xbap",
		"application/vnd.ms-xpsdocument":                             ".xps",
		"audio/midi":                                                 ".mid",
		"application/vnd.ibm.minipay":                                ".mpy",
		"application/vnd.ibm.modcap":                                 ".afp",
		"application/vnd.jcp.javame.midlet-rms":                      ".rms",
		"application/vnd.tmobile-livetv":                             ".tmo",
		"application/x-mobipocket-ebook":                             ".prc",
		"application/vnd.mobius.mbk":                                 ".mbk",
		"application/vnd.mobius.dis":                                 ".dis",
		"application/vnd.mobius.plc":                                 ".plc",
		"application/vnd.mobius.mqy":                                 ".mqy",
		"application/vnd.mobius.msl":                                 ".msl",
		"application/vnd.mobius.txf":                                 ".txf",
		"application/vnd.mobius.daf":                                 ".daf",
		"text/vnd.fly":                                               ".fly",
		"application/vnd.mophun.certificate":                         ".mpc",
		"application/vnd.mophun.application":                         ".mpn",
		"video/mj2":                                                  ".mj2",
		"audio/mpeg":                                                 ".mpga",
		"video/vnd.mpegurl":                                          ".mxu",
		"video/mpeg":                                                 ".mpeg",
		"application/mp21":                                           ".m21",
		"audio/mp4":                                                  ".mp4a",
		"video/mp4":                                                  ".mp4",
		"application/mp4":                                            ".mp4",
		"application/vnd.apple.mpegurl":                              ".m3u8",
		"application/vnd.musician":                                   ".mus",
		"application/vnd.muvee.style":                                ".msty",
		"application/xv+xml":                                         ".mxml",
		"application/vnd.nokia.n-gage.data":                          ".ngdat",
		"application/vnd.nokia.n-gage.symbian.install":               ".n-gage",
		"application/x-dtbncx+xml":                                   ".ncx",
		"application/x-netcdf":                                       ".nc",
		"application/vnd.neurolanguage.nlu":                          ".nlu",
		"application/vnd.dna":                                        ".dna",
		"application/vnd.noblenet-directory":                         ".nnd",
		"application/vnd.noblenet-sealer":                            ".nns",
		"application/vnd.noblenet-web":                               ".nnw",
		"application/vnd.nokia.radio-preset":                         ".rpst",
		"application/vnd.nokia.radio-presets":                        ".rpss",
		"text/n3":                                                    ".n3",
		"application/vnd.novadigm.edm":                               ".edm",
		"application/vnd.novadigm.edx":                               ".edx",
		"application/vnd.novadigm.ext":                               ".ext",
		"application/vnd.flographit":                                 ".gph",
		"audio/vnd.nuera.ecelp4800":                                  ".ecelp4800",
		"audio/vnd.nuera.ecelp7470":                                  ".ecelp7470",
		"audio/vnd.nuera.ecelp9600":                                  ".ecelp9600",
		"application/oda":                                            ".oda",
		"application/ogg":                                            ".ogx",
		"audio/ogg":                                                  ".oga",
		"video/ogg":                                                  ".ogv",
		"application/vnd.oma.dd2+xml":                                ".dd2",
		"application/vnd.oasis.opendocument.text-web":                ".oth",
		"application/oebps-package+xml":                              ".opf",
		"application/vnd.intu.qbo":                                   ".qbo",
		"application/vnd.openofficeorg.extension":                    ".oxt",
		"application/vnd.yamaha.openscoreformat":                     ".osf",
		"audio/webm":                                                 ".weba",
		"video/webm":                                                 ".webm",
		"application/vnd.oasis.opendocument.chart":                   ".odc",
		"application/vnd.oasis.opendocument.chart-template":          ".otc",
		"application/vnd.oasis.opendocument.database":                ".odb",
		"application/vnd.oasis.opendocument.formula":                 ".odf",
		"application/vnd.oasis.opendocument.formula-template":        ".odft",
		"application/vnd.oasis.opendocument.graphics":                ".odg",
		"application/vnd.oasis.opendocument.graphics-template":       ".otg",
		"application/vnd.oasis.opendocument.image":                   ".odi",
		"application/vnd.oasis.opendocument.image-template":          ".oti",
		"application/vnd.oasis.opendocument.presentation":            ".odp",
		"application/vnd.oasis.opendocument.presentation-template":   ".otp",
		"application/vnd.oasis.opendocument.spreadsheet":             ".ods",
		"application/vnd.oasis.opendocument.spreadsheet-template":    ".ots",
		"application/vnd.oasis.opendocument.text":                    ".odt",
		"application/vnd.oasis.opendocument.text-master":             ".odm",
		"application/vnd.oasis.opendocument.text-template":           ".ott",
		"image/ktx":                                         ".ktx",
		"application/vnd.sun.xml.calc":                      ".sxc",
		"application/vnd.sun.xml.calc.template":             ".stc",
		"application/vnd.sun.xml.draw":                      ".sxd",
		"application/vnd.sun.xml.draw.template":             ".std",
		"application/vnd.sun.xml.impress":                   ".sxi",
		"application/vnd.sun.xml.impress.template":          ".sti",
		"application/vnd.sun.xml.math":                      ".sxm",
		"application/vnd.sun.xml.writer":                    ".sxw",
		"application/vnd.sun.xml.writer.global":             ".sxg",
		"application/vnd.sun.xml.writer.template":           ".stw",
		"application/x-font-otf":                            ".otf",
		"application/vnd.yamaha.openscoreformat.osfpvg+xml": ".osfpvg",
		"application/vnd.osgi.dp":                           ".dp",
		"application/vnd.palm":                              ".pdb",
		"text/x-pascal":                                     ".p",
		"application/vnd.pawaafile":                         ".paw",
		"application/vnd.hp-pclxl":                          ".pclxl",
		"application/vnd.picsel":                            ".efif",
		"image/x-pcx":                                       ".pcx",
		"image/vnd.adobe.photoshop":                         ".psd",
		"application/pics-rules":                            ".prf",
		"image/x-pict":                                      ".pic",
		"application/x-chat":                                ".chat",
		"application/pkcs10":                                ".p10",
		"application/x-pkcs12":                              ".p12",
		"application/pkcs7-mime":                            ".p7m",
		"application/pkcs7-signature":                       ".p7s",
		"application/x-pkcs7-certreqresp":                   ".p7r",
		"application/x-pkcs7-certificates":                  ".p7b",
		"application/pkcs8":                                 ".p8",
		"application/vnd.pocketlearn":                       ".plf",
		"image/x-portable-anymap":                           ".pnm",
		"image/x-portable-bitmap":                           ".pbm",
		"application/x-font-pcf":                            ".pcf",
		"application/font-tdpfr":                            ".pfr",
		"application/x-chess-pgn":                           ".pgn",
		"image/x-portable-graymap":                          ".pgm",
		"image/png":                                         ".png",
		"image/x-citrix-png":                                ".png",
		"image/x-png":                                       ".png",
		"image/x-portable-pixmap":                           ".ppm",
		"application/pskc+xml":                              ".pskcxml",
		"application/vnd.ctc-posml":                         ".pml",
		"application/postscript":                            ".ai",
		"application/x-font-type1":                          ".pfa",
		"application/vnd.powerbuilder6":                     ".pbd",
		"application/pgp-encrypted":                         ".pgp",
		"application/pgp-signature":                         ".pgp",
		"application/vnd.previewsystems.box":                ".box",
		"application/vnd.pvi.ptid1":                         ".ptid",
		"application/pls+xml":                               ".pls",
		"application/vnd.pg.format":                         ".str",
		"application/vnd.pg.osasli":                         ".ei6",
		"text/prs.lines.tag":                                ".dsc",
		"application/x-font-linux-psf":                      ".psf",
		"application/vnd.publishare-delta-tree":             ".qps",
		"application/vnd.pmi.widget":                        ".wg",
		"application/vnd.quark.quarkxpress":                 ".qxd",
		"application/vnd.epson.esf":                         ".esf",
		"application/vnd.epson.msf":                         ".msf",
		"application/vnd.epson.ssf":                         ".ssf",
		"application/vnd.epson.quickanime":                  ".qam",
		"application/vnd.intu.qfx":                          ".qfx",
		"video/quicktime":                                   ".qt",
		"application/x-rar-compressed":                      ".rar",
		"audio/x-pn-realaudio":                              ".ram",
		"audio/x-pn-realaudio-plugin":                       ".rmp",
		"application/rsd+xml":                               ".rsd",
		"application/vnd.rn-realmedia":                      ".rm",
		"application/vnd.realvnc.bed":                       ".bed",
		"application/vnd.recordare.musicxml":                ".mxl",
		"application/vnd.recordare.musicxml+xml":            ".musicxml",
		"application/relax-ng-compact-syntax":               ".rnc",
		"application/vnd.data-vision.rdz":                   ".rdz",
		"application/rdf+xml":                               ".rdf",
		"application/vnd.cloanto.rp9":                       ".rp9",
		"application/vnd.jisp":                              ".jisp",
		"application/rtf":                                   ".rtf",
		"text/richtext":                                     ".rtx",
		"application/vnd.route66.link66+xml":                ".link66",
		"application/rss+xml":                               ".rss, .xml",
		"application/shf+xml":                               ".shf",
		"application/vnd.sailingtracker.track":              ".st",
		"image/svg+xml":                                     ".svg",
		"application/vnd.sus-calendar":                      ".sus",
		"application/sru+xml":                               ".sru",
		"application/set-payment-initiation":                ".setpay",
		"application/set-registration-initiation":           ".setreg",
		"application/vnd.sema":                              ".sema",
		"application/vnd.semd":                              ".semd",
		"application/vnd.semf":                              ".semf",
		"application/vnd.seemail":                           ".see",
		"application/x-font-snf":                            ".snf",
		"application/scvp-vp-request":                       ".spq",
		"application/scvp-vp-response":                      ".spp",
		"application/scvp-cv-request":                       ".scq",
		"application/scvp-cv-response":                      ".scs",
		"application/sdp":                                   ".sdp",
		"text/x-setext":                                     ".etx",
		"video/x-sgi-movie":                                 ".movie",
		"application/vnd.shana.informed.formdata":           ".ifm",
		"application/vnd.shana.informed.formtemplate":       ".itp",
		"application/vnd.shana.informed.interchange":        ".iif",
		"application/vnd.shana.informed.package":            ".ipk",
		"application/thraud+xml":                            ".tfi",
		"application/x-shar":                                ".shar",
		"image/x-rgb":                                       ".rgb",
		"application/vnd.epson.salt":                        ".slt",
		"application/vnd.accpac.simply.aso":                 ".aso",
		"application/vnd.accpac.simply.imp":                 ".imp",
		"application/vnd.simtech-mindmapper":                ".twd",
		"application/vnd.commonspace":                       ".csp",
		"application/vnd.yamaha.smaf-audio":                 ".saf",
		"application/vnd.smaf":                              ".mmf",
		"application/vnd.yamaha.smaf-phrase":                ".spf",
		"application/vnd.smart.teacher":                     ".teacher",
		"application/vnd.svd":                               ".svd",
		"application/sparql-query":                          ".rq",
		"application/sparql-results+xml":                    ".srx",
		"application/srgs":                                  ".gram",
		"application/srgs+xml":                              ".grxml",
		"application/ssml+xml":                              ".ssml",
		"application/vnd.koan":                              ".skp",
		"text/sgml":                                         ".sgml",
		"application/vnd.stardivision.calc":                 ".sdc",
		"application/vnd.stardivision.draw":                 ".sda",
		"application/vnd.stardivision.impress":              ".sdd",
		"application/vnd.stardivision.math":                 ".smf",
		"application/vnd.stardivision.writer":               ".sdw",
		"application/vnd.stardivision.writer-global":        ".sgl",
		"application/vnd.stepmania.stepchart":               ".sm",
		"application/x-stuffit":                             ".sit",
		"application/x-stuffitx":                            ".sitx",
		"application/vnd.solent.sdkm+xml":                   ".sdkm",
		"application/vnd.olpc-sugar":                        ".xo",
		"audio/basic":                                       ".au",
		"application/vnd.wqd":                               ".wqd",
		"application/vnd.symbian.install":                   ".sis",
		"application/smil+xml":                              ".smi",
		"application/vnd.syncml+xml":                        ".xsm",
		"application/vnd.syncml.dm+wbxml":                   ".bdm",
		"application/vnd.syncml.dm+xml":                     ".xdm",
		"application/x-sv4cpio":                             ".sv4cpio",
		"application/x-sv4crc":                              ".sv4crc",
		"application/sbml+xml":                              ".sbml",
		"text/tab-separated-values":                         ".tsv",
		"image/tiff":                                        ".tiff",
		"application/vnd.tao.intent-module-archive":         ".tao",
		"application/x-tar":                                 ".tar",
		"application/x-tcl":                                 ".tcl",
		"application/x-tex":                                 ".tex",
		"application/x-tex-tfm":                             ".tfm",
		"application/tei+xml":                               ".tei",
		"text/plain":                                        ".txt",
		"application/vnd.spotfire.dxp":                      ".dxp",
		"application/vnd.spotfire.sfs":                      ".sfs",
		"application/timestamped-data":                      ".tsd",
		"application/vnd.trid.tpt":                          ".tpt",
		"application/vnd.triscape.mxs":                      ".mxs",
		"text/troff":                                        ".t",
		"application/vnd.trueapp":                           ".tra",
		"application/x-font-ttf":                            ".ttf",
		"text/turtle":                                       ".ttl",
		"application/vnd.umajin":                            ".umj",
		"application/vnd.uoml+xml":                          ".uoml",
		"application/vnd.unity":                             ".unityweb",
		"application/vnd.ufdl":                              ".ufd",
		"text/uri-list":                                     ".uri",
		"application/vnd.uiq.theme":                         ".utz",
		"application/x-ustar":                               ".ustar",
		"text/x-uuencode":                                   ".uu",
		"text/x-vcalendar":                                  ".vcs",
		"text/x-vcard":                                      ".vcf",
		"application/x-cdlink":                              ".vcd",
		"application/vnd.vsf":                               ".vsf",
		"model/vrml":                                        ".wrl",
		"application/vnd.vcx":                               ".vcx",
		"model/vnd.mts":                                     ".mts",
		"model/vnd.vtu":                                     ".vtu",
		"application/vnd.visionary":                         ".vis",
		"video/vnd.vivo":                                    ".viv",
		"application/ccxml+xml,":                            ".ccxml",
		"application/voicexml+xml":                          ".vxml",
		"application/x-wais-source":                         ".src",
		"application/vnd.wap.wbxml":                         ".wbxml",
		"image/vnd.wap.wbmp":                                ".wbmp",
		"audio/x-wav":                                       ".wav",
		"application/davmount+xml":                          ".davmount",
		"application/x-font-woff":                           ".woff",
		"application/wspolicy+xml":                          ".wspolicy",
		"image/webp":                                        ".webp",
		"application/vnd.webturbo":                          ".wtb",
		"application/widget":                                ".wgt",
		"application/winhlp":                                ".hlp",
		"text/vnd.wap.wml":                                  ".wml",
		"text/vnd.wap.wmlscript":                            ".wmls",
		"application/vnd.wap.wmlscriptc":                    ".wmlsc",
		"application/vnd.wordperfect":                       ".wpd",
		"application/vnd.wt.stf":                            ".stf",
		"application/wsdl+xml":                              ".wsdl",
		"image/x-xbitmap":                                   ".xbm",
		"image/x-xpixmap":                                   ".xpm",
		"image/x-xwindowdump":                               ".xwd",
		"application/x-x509-ca-cert":                        ".der",
		"application/x-xfig":                                ".fig",
		"application/xhtml+xml":                             ".xhtml",
		"application/xml":                                   ".xml",
		"application/xcap-diff+xml":                         ".xdf",
		"application/xenc+xml":                              ".xenc",
		"application/patch-ops-error+xml":                   ".xer",
		"application/resource-lists+xml":                    ".rl",
		"application/rls-services+xml":                      ".rs",
		"application/resource-lists-diff+xml":               ".rld",
		"application/xslt+xml":                              ".xslt",
		"application/xop+xml":                               ".xop",
		"application/x-xpinstall":                           ".xpi",
		"application/xspf+xml":                              ".xspf",
		"application/vnd.mozilla.xul+xml":                   ".xul",
		"chemical/x-xyz":                                    ".xyz",
		"text/yaml":                                         ".yaml",
		"application/yang":                                  ".yang",
		"application/yin+xml":                               ".yin",
		"application/vnd.zul":                               ".zir",
		"application/zip":                                   ".zip",
		"application/vnd.handheld-entertainment+xml":        ".zmm",
		"application/vnd.zzazz.deck+xml":                    ".zaz",
	}
}

type chunk struct {
	FileName  string
	StartByte int
	EndByte   int
	Index     int
	TmpFile   string
	Counter   writeCounter
}

type writeCounter struct {
	Total uint64
	Index int
	file  *DownloadFile
}

// DownloadFile takes file's url and has a download method
type DownloadFile struct {
	Link     string
	Filename string
	Size     int
	Chunks   []chunk
	Bar      *pb.ProgressBar
}
type browserDownload struct {
	URL string `json:"url"`
}

var wg sync.WaitGroup

// Download starts downloading the file it's called on and prints a progress bar to the terminal
func (file *DownloadFile) Download() {
	var filename, ext string
	request, err := http.NewRequest(http.MethodGet, file.Link, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	contentDisposition := response.Header.Get("Content-Disposition")

	if getMime(response.Header.Get("Content-Type")) == "" {
		filename, ext = file.handleFileName(contentDisposition)
	} else {
		filename, ext = file.handleFileName(contentDisposition)
		if ext == "" {
			ext = getMime(response.Header.Get("Content-Type"))
		}
	}
	file.Filename = filename
	file.Size, err = strconv.Atoi(response.Header.Get("Content-Length"))
	if err != nil {
		log.Fatal("Error converting size to integer", err)
	}
	start := time.Now()
	if file.existAndHealthy() {
		fmt.Println("File already exists in the same directory.")
	}
	if file.Size > 0 && !file.existAndHealthy() {
		file.Chunks = file.handleParts(4)
		if response.Header.Get("Accept-Ranges") != "bytes" {
			file.Chunks = file.handleParts(1)
		}
		runtime.GOMAXPROCS(len(file.Chunks))
		wg.Add(len(file.Chunks))

		for i := 0; i < len(file.Chunks); i++ {
			go func(chunkie chunk) {
				var tmpFile *os.File
				client := http.Client{}
				request, err := http.NewRequest(http.MethodGet, file.Link, nil)
				request.Header.Set("Range", "bytes="+strconv.Itoa(chunkie.StartByte)+"-"+strconv.Itoa(chunkie.EndByte))
				response, err := client.Do(request)
				if err != nil {
					log.Fatal(err)
				}
				tmpFile, err = ioutil.TempFile("./", "tmp-*.part")
				if err != nil {
					log.Fatal("Couldn't create temporary file", err)
				}
				editChunk(file.Chunks, chunkie, tmpFile.Name())
				defer response.Body.Close()
				_, err = io.Copy(tmpFile, io.TeeReader(response.Body, &chunkie.Counter))
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(file.Chunks[i])
		}
		file.Bar = getBar(file.Size)
		wg.Wait()
		downloaded, err := os.Create(file.Filename + ext)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(file.Chunks); i++ {
			tmp, err := ioutil.ReadFile(file.Chunks[i].TmpFile)
			if err != nil {
				log.Fatal(err)
			}
			if _, err = downloaded.Write(tmp); err != nil {
				log.Fatal(err)
			}
			if removalErr := os.Remove(file.Chunks[i].TmpFile); removalErr != nil {
				log.Fatal(removalErr)
			}
		}
	}
	err = os.Remove(filename + "-dl.gost")
	if err != nil {
		log.Fatal("Error deleting state file from disk :", err)
	}
	operationTime := time.Since(start)
	fmt.Printf("Downloaded %s in %s", humanize.Bytes(uint64(file.Size)), operationTime.String())
}

func (file *DownloadFile) handleParts(numOfChunks int) []chunk {
	var chunks []chunk
	chunkie := file.Size / numOfChunks
	remaining := file.Size % numOfChunks
	startByte := 0
	endByte := chunkie
	for i := 0; i < numOfChunks; i++ {
		chunks = append(chunks, chunk{file.Filename, startByte, endByte, i + 1, "", writeCounter{Index: i + 1, file: file}})
		startByte = endByte + 1
		endByte += chunkie + 1
	}
	chunks[numOfChunks-1].EndByte += remaining
	return chunks
}

func (file *DownloadFile) handleFileName(headerResult string) (string, string) {
	var fileName string
	fields := strings.Split(headerResult, ";")
	if len(fields) != 0 {
		if strings.Contains(fields[len(fields)-1], "filename") {
			fileName = strings.Replace(strings.Replace(strings.Replace(fields[len(fields)-1], "filename=", "", -1), "\"", "", -1), strings.Replace(strings.Replace(fields[len(fields)-1], "filename=", "", -1), "\"", "", -1)[strings.LastIndex(strings.Replace(strings.Replace(fields[len(fields)-1], "filename=", "", -1), "\"", "", -1), "."):], "", -1)
			return fileName, ""
		}
	}
	if fileName == "" {
		segments := strings.Split(file.Link, "/")
		if strings.Contains(segments[len(segments)-1], ".") {
			//removes what is after the last dot in the last segment of the url as the preceding is likely to be the file name
			fileName = strings.Replace(segments[len(segments)-1], segments[len(segments)-1][strings.LastIndex(segments[len(segments)-1], "."):], "", -1)
			extension := segments[len(segments)-1][strings.LastIndex(segments[len(segments)-1], ".") : strings.LastIndex(segments[len(segments)-1], ".")+4]
			return fileName, extension
		}
	}
	return "untitled", ""
}

func (file *DownloadFile) existAndHealthy() bool {
	var files []string

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal("Error list directory files :", err)
	}
	for _, aFile := range files {
		if strings.Contains(aFile, file.Filename) {
			info, err := os.Stat(aFile)
			if err != nil {
				log.Fatal("Error reading file info", err)
			}
			if int(info.Size()) == file.Size {
				return true
			}
		}
	}
	return false
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	stateFile, err := os.Create(wc.file.Filename + "-dl.gost")
	if err != nil {
		log.Fatal(err)
	}
	wc.file.Chunks[wc.Index-1].Counter.Total += uint64(n)
	wc.file.Bar.Add(n)
	defer stateFile.Close()
	enc := gob.NewEncoder(stateFile)
	err = enc.Encode(&wc.file.Chunks)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return n, nil
}

func getMime(key string) string {
	val, ok := mimesCons()[key]
	if ok {
		return val
	}
	return ""
}

func editChunk(chunks []chunk, chunkie chunk, fileName string) {
	chunks[chunkie.Index-1].TmpFile = fileName
}

func getBar(size int) *pb.ProgressBar {
	bar := pb.StartNew(size)
	bar.SetUnits(pb.U_BYTES)
	bar.ShowSpeed = true
	bar.ShowPercent = true
	return bar
}
