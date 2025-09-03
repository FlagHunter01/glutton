// [CUSTOM]
package custom

// Returns a string describing the port (empty if the port is unknown)
func PortDescription(port uint16) string {

	// Describes the port
	var description string

	// Assigns a value to the description for some intresting ports
	switch port {
	case 5:
		description = "Remote Job Entry"
	case 7:
		description = "Echo"
	case 11:
		description = "Active Users"
	case 13:
		description = "Daytime"
	case 17:
		description = "Quote of the Day"
	case 18:
		description = "Message Send Protocol (historic)"
	case 19:
		description = "Character Generator"
	case 20:
		description = "File Transfer [Default Data]"
	case 21:
		description = "File Transfer Protocol [Control]"
	case 22:
		description = "The Secure Shell (SSH) Protocol"
	case 23:
		description = "Telnet"
	case 24:
		description = "any private mail system"
	case 25:
		description = "Simple Mail Transfer"
	case 27:
		description = "NSW User System FE"
	case 29:
		description = "MSG ICP"
	case 31:
		description = "MSG Authentication"
	case 33:
		description = "Display Support Protocol"
	case 35:
		description = "any private printer server"
	case 37:
		description = "Time"
	case 38:
		description = "Route Access Protocol"
	case 39:
		description = "Resource Location Protocol"
	case 41:
		description = "Graphics"
	case 42:
		description = "Host Name Server"
	case 43:
		description = "Who Is"
	case 44:
		description = "MPM FLAGS Protocol"
	case 45:
		description = "Message Processing Module [recv]"
	case 46:
		description = "MPM [default send]"
	case 48:
		description = "Digital Audit Daemon"
	case 49:
		description = "Login Host Protocol (TACACS)"
	case 50:
		description = "Remote Mail Checking Protocol"
	case 52:
		description = "XNS Time Protocol"
	case 53:
		description = "Domain Name Server"
	case 54:
		description = "XNS Clearinghouse"
	case 55:
		description = "ISI Graphics Language"
	case 56:
		description = "XNS Authentication"
	case 57:
		description = "any private terminal access"
	case 58:
		description = "XNS Mail"
	case 59:
		description = "any private file service"
	case 62:
		description = "ACA Services"
	case 63:
		description = "whois++"
	case 64:
		description = "Communications Integrator (CI)"
	case 65:
		description = "TACACS-Database Service"
	case 66:
		description = "Oracle SQL*NET"
	case 67:
		description = "Bootstrap Protocol Server"
	case 68:
		description = "Bootstrap Protocol Client"
	case 69:
		description = "Trivial File Transfer"
	case 70:
		description = "Gopher"
	case 71:
		description = "Remote Job Service"
	case 72:
		description = "Remote Job Service"
	case 73:
		description = "Remote Job Service"
	case 74:
		description = "Remote Job Service"
	case 75:
		description = "any private dial out service"
	case 76:
		description = "Distributed External Object Store"
	case 77:
		description = "any private RJE service"
	case 78:
		description = "vettcp"
	case 79:
		description = "Finger"
	case 80:
		description = "HTTP"
	case 82:
		description = "XFER Utility"
	case 83:
		description = "MIT ML Device"
	case 84:
		description = "Common Trace Facility"
	case 85:
		description = "MIT ML Device"
	case 86:
		description = "Micro Focus Cobol"
	case 87:
		description = "any private terminal link"
	case 88:
		description = "Kerberos"
	case 89:
		description = "SU/MIT Telnet Gateway"
	case 90:
		description = "DNSIX Securit Attribute Token Map"
	case 91:
		description = "MIT Dover Spooler"
	case 92:
		description = "Network Printing Protocol"
	case 93:
		description = "Device Control Protocol"
	case 94:
		description = "Tivoli Object Dispatcher"
	case 95:
		description = "SUPDUP"
	case 96:
		description = "DIXIE Protocol Specification"
	case 97:
		description = "Swift Remote Virtural File Protocol"
	case 98:
		description = "TAC News"
	case 99:
		description = "Metagram Relay"
	case 101:
		description = "NIC Host Name Server"
	case 102:
		description = "ISO-TSAP Class 0"
	case 103:
		description = "Genesis Point-to-Point Trans Net"
	case 104:
		description = "ACR-NEMA Digital Imag. & Comm. 300"
	case 105:
		description = "Mailbox Name Nameserver"
	case 106:
		description = "3COM-TSMUX"
	case 107:
		description = "Remote Telnet Service"
	case 108:
		description = "SNA Gateway Access Server"
	case 109:
		description = "Post Office Protocol - Version 2"
	case 110:
		description = "Post Office Protocol - Version 3"
	case 111:
		description = "SUN Remote Procedure Call"
	case 112:
		description = "McIDAS Data Transmission Protocol"
	case 113:
		description = "Authentication Service"
	case 115:
		description = "Simple File Transfer Protocol"
	case 116:
		description = "ANSA REX Notify"
	case 117:
		description = "UUCP Path Service"
	case 118:
		description = "SQL Services"
	case 119:
		description = "Network News Transfer Protocol"
	case 120:
		description = "CFDPTKT"
	case 121:
		description = "Encore Expedited Remote Pro.Call"
	case 122:
		description = "SMAKYNET"
	case 123:
		description = "Network Time Protocol"
	case 124:
		description = "ANSA REX Trader"
	case 125:
		description = "Locus PC-Interface Net Map Ser"
	case 126:
		description = "NXEdit"
	case 127:
		description = "Locus PC-Interface Conn Server"
	case 128:
		description = "GSS X License Verification"
	case 129:
		description = "Password Generator Protocol"
	case 130:
		description = "cisco FNATIVE"
	case 131:
		description = "cisco TNATIVE"
	case 132:
		description = "cisco SYSMAINT"
	case 133:
		description = "Statistics Service"
	case 134:
		description = "INGRES-NET Service"
	case 135:
		description = "DCE endpoint resolution"
	case 136:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 137:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 138:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 139:
		description = "NETBIOS Session Service"
	case 140:
		description = "EMFIS Data Service"
	case 141:
		description = "EMFIS Control Service"
	case 142:
		description = "Britton-Lee IDM"
	case 143:
		description = "Internet Message Access Protocol"
	case 144:
		description = "Universal Management Architecture"
	case 145:
		description = "UAAC Protocol"
	case 146:
		description = "ISO-IP0"
	case 147:
		description = "ISO-IP"
	case 148:
		description = "Jargon"
	case 149:
		description = "AED 512 Emulation Service"
	case 150:
		description = "SQL-NET"
	case 151:
		description = "HEMS"
	case 152:
		description = "Background File Transfer Program"
	case 153:
		description = "SGMP"
	case 154:
		description = "NETSC"
	case 155:
		description = "NETSC"
	case 156:
		description = "SQL Service"
	case 157:
		description = "KNET/VM Command/Message Protocol"
	case 158:
		description = "PCMail Server"
	case 159:
		description = "NSS-Routing"
	case 160:
		description = "SGMP-TRAPS"
	case 161:
		description = "SNMP"
	case 162:
		description = "SNMPTRAP"
	case 163:
		description = "CMIP/TCP Manager"
	case 164:
		description = "CMIP/TCP Agent"
	case 165:
		description = "Xerox"
	case 166:
		description = "Sirius Systems"
	case 167:
		description = "NAMP"
	case 168:
		description = "RSVD"
	case 169:
		description = "SEND"
	case 170:
		description = "Network PostScript"
	case 171:
		description = "Network Innovations Multiplex"
	case 172:
		description = "Network Innovations CL/1"
	case 173:
		description = "Xyplex"
	case 174:
		description = "MAILQ"
	case 175:
		description = "VMNET"
	case 176:
		description = "GENRAD-MUX"
	case 177:
		description = "X Display Manager Control Protocol"
	case 178:
		description = "NextStep Window Server"
	case 179:
		description = "Border Gateway Protocol"
	case 180:
		description = "Intergraph"
	case 181:
		description = "Unify"
	case 182:
		description = "Unisys Audit SITP"
	case 183:
		description = "OCBinder"
	case 184:
		description = "OCServer"
	case 185:
		description = "Remote-KIS"
	case 186:
		description = "KIS Protocol"
	case 187:
		description = "Application Communication Interface"
	case 188:
		description = "Plus Five's MUMPS"
	case 189:
		description = "Queued File Transport"
	case 190:
		description = "Gateway Access Control Protocol"
	case 191:
		description = "Prospero Directory Service"
	case 192:
		description = "OSU Network Monitoring System"
	case 193:
		description = "Spider Remote Monitoring Protocol"
	case 194:
		description = "Internet Relay Chat Protocol"
	case 195:
		description = "DNSIX Network Level Module Audit"
	case 196:
		description = "DNSIX Session Mgt Module Audit Redir"
	case 197:
		description = "Directory Location Service"
	case 198:
		description = "Directory Location Service Monitor"
	case 199:
		description = "SMUX"
	case 200:
		description = "IBM System Resource Controller"
	case 201:
		description = "AppleTalk Routing Maintenance"
	case 202:
		description = "AppleTalk Name Binding"
	case 203:
		description = "AppleTalk Unused"
	case 204:
		description = "AppleTalk Echo"
	case 205:
		description = "AppleTalk Unused"
	case 206:
		description = "AppleTalk Zone Information"
	case 207:
		description = "AppleTalk Unused"
	case 208:
		description = "AppleTalk Unused"
	case 209:
		description = "The Quick Mail Transfer Protocol"
	case 210:
		description = "ANSI Z39.50"
	case 211:
		description = "Texas Instruments 914C/G Terminal"
	case 212:
		description = "ATEXSSTR"
	case 213:
		description = "IPX"
	case 214:
		description = "VM PWSCS"
	case 215:
		description = "Insignia Solutions"
	case 216:
		description = "Computer Associates Int'l License Server"
	case 217:
		description = "dBASE Unix"
	case 218:
		description = "Netix Message Posting Protocol"
	case 219:
		description = "Unisys ARPs"
	case 220:
		description = "Interactive Mail Access Protocol v3"
	case 221:
		description = "Berkeley rlogind with SPX auth"
	case 222:
		description = "Berkeley rshd with SPX auth"
	case 223:
		description = "Certificate Distribution Center"
	case 224:
		description = "masqdialer"
	case 242:
		description = "Direct"
	case 243:
		description = "Survey Measurement"
	case 244:
		description = "inbusiness"
	case 245:
		description = "LINK"
	case 246:
		description = "Display Systems Protocol"
	case 247:
		description = "SUBNTBCST_TFTP"
	case 248:
		description = "bhfhs"
	case 256:
		description = "RAP"
	case 257:
		description = "Secure Electronic Transaction"
	case 259:
		description = "Efficient Short Remote Operations"
	case 260:
		description = "Openport"
	case 261:
		description = "IIOP Name Service over TLS/SSL"
	case 262:
		description = "Arcisdms"
	case 263:
		description = "HDAP"
	case 264:
		description = "BGMP"
	case 265:
		description = "X-Bone CTL"
	case 266:
		description = "SCSI on ST"
	case 267:
		description = "Tobit David Service Layer"
	case 268:
		description = "Tobit David Replica"
	case 269:
		description = "MANET Protocols"
	case 270:
		description = "Q-mode encapsulation for GIST messages"
	case 271:
		description = "IETF PT-TLS"
	case 280:
		description = "http-mgmt"
	case 281:
		description = "Personal Link"
	case 282:
		description = "Cable Port A/X"
	case 283:
		description = "rescap"
	case 284:
		description = "corerjd"
	case 286:
		description = "FXP Communication"
	case 287:
		description = "K-BLOCK"
	case 300:
		description = "TLS Secure Login Host Protocol (TACACSS)"
	case 308:
		description = "Novastor Backup"
	case 309:
		description = "EntrustTime"
	case 310:
		description = "bhmds"
	case 311:
		description = "AppleShare IP WebAdmin"
	case 312:
		description = "VSLMP"
	case 313:
		description = "Magenta Logic"
	case 314:
		description = "Opalis Robot"
	case 315:
		description = "DPSI"
	case 316:
		description = "decAuth"
	case 317:
		description = "Zannet"
	case 318:
		description = "PKIX TimeStamp"
	case 319:
		description = "PTP Event"
	case 320:
		description = "PTP General"
	case 321:
		description = "PIP"
	case 322:
		description = "RTSPS"
	case 323:
		description = "Resource PKI to Router Protocol"
	case 324:
		description = "Resource PKI to Router Protocol over TLS"
	case 333:
		description = "Texar Security Port"
	case 344:
		description = "Prospero Data Access Protocol"
	case 345:
		description = "Perf Analysis Workbench"
	case 346:
		description = "Zebra server"
	case 347:
		description = "Fatmen Server"
	case 348:
		description = "Cabletron Management Protocol"
	case 349:
		description = "mftp"
	case 350:
		description = "MATIP Type A"
	case 351:
		description = "MATIP Type B"
	case 352:
		description = "bhoedap4"
	case 353:
		description = "NDSAUTH"
	case 354:
		description = "bh611"
	case 355:
		description = "DATEX-ASN"
	case 356:
		description = "Cloanto Net 1"
	case 357:
		description = "bhevent"
	case 358:
		description = "Shrinkwrap"
	case 360:
		description = "scoi2odialog"
	case 361:
		description = "Semantix"
	case 362:
		description = "SRS Send"
	case 363:
		description = "RSVP Tunnel"
	case 364:
		description = "Aurora CMGR"
	case 365:
		description = "DTK"
	case 366:
		description = "ODMR"
	case 367:
		description = "MortgageWare"
	case 368:
		description = "QbikGDP"
	case 369:
		description = "rpc2portmap"
	case 370:
		description = "codaauth2"
	case 371:
		description = "Clearcase"
	case 372:
		description = "ListProcessor"
	case 373:
		description = "Legent Corporation"
	case 374:
		description = "Legent Corporation"
	case 375:
		description = "Hassle"
	case 376:
		description = "Amiga Envoy Network Inquiry Protocol"
	case 377:
		description = "NEC Corporation"
	case 378:
		description = "NEC Corporation"
	case 379:
		description = "TIA/EIA/IS-99 modem client"
	case 380:
		description = "TIA/EIA/IS-99 modem server"
	case 381:
		description = "hp performance data collector"
	case 382:
		description = "hp performance data managed node"
	case 383:
		description = "hp performance data alarm manager"
	case 384:
		description = "A Remote Network Server System"
	case 385:
		description = "IBM Application"
	case 386:
		description = "ASA Message Router Object Def."
	case 387:
		description = "Appletalk Update-Based Routing Pro."
	case 388:
		description = "Unidata LDM"
	case 389:
		description = "Lightweight Directory Access Protocol"
	case 390:
		description = "UIS"
	case 391:
		description = "SynOptics SNMP Relay Port"
	case 392:
		description = "SynOptics Port Broker Port"
	case 393:
		description = "Meta5"
	case 394:
		description = "EMBL Nucleic Data Transfer"
	case 395:
		description = "NetScout Control Protocol"
	case 396:
		description = "Novell Netware over IP"
	case 397:
		description = "Multi Protocol Trans. Net."
	case 398:
		description = "Kryptolan"
	case 399:
		description = "ISO Transport Class 2 Non-Control"
	case 400:
		description = "Oracle Secure Backup"
	case 401:
		description = "Uninterruptible Power Supply"
	case 402:
		description = "Genie Protocol"
	case 403:
		description = "decap"
	case 404:
		description = "nced"
	case 405:
		description = "ncld"
	case 406:
		description = "Interactive Mail Support Protocol"
	case 407:
		description = "Timbuktu"
	case 408:
		description = "Prospero Resource Manager Sys. Man."
	case 409:
		description = "Prospero Resource Manager Node Man."
	case 410:
		description = "DECLadebug Remote Debug Protocol"
	case 411:
		description = "Remote MT Protocol"
	case 412:
		description = "Trap Convention Port"
	case 413:
		description = "Storage Management Services Protocol"
	case 414:
		description = "InfoSeek"
	case 415:
		description = "BNet"
	case 416:
		description = "Silverplatter"
	case 417:
		description = "Onmux"
	case 418:
		description = "Hyper-G"
	case 419:
		description = "Ariel 1"
	case 420:
		description = "SMPTE"
	case 421:
		description = "Ariel 2"
	case 422:
		description = "Ariel 3"
	case 423:
		description = "IBM Operations Planning and Control Start"
	case 424:
		description = "IBM Operations Planning and Control Track"
	case 425:
		description = "ICAD"
	case 426:
		description = "smartsdp"
	case 427:
		description = "Server Location"
	case 428:
		description = "OCS_CMU"
	case 429:
		description = "OCS_AMU"
	case 430:
		description = "UTMPSD"
	case 431:
		description = "UTMPCD"
	case 432:
		description = "IASD"
	case 433:
		description = "NNTP for transit servers (NNSP)"
	case 434:
		description = "MobileIP-Agent"
	case 435:
		description = "MobilIP-MN"
	case 436:
		description = "DNA-CML"
	case 437:
		description = "comscm"
	case 438:
		description = "dsfgw"
	case 439:
		description = "dasp"
	case 440:
		description = "sgcp"
	case 441:
		description = "decvms-sysmgt"
	case 442:
		description = "cvc_hostd"
	case 443:
		description = "HTTPS"
	case 444:
		description = "Simple Network Paging Protocol"
	case 445:
		description = "Microsoft-DS"
	case 446:
		description = "DDM-Remote Relational Database Access"
	case 447:
		description = "DDM-Distributed File Management"
	case 448:
		description = "DDM-Remote DB Access Using Secure Sockets"
	case 449:
		description = "AS Server Mapper"
	case 450:
		description = "Computer Supported Telecomunication Applications"
	case 451:
		description = "Cray Network Semaphore server"
	case 452:
		description = "Cray SFS config server"
	case 453:
		description = "CreativeServer"
	case 454:
		description = "ContentServer"
	case 455:
		description = "CreativePartnr"
	case 456:
		description = "macon"
	case 457:
		description = "scohelp"
	case 458:
		description = "apple quick time"
	case 459:
		description = "ampr-rcmd"
	case 460:
		description = "skronk"
	case 461:
		description = "DataRampSrv"
	case 462:
		description = "DataRampSrvSec"
	case 463:
		description = "alpes"
	case 464:
		description = "kpasswd"
	case 465:
		description = "Message Submission over TLS protocol"
	case 466:
		description = "digital-vrc"
	case 467:
		description = "mylex-mapd"
	case 468:
		description = "proturis"
	case 469:
		description = "Radio Control Protocol"
	case 470:
		description = "scx-proxy"
	case 471:
		description = "Mondex"
	case 472:
		description = "ljk-login"
	case 473:
		description = "hybrid-pop"
	case 474:
		description = "tn-tl-w1"
	case 475:
		description = "tcpnethaspsrv"
	case 476:
		description = "tn-tl-fd1"
	case 477:
		description = "ss7ns"
	case 478:
		description = "spsc"
	case 479:
		description = "iafserver"
	case 480:
		description = "iafdbase"
	case 481:
		description = "Ph service"
	case 482:
		description = "bgs-nsi"
	case 483:
		description = "ulpnet"
	case 484:
		description = "Integra Software Management Environment"
	case 485:
		description = "Air Soft Power Burst"
	case 486:
		description = "avian"
	case 487:
		description = "saft Simple Asynchronous File Transfer"
	case 488:
		description = "gss-http"
	case 489:
		description = "nest-protocol"
	case 490:
		description = "micom-pfs"
	case 491:
		description = "go-login"
	case 492:
		description = "Transport Independent Convergence for FNA"
	case 493:
		description = "Transport Independent Convergence for FNA"
	case 494:
		description = "POV-Ray"
	case 495:
		description = "intecourier"
	case 496:
		description = "PIM-RP-DISC"
	case 497:
		description = "Retrospect backup and restore service"
	case 498:
		description = "siam"
	case 499:
		description = "ISO ILL Protocol"
	case 500:
		description = "isakmp"
	case 501:
		description = "STMF"
	case 502:
		description = "Modbus Application Protocol"
	case 503:
		description = "Intrinsa"
	case 504:
		description = "citadel"
	case 505:
		description = "mailbox-lm"
	case 506:
		description = "ohimsrv"
	case 507:
		description = "crs"
	case 508:
		description = "xvttp"
	case 509:
		description = "snare"
	case 510:
		description = "FirstClass Protocol"
	case 511:
		description = "PassGo"
	case 512:
		description = "remote process execution"
	case 513:
		description = "remote login a la telnet"
	case 514:
		description = "cmd like exec"
	case 515:
		description = "spooler"
	case 516:
		description = "videotex"
	case 517:
		description = "like tenex link"
	case 519:
		description = "unixtime"
	case 520:
		description = "extended file name server"
	case 521:
		description = "ripng"
	case 522:
		description = "ULP"
	case 523:
		description = "IBM-DB2"
	case 524:
		description = "NCP"
	case 525:
		description = "timeserver"
	case 526:
		description = "newdate"
	case 527:
		description = "Stock IXChange"
	case 528:
		description = "Customer IXChange"
	case 529:
		description = "IRC-SERV"
	case 530:
		description = "rpc"
	case 531:
		description = "chat"
	case 532:
		description = "readnews"
	case 533:
		description = "for emergency broadcasts"
	case 534:
		description = "windream Admin"
	case 535:
		description = "iiop"
	case 536:
		description = "opalis-rdv"
	case 537:
		description = "Networked Media Streaming Protocol"
	case 538:
		description = "gdomap"
	case 539:
		description = "Apertus Technologies Load Determination"
	case 540:
		description = "uucpd"
	case 541:
		description = "uucp-rlogin"
	case 542:
		description = "commerce"
	case 544:
		description = "krcmd"
	case 545:
		description = "appleqtcsrvr"
	case 546:
		description = "DHCPv6 Client"
	case 547:
		description = "DHCPv6 Server"
	case 548:
		description = "AFP over TCP"
	case 549:
		description = "IDFP"
	case 550:
		description = "new-who"
	case 551:
		description = "cybercash"
	case 552:
		description = "DeviceShare"
	case 553:
		description = "pirp"
	case 554:
		description = "Real Time Streaming Protocol (RTSP)"
	case 556:
		description = "rfs server"
	case 557:
		description = "openvms-sysipc"
	case 558:
		description = "SDNSKMP"
	case 559:
		description = "TEEDTAP"
	case 560:
		description = "rmonitord"
	case 562:
		description = "chcmd"
	case 563:
		description = "nntp protocol over TLS/SSL (was snntp)"
	case 564:
		description = "plan 9 file service"
	case 565:
		description = "whoami"
	case 566:
		description = "streettalk"
	case 567:
		description = "banyan-rpc"
	case 568:
		description = "microsoft shuttle"
	case 569:
		description = "microsoft rome"
	case 570:
		description = "demon"
	case 571:
		description = "udemon"
	case 572:
		description = "sonar"
	case 573:
		description = "banyan-vip"
	case 574:
		description = "FTP Software Agent System"
	case 575:
		description = "VEMMI"
	case 576:
		description = "ipcd"
	case 577:
		description = "vnas"
	case 578:
		description = "ipdd"
	case 579:
		description = "decbsrv"
	case 580:
		description = "SNTP HEARTBEAT"
	case 581:
		description = "Bundle Discovery Protocol"
	case 582:
		description = "SCC Security"
	case 583:
		description = "Philips Video-Conferencing"
	case 584:
		description = "Key Server"
	case 585:
		description = "De-registered"
	case 586:
		description = "Password Change"
	case 587:
		description = "Message Submission"
	case 588:
		description = "CAL"
	case 589:
		description = "EyeLink"
	case 590:
		description = "TNS CML"
	case 591:
		description = "FileMaker"
	case 592:
		description = "Eudora Set"
	case 593:
		description = "HTTP RPC Ep Map"
	case 594:
		description = "TPIP"
	case 595:
		description = "CAB Protocol"
	case 596:
		description = "SMSD"
	case 597:
		description = "PTC Name Service"
	case 598:
		description = "SCO Web Server Manager 3"
	case 599:
		description = "Aeolon Core Protocol"
	case 600:
		description = "Sun IPC server"
	case 601:
		description = "Reliable Syslog Service"
	case 602:
		description = "XML-RPC over BEEP"
	case 603:
		description = "IDXP"
	case 604:
		description = "TUNNEL"
	case 605:
		description = "SOAP over BEEP"
	case 606:
		description = "Cray Unified Resource Manager"
	case 607:
		description = "nqs"
	case 608:
		description = "Sender-Initiated/Unsolicited File Transfer"
	case 609:
		description = "npmp-trap"
	case 610:
		description = "npmp-local"
	case 611:
		description = "npmp-gui"
	case 612:
		description = "HMMP Indication"
	case 613:
		description = "HMMP Operation"
	case 614:
		description = "SSLshell"
	case 615:
		description = "Internet Configuration Manager"
	case 616:
		description = "SCO System Administration Server"
	case 617:
		description = "SCO Desktop Administration Server"
	case 618:
		description = "DEI-ICDA"
	case 619:
		description = "Compaq EVM"
	case 620:
		description = "SCO WebServer Manager"
	case 621:
		description = "ESCP"
	case 622:
		description = "Collaborator"
	case 623:
		description = "ASF Remote Management and Control Protocol"
	case 624:
		description = "Crypto Admin"
	case 625:
		description = "DEC DLM"
	case 626:
		description = "ASIA"
	case 627:
		description = "PassGo Tivoli"
	case 628:
		description = "QMQP"
	case 629:
		description = "3Com AMP3"
	case 630:
		description = "RDA"
	case 631:
		description = "Internet Printing Protocol over HTTPS"
	case 632:
		description = "bmpp"
	case 633:
		description = "Service Status update (Sterling Software)"
	case 634:
		description = "ginad"
	case 635:
		description = "RLZ DBase"
	case 636:
		description = "ldap protocol over TLS/SSL (was sldap)"
	case 637:
		description = "lanserver"
	case 638:
		description = "mcns-sec"
	case 639:
		description = "MSDP"
	case 640:
		description = "entrust-sps"
	case 641:
		description = "repcmd"
	case 642:
		description = "ESRO-EMSDP V1.3"
	case 643:
		description = "SANity"
	case 644:
		description = "dwr"
	case 645:
		description = "PSSC"
	case 646:
		description = "LDP"
	case 647:
		description = "DHCP Failover"
	case 648:
		description = "Registry Registrar Protocol (RRP)"
	case 649:
		description = "Cadview-3d"
	case 650:
		description = "OBEX"
	case 651:
		description = "IEEE MMS"
	case 652:
		description = "HELLO_PORT"
	case 653:
		description = "RepCmd"
	case 654:
		description = "AODV"
	case 655:
		description = "TINC"
	case 656:
		description = "SPMP"
	case 657:
		description = "RMC"
	case 658:
		description = "TenFold"
	case 659:
		description = "Removed"
	case 660:
		description = "MacOS Server Admin"
	case 661:
		description = "HAP"
	case 662:
		description = "PFTP"
	case 663:
		description = "PureNoise"
	case 664:
		description = "ASF Secure Remote Management and Control Protocol"
	case 665:
		description = "Sun DR"
	case 666:
		description = "DOOM"
	case 667:
		description = "campaign contribution disclosures - SDR Technologies"
	case 668:
		description = "MeComm"
	case 669:
		description = "MeRegister"
	case 670:
		description = "VACDSM-SWS"
	case 671:
		description = "VACDSM-APP"
	case 672:
		description = "VPPS-QUA"
	case 673:
		description = "CIMPLEX"
	case 674:
		description = "ACAP"
	case 675:
		description = "DCTP"
	case 676:
		description = "VPPS Via"
	case 677:
		description = "Virtual Presence Protocol"
	case 678:
		description = "GNU Generation Foundation NCP"
	case 679:
		description = "MRM"
	case 680:
		description = "entrust-aaas"
	case 681:
		description = "entrust-aams"
	case 682:
		description = "XFR"
	case 683:
		description = "CORBA IIOP"
	case 684:
		description = "CORBA IIOP SSL"
	case 685:
		description = "MDC Port Mapper"
	case 686:
		description = "Hardware Control Protocol Wismar"
	case 687:
		description = "asipregistry"
	case 688:
		description = "ApplianceWare managment protocol"
	case 689:
		description = "NMAP"
	case 690:
		description = "Velneo Application Transfer Protocol"
	case 691:
		description = "MS Exchange Routing"
	case 692:
		description = "Hyperwave-ISP"
	case 693:
		description = "almanid Connection Endpoint"
	case 694:
		description = "ha-cluster"
	case 695:
		description = "IEEE-MMS-SSL"
	case 696:
		description = "RUSHD"
	case 697:
		description = "UUIDGEN"
	case 698:
		description = "OLSR"
	case 699:
		description = "Access Network"
	case 700:
		description = "Extensible Provisioning Protocol"
	case 701:
		description = "Link Management Protocol (LMP)"
	case 702:
		description = "IRIS over BEEP"
	case 704:
		description = "errlog copy/server daemon"
	case 705:
		description = "AgentX"
	case 706:
		description = "SILC"
	case 707:
		description = "Borland DSJ"
	case 709:
		description = "Entrust Key Management Service Handler"
	case 710:
		description = "Entrust Administration Service Handler"
	case 711:
		description = "Cisco TDP"
	case 712:
		description = "TBRPF"
	case 713:
		description = "IRIS over XPC"
	case 714:
		description = "IRIS over XPCS"
	case 715:
		description = "IRIS-LWZ"
	case 716:
		description = "PANA Messages"
	case 729:
		description = "IBM NetView DM/6000 Server/Client"
	case 730:
		description = "IBM NetView DM/6000 send/tcp"
	case 731:
		description = "IBM NetView DM/6000 receive/tcp"
	case 741:
		description = "netGW"
	case 742:
		description = "Network based Rev. Cont. Sys."
	case 744:
		description = "Flexible License Manager"
	case 747:
		description = "Fujitsu Device Control"
	case 748:
		description = "Russell Info Sci Calendar Manager"
	case 749:
		description = "kerberos administration"
	case 750:
		description = "kerberos version iv"
	case 754:
		description = "send"
	case 767:
		description = "phone"
	case 777:
		description = "Multiling HTTP"
	case 802:
		description = "Modbus Application Protocol Secure"
	case 810:
		description = "FCP"
	case 828:
		description = "itm-mcell-s"
	case 829:
		description = "PKIX-3 CA/RA"
	case 830:
		description = "NETCONF over SSH"
	case 831:
		description = "NETCONF over BEEP"
	case 832:
		description = "NETCONF for SOAP over HTTPS"
	case 833:
		description = "NETCONF for SOAP over BEEP"
	case 847:
		description = "dhcp-failover 2"
	case 848:
		description = "GDOI"
	case 853:
		description = "DNS query-response protocol run over TLS"
	case 854:
		description = "Dynamic Link Exchange Protocol (DLEP)"
	case 860:
		description = "iSCSI"
	case 861:
		description = "OWAMP-Control"
	case 862:
		description = "TWAMP-Control"
	case 873:
		description = "rsync"
	case 886:
		description = "ICL coNETion locate server"
	case 887:
		description = "ICL coNETion server info"
	case 888:
		description = "CD Database Protocol"
	case 900:
		description = "OMG Initial Refs"
	case 901:
		description = "SMPNAMERES"
	case 902:
		description = "self documenting Telnet Door"
	case 903:
		description = "self documenting Telnet Panic Door"
	case 910:
		description = "Kerberized Internet Negotiation of Keys (KINK)"
	case 911:
		description = "xact-backup"
	case 912:
		description = "APEX relay-relay service"
	case 913:
		description = "APEX endpoint-relay service"
	case 914:
		description = "Routing in Fat Trees Link Information Element"
	case 915:
		description = "Routing in Fat Trees Topology Information Element"
	case 953:
		description = "BIND9 remote name daemon controller"
	case 989:
		description = "ftp protocol"
	case 990:
		description = "ftp protocol"
	case 991:
		description = "Netnews Administration System"
	case 992:
		description = "telnet protocol over TLS/SSL"
	case 993:
		description = "IMAP over TLS protocol"
	case 995:
		description = "POP3 over TLS protocol"
	case 996:
		description = "vsinet"
	case 999:
		description = "Applix ac"
	case 1001:
		description = "HTTP Web Push"
	case 1008:
		description = "Possibly used by Sun Solaris????"
	case 1010:
		description = "surf"
	case 1021:
		description = "RFC3692-style Experiment 1"
	case 1022:
		description = "RFC3692-style Experiment 2"
	case 1025:
		description = "network blackjack"
	case 1026:
		description = "Calendar Access Protocol"
	case 1027:
		description = "IPv6 Behind NAT44 CPEs"
	case 1028:
		description = "Deprecated"
	case 1029:
		description = "Solid Mux Server"
	case 1033:
		description = "local netinfo port"
	case 1034:
		description = "ActiveSync Notifications"
	case 1035:
		description = "MX-XR RPC"
	case 1036:
		description = "Nebula Secure Segment Transfer Protocol"
	case 1037:
		description = "AMS"
	case 1038:
		description = "Message Tracking Query Protocol"
	case 1039:
		description = "Streamlined Blackhole"
	case 1040:
		description = "Netarx Netcare"
	case 1041:
		description = "AK2 Product"
	case 1042:
		description = "Subnet Roaming"
	case 1043:
		description = "BOINC Client Control"
	case 1044:
		description = "Dev Consortium Utility"
	case 1045:
		description = "Fingerprint Image Transfer Protocol"
	case 1046:
		description = "WebFilter Remote Monitor"
	case 1047:
		description = "Sun's NEO Object Request Broker"
	case 1048:
		description = "Sun's NEO Object Request Broker"
	case 1049:
		description = "Tobit David Postman VPMN"
	case 1050:
		description = "CORBA Management Agent"
	case 1051:
		description = "Optima VNET"
	case 1052:
		description = "Dynamic DNS Tools"
	case 1053:
		description = "Remote Assistant (RA)"
	case 1054:
		description = "BRVREAD"
	case 1055:
		description = "ANSYS - License Manager"
	case 1056:
		description = "VFO"
	case 1057:
		description = "STARTRON"
	case 1058:
		description = "nim"
	case 1059:
		description = "nimreg"
	case 1060:
		description = "POLESTAR"
	case 1061:
		description = "KIOSK"
	case 1062:
		description = "Veracity"
	case 1063:
		description = "KyoceraNetDev"
	case 1064:
		description = "JSTEL"
	case 1065:
		description = "SYSCOMLAN"
	case 1066:
		description = "FPO-FNS"
	case 1067:
		description = "Installation Bootstrap Proto. Serv."
	case 1068:
		description = "Installation Bootstrap Proto. Cli."
	case 1069:
		description = "COGNEX-INSIGHT"
	case 1070:
		description = "GMRUpdateSERV"
	case 1071:
		description = "BSQUARE-VOIP"
	case 1072:
		description = "CARDAX"
	case 1073:
		description = "Bridge Control"
	case 1074:
		description = "Warmspot Management Protocol"
	case 1075:
		description = "RDRMSHC"
	case 1076:
		description = "DAB STI-C"
	case 1077:
		description = "IMGames"
	case 1078:
		description = "Avocent Proxy Protocol"
	case 1079:
		description = "ASPROVATalk"
	case 1080:
		description = "Socks"
	case 1081:
		description = "PVUNIWIEN"
	case 1082:
		description = "AMT-ESD-PROT"
	case 1083:
		description = "Anasoft License Manager"
	case 1084:
		description = "Anasoft License Manager"
	case 1085:
		description = "Web Objects"
	case 1086:
		description = "CPL Scrambler Logging"
	case 1087:
		description = "CPL Scrambler Internal"
	case 1088:
		description = "CPL Scrambler Alarm Log"
	case 1089:
		description = "FF Annunciation"
	case 1090:
		description = "FF Fieldbus Message Specification"
	case 1091:
		description = "FF System Management"
	case 1092:
		description = "Open Business Reporting Protocol"
	case 1093:
		description = "PROOFD"
	case 1094:
		description = "ROOTD"
	case 1095:
		description = "NICELink"
	case 1096:
		description = "Common Name Resolution Protocol"
	case 1097:
		description = "Sun Cluster Manager"
	case 1098:
		description = "RMI Activation"
	case 1099:
		description = "RMI Registry"
	case 1100:
		description = "MCTP"
	case 1101:
		description = "PT2-DISCOVER"
	case 1102:
		description = "ADOBE SERVER 1"
	case 1103:
		description = "ADOBE SERVER 2"
	case 1104:
		description = "XRL"
	case 1105:
		description = "FTRANHC"
	case 1106:
		description = "ISOIPSIGPORT-1"
	case 1107:
		description = "ISOIPSIGPORT-2"
	case 1108:
		description = "ratio-adp"
	case 1110:
		description = "Start web admin server"
	case 1111:
		description = "LM Social Server"
	case 1112:
		description = "Intelligent Communication Protocol"
	case 1113:
		description = "Licklider Transmission Protocol"
	case 1114:
		description = "Mini SQL"
	case 1115:
		description = "ARDUS Transfer"
	case 1116:
		description = "ARDUS Control"
	case 1117:
		description = "ARDUS Multicast Transfer"
	case 1118:
		description = "SACRED"
	case 1119:
		description = "Battle.net Chat/Game Protocol"
	case 1120:
		description = "Battle.net File Transfer Protocol"
	case 1121:
		description = "Datalode RMPP"
	case 1122:
		description = "availant-mgr"
	case 1123:
		description = "Murray"
	case 1124:
		description = "HP VMM Control"
	case 1125:
		description = "HP VMM Agent"
	case 1126:
		description = "HP VMM Agent"
	case 1127:
		description = "KWDB Remote Communication"
	case 1128:
		description = "SAPHostControl over SOAP/HTTP"
	case 1129:
		description = "SAPHostControl over SOAP/HTTPS"
	case 1130:
		description = "CAC App Service Protocol"
	case 1131:
		description = "CAC App Service Protocol Encripted"
	case 1132:
		description = "KVM-via-IP Management Service"
	case 1133:
		description = "Data Flow Network"
	case 1134:
		description = "MicroAPL APLX"
	case 1135:
		description = "OmniVision Communication Service"
	case 1136:
		description = "HHB Gateway Control"
	case 1137:
		description = "TRIM Workgroup Service"
	case 1138:
		description = "encrypted admin requests"
	case 1139:
		description = "Enterprise Virtual Manager"
	case 1140:
		description = "AutoNOC Network Operations Protocol"
	case 1141:
		description = "User Message Service"
	case 1142:
		description = "User Discovery Service"
	case 1143:
		description = "Infomatryx Exchange"
	case 1144:
		description = "Fusion Script"
	case 1145:
		description = "X9 iCue Show Control"
	case 1146:
		description = "audit transfer"
	case 1147:
		description = "CAPIoverLAN"
	case 1148:
		description = "Elfiq Replication Service"
	case 1149:
		description = "BlueView Sonar Service"
	case 1150:
		description = "Blaze File Server"
	case 1151:
		description = "Unizensus Login Server"
	case 1152:
		description = "Winpopup LAN Messenger"
	case 1153:
		description = "ANSI C12.22 Port"
	case 1154:
		description = "Community Service"
	case 1155:
		description = "Network File Access"
	case 1156:
		description = "iasControl OMS"
	case 1157:
		description = "Oracle iASControl"
	case 1158:
		description = "dbControl OMS"
	case 1159:
		description = "Oracle OMS"
	case 1160:
		description = "DB Lite Mult-User Server"
	case 1161:
		description = "Health Polling"
	case 1162:
		description = "Health Trap"
	case 1163:
		description = "SmartDialer Data Protocol"
	case 1164:
		description = "QSM Proxy Service"
	case 1165:
		description = "QSM GUI Service"
	case 1166:
		description = "QSM RemoteExec"
	case 1167:
		description = "Cisco IP SLAs Control Protocol"
	case 1168:
		description = "VChat Conference Service"
	case 1169:
		description = "TRIPWIRE"
	case 1170:
		description = "AT+C License Manager"
	case 1171:
		description = "AT+C FmiApplicationServer"
	case 1172:
		description = "DNA Protocol"
	case 1173:
		description = "D-Cinema Request-Response"
	case 1174:
		description = "FlashNet Remote Admin"
	case 1175:
		description = "Dossier Server"
	case 1176:
		description = "Indigo Home Server"
	case 1177:
		description = "DKMessenger Protocol"
	case 1178:
		description = "SGI Storage Manager"
	case 1179:
		description = "Backup To Neighbor"
	case 1180:
		description = "Millicent Client Proxy"
	case 1181:
		description = "3Com Net Management"
	case 1182:
		description = "AcceleNet Control"
	case 1183:
		description = "LL Surfup HTTP"
	case 1184:
		description = "LL Surfup HTTPS"
	case 1185:
		description = "Catchpole port"
	case 1186:
		description = "MySQL Cluster Manager"
	case 1187:
		description = "Alias Service"
	case 1188:
		description = "HP Web Admin"
	case 1189:
		description = "Unet Connection"
	case 1190:
		description = "CommLinx GPS / AVL System"
	case 1191:
		description = "General Parallel File System"
	case 1192:
		description = "caids sensors channel"
	case 1193:
		description = "Five Across Server"
	case 1194:
		description = "OpenVPN"
	case 1195:
		description = "RSF-1 clustering"
	case 1196:
		description = "Network Magic"
	case 1197:
		description = "Carrius Remote Access"
	case 1198:
		description = "cajo reference discovery"
	case 1199:
		description = "DMIDI"
	case 1200:
		description = "SCOL"
	case 1201:
		description = "Nucleus Sand Database Server"
	case 1202:
		description = "caiccipc"
	case 1203:
		description = "License Validation"
	case 1204:
		description = "Log Request Listener"
	case 1205:
		description = "Accord-MGC"
	case 1206:
		description = "Anthony Data"
	case 1207:
		description = "MetaSage"
	case 1208:
		description = "SEAGULL AIS"
	case 1209:
		description = "IPCD3"
	case 1210:
		description = "EOSS"
	case 1211:
		description = "Groove DPP"
	case 1212:
		description = "lupa"
	case 1213:
		description = "Medtronic/Physio-Control LIFENET"
	case 1214:
		description = "KAZAA"
	case 1215:
		description = "scanSTAT 1.0"
	case 1216:
		description = "ETEBAC 5"
	case 1217:
		description = "HPSS NonDCE Gateway"
	case 1218:
		description = "AeroFlight-ADs"
	case 1219:
		description = "AeroFlight-Ret"
	case 1220:
		description = "QT SERVER ADMIN"
	case 1221:
		description = "SweetWARE Apps"
	case 1222:
		description = "SNI R&D network"
	case 1223:
		description = "TrulyGlobal Protocol"
	case 1224:
		description = "VPNz"
	case 1225:
		description = "SLINKYSEARCH"
	case 1226:
		description = "STGXFWS"
	case 1227:
		description = "DNS2Go"
	case 1228:
		description = "FLORENCE"
	case 1229:
		description = "ZENworks Tiered Electronic Distribution"
	case 1230:
		description = "Periscope"
	case 1231:
		description = "menandmice-lpm"
	case 1232:
		description = "Remote systems monitoring"
	case 1233:
		description = "Universal App Server"
	case 1234:
		description = "Infoseek Search Agent"
	case 1235:
		description = "mosaicsyssvc1"
	case 1236:
		description = "bvcontrol"
	case 1237:
		description = "tsdos390"
	case 1238:
		description = "hacl-qs"
	case 1239:
		description = "NMSD"
	case 1240:
		description = "Instantia"
	case 1241:
		description = "nessus"
	case 1242:
		description = "NMAS over IP"
	case 1243:
		description = "SerialGateway"
	case 1244:
		description = "isbconference1"
	case 1245:
		description = "isbconference2"
	case 1246:
		description = "payrouter"
	case 1247:
		description = "VisionPyramid"
	case 1248:
		description = "hermes"
	case 1249:
		description = "Mesa Vista Co"
	case 1250:
		description = "swldy-sias"
	case 1251:
		description = "servergraph"
	case 1252:
		description = "bspne-pcc"
	case 1253:
		description = "q55-pcc"
	case 1254:
		description = "de-noc"
	case 1255:
		description = "de-cache-query"
	case 1256:
		description = "de-server"
	case 1257:
		description = "Shockwave 2"
	case 1258:
		description = "Open Network Library"
	case 1259:
		description = "Open Network Library Voice"
	case 1260:
		description = "ibm-ssd"
	case 1261:
		description = "mpshrsv"
	case 1262:
		description = "QNTS-ORB"
	case 1263:
		description = "dka"
	case 1264:
		description = "PRAT"
	case 1265:
		description = "DSSIAPI"
	case 1266:
		description = "DELLPWRAPPKS"
	case 1267:
		description = "eTrust Policy Compliance"
	case 1268:
		description = "PROPEL-MSGSYS"
	case 1269:
		description = "WATiLaPP"
	case 1270:
		description = "Microsoft Operations Manager"
	case 1271:
		description = "eXcW"
	case 1272:
		description = "CSPMLockMgr"
	case 1273:
		description = "EMC-Gateway"
	case 1274:
		description = "t1distproc"
	case 1275:
		description = "ivcollector"
	case 1277:
		description = "mqs"
	case 1278:
		description = "Dell Web Admin 1"
	case 1279:
		description = "Dell Web Admin 2"
	case 1280:
		description = "Pictrography"
	case 1281:
		description = "healthd"
	case 1282:
		description = "Emperion"
	case 1283:
		description = "Product Information"
	case 1284:
		description = "IEE-QFX"
	case 1285:
		description = "neoiface"
	case 1286:
		description = "netuitive"
	case 1287:
		description = "RouteMatch Com"
	case 1288:
		description = "NavBuddy"
	case 1289:
		description = "JWalkServer"
	case 1290:
		description = "WinJaServer"
	case 1291:
		description = "SEAGULLLMS"
	case 1292:
		description = "dsdn"
	case 1293:
		description = "PKT-KRB-IPSec"
	case 1294:
		description = "CMMdriver"
	case 1295:
		description = "End-by-Hop Transmission Protocol"
	case 1296:
		description = "dproxy"
	case 1297:
		description = "sdproxy"
	case 1298:
		description = "lpcp"
	case 1299:
		description = "hp-sci"
	case 1300:
		description = "H.323 Secure Call Control Signalling"
	case 1303:
		description = "sftsrv"
	case 1304:
		description = "Boomerang"
	case 1305:
		description = "pe-mike"
	case 1306:
		description = "RE-Conn-Proto"
	case 1307:
		description = "Pacmand"
	case 1308:
		description = "Optical Domain Service Interconnect (ODSI)"
	case 1309:
		description = "JTAG server"
	case 1310:
		description = "Husky"
	case 1311:
		description = "RxMon"
	case 1312:
		description = "STI Envision"
	case 1313:
		description = "BMC_PATROLDB"
	case 1314:
		description = "Photoscript Distributed Printing System"
	case 1315:
		description = "E.L.S."
	case 1316:
		description = "Exbit-ESCP"
	case 1317:
		description = "vrts-ipcserver"
	case 1318:
		description = "krb5gatekeeper"
	case 1319:
		description = "AMX-ICSP"
	case 1320:
		description = "AMX-AXBNET"
	case 1321:
		description = "PIP"
	case 1322:
		description = "Novation"
	case 1323:
		description = "brcd"
	case 1324:
		description = "delta-mcp"
	case 1325:
		description = "DX-Instrument"
	case 1326:
		description = "WIMSIC"
	case 1327:
		description = "Ultrex"
	case 1328:
		description = "EWALL"
	case 1329:
		description = "netdb-export"
	case 1330:
		description = "StreetPerfect"
	case 1331:
		description = "intersan"
	case 1332:
		description = "PCIA RXP-B"
	case 1333:
		description = "Password Policy"
	case 1334:
		description = "writesrv"
	case 1335:
		description = "Digital Notary Protocol"
	case 1336:
		description = "Instant Service Chat"
	case 1337:
		description = "menandmice DNS"
	case 1338:
		description = "WMC-log-svr"
	case 1339:
		description = "kjtsiteserver"
	case 1340:
		description = "NAAP"
	case 1341:
		description = "QuBES"
	case 1342:
		description = "ESBroker"
	case 1343:
		description = "re101"
	case 1344:
		description = "ICAP"
	case 1345:
		description = "VPJP"
	case 1346:
		description = "Alta Analytics License Manager"
	case 1347:
		description = "multi media conferencing"
	case 1348:
		description = "multi media conferencing"
	case 1349:
		description = "Registration Network Protocol"
	case 1350:
		description = "Registration Network Protocol"
	case 1351:
		description = "Digital Tool Works (MIT)"
	case 1352:
		description = "Lotus Note"
	case 1353:
		description = "Relief Consulting"
	case 1354:
		description = "Five Across XSIP Network"
	case 1355:
		description = "Intuitive Edge"
	case 1356:
		description = "CuillaMartin Company"
	case 1357:
		description = "Electronic PegBoard"
	case 1358:
		description = "CONNLCLI"
	case 1359:
		description = "FTSRV"
	case 1360:
		description = "MIMER"
	case 1361:
		description = "LinX"
	case 1362:
		description = "TimeFlies"
	case 1363:
		description = "Network DataMover Requester"
	case 1364:
		description = "Network DataMover Server"
	case 1365:
		description = "Network Software Associates"
	case 1366:
		description = "Novell NetWare Comm Service Platform"
	case 1367:
		description = "DCS"
	case 1368:
		description = "ScreenCast"
	case 1369:
		description = "GlobalView to Unix Shell"
	case 1370:
		description = "Unix Shell to GlobalView"
	case 1371:
		description = "Fujitsu Config Protocol"
	case 1372:
		description = "Fujitsu Config Protocol"
	case 1373:
		description = "Chromagrafx"
	case 1374:
		description = "EPI Software Systems"
	case 1375:
		description = "Bytex"
	case 1376:
		description = "IBM Person to Person Software"
	case 1377:
		description = "Cichlid License Manager"
	case 1378:
		description = "Elan License Manager"
	case 1379:
		description = "Integrity Solutions"
	case 1380:
		description = "Telesis Network License Manager"
	case 1381:
		description = "Apple Network License Manager"
	case 1382:
		description = "udt_os"
	case 1383:
		description = "GW Hannaway Network License Manager"
	case 1384:
		description = "Objective Solutions License Manager"
	case 1385:
		description = "Atex Publishing License Manager"
	case 1386:
		description = "CheckSum License Manager"
	case 1387:
		description = "Computer Aided Design Software Inc LM"
	case 1388:
		description = "Objective Solutions DataBase Cache"
	case 1389:
		description = "Document Manager"
	case 1390:
		description = "Storage Controller"
	case 1391:
		description = "Storage Access Server"
	case 1392:
		description = "Print Manager"
	case 1393:
		description = "Network Log Server"
	case 1394:
		description = "Network Log Client"
	case 1395:
		description = "PC Workstation Manager software"
	case 1396:
		description = "DVL Active Mail"
	case 1397:
		description = "Audio Active Mail"
	case 1398:
		description = "Video Active Mail"
	case 1399:
		description = "Cadkey License Manager"
	case 1400:
		description = "Cadkey Tablet Daemon"
	case 1401:
		description = "Goldleaf License Manager"
	case 1402:
		description = "Prospero Resource Manager"
	case 1403:
		description = "Prospero Resource Manager"
	case 1404:
		description = "Infinite Graphics License Manager"
	case 1405:
		description = "IBM Remote Execution Starter"
	case 1406:
		description = "NetLabs License Manager"
	case 1407:
		description = "TIBET Data Server"
	case 1408:
		description = "Sophia License Manager"
	case 1409:
		description = "Here License Manager"
	case 1410:
		description = "HiQ License Manager"
	case 1411:
		description = "AudioFile"
	case 1412:
		description = "InnoSys"
	case 1413:
		description = "Innosys-ACL"
	case 1414:
		description = "IBM MQSeries"
	case 1415:
		description = "DBStar"
	case 1416:
		description = "Novell LU6.2"
	case 1417:
		description = "Timbuktu Service 1 Port"
	case 1418:
		description = "Timbuktu Service 2 Port"
	case 1419:
		description = "Timbuktu Service 3 Port"
	case 1420:
		description = "Timbuktu Service 4 Port"
	case 1421:
		description = "Gandalf License Manager"
	case 1422:
		description = "Autodesk License Manager"
	case 1423:
		description = "Essbase Arbor Software"
	case 1424:
		description = "Hybrid Encryption Protocol"
	case 1425:
		description = "Zion Software License Manager"
	case 1426:
		description = "Satellite-data Acquisition System 1"
	case 1427:
		description = "mloadd monitoring tool"
	case 1428:
		description = "Informatik License Manager"
	case 1429:
		description = "Hypercom NMS"
	case 1430:
		description = "Hypercom TPDU"
	case 1431:
		description = "Reverse Gossip Transport"
	case 1432:
		description = "Blueberry Software License Manager"
	case 1433:
		description = "Microsoft-SQL-Server"
	case 1434:
		description = "Microsoft-SQL-Monitor"
	case 1435:
		description = "IBM CICS"
	case 1436:
		description = "Satellite-data Acquisition System 2"
	case 1437:
		description = "Tabula"
	case 1438:
		description = "Eicon Security Agent/Server"
	case 1439:
		description = "Eicon X25/SNA Gateway"
	case 1440:
		description = "Eicon Service Location Protocol"
	case 1441:
		description = "Cadis License Management"
	case 1442:
		description = "Cadis License Management"
	case 1443:
		description = "Integrated Engineering Software"
	case 1444:
		description = "Marcam  License Management"
	case 1445:
		description = "Proxima License Manager"
	case 1446:
		description = "Optical Research Associates License Manager"
	case 1447:
		description = "Applied Parallel Research LM"
	case 1448:
		description = "OpenConnect License Manager"
	case 1449:
		description = "PEport"
	case 1450:
		description = "Tandem Distributed Workbench Facility"
	case 1451:
		description = "IBM Information Management"
	case 1452:
		description = "GTE Government Systems License Man"
	case 1453:
		description = "Genie License Manager"
	case 1454:
		description = "interHDL License Manager"
	case 1455:
		description = "ESL License Manager"
	case 1456:
		description = "DCA"
	case 1457:
		description = "Valisys License Manager"
	case 1458:
		description = "Nichols Research Corp."
	case 1459:
		description = "Proshare Notebook Application"
	case 1460:
		description = "Proshare Notebook Application"
	case 1461:
		description = "IBM Wireless LAN"
	case 1462:
		description = "World License Manager"
	case 1463:
		description = "Nucleus"
	case 1464:
		description = "MSL License Manager"
	case 1465:
		description = "Pipes Platform"
	case 1466:
		description = "Ocean Software License Manager"
	case 1467:
		description = "CSDMBASE"
	case 1468:
		description = "CSDM"
	case 1469:
		description = "Active Analysis Limited License Manager"
	case 1470:
		description = "Universal Analytics"
	case 1471:
		description = "csdmbase"
	case 1472:
		description = "csdm"
	case 1473:
		description = "OpenMath"
	case 1474:
		description = "Telefinder"
	case 1475:
		description = "Taligent License Manager"
	case 1476:
		description = "clvm-cfg"
	case 1477:
		description = "ms-sna-server"
	case 1478:
		description = "ms-sna-base"
	case 1479:
		description = "dberegister"
	case 1480:
		description = "PacerForum"
	case 1481:
		description = "AIRS"
	case 1482:
		description = "Miteksys License Manager"
	case 1483:
		description = "AFS License Manager"
	case 1484:
		description = "Confluent License Manager"
	case 1485:
		description = "LANSource"
	case 1486:
		description = "nms_topo_serv"
	case 1487:
		description = "LocalInfoSrvr"
	case 1488:
		description = "DocStor"
	case 1489:
		description = "dmdocbroker"
	case 1490:
		description = "insitu-conf"
	case 1492:
		description = "stone-design-1"
	case 1493:
		description = "netmap_lm"
	case 1494:
		description = "ica"
	case 1495:
		description = "cvc"
	case 1496:
		description = "liberty-lm"
	case 1497:
		description = "rfx-lm"
	case 1498:
		description = "Sybase SQL Any"
	case 1499:
		description = "Federico Heinz Consultora"
	case 1500:
		description = "VLSI License Manager"
	case 1501:
		description = "Satellite-data Acquisition System 3"
	case 1502:
		description = "Shiva"
	case 1503:
		description = "Databeam"
	case 1504:
		description = "EVB Software Engineering License Manager"
	case 1505:
		description = "Funk Software"
	case 1506:
		description = "Universal Time daemon (utcd)"
	case 1507:
		description = "symplex"
	case 1508:
		description = "diagmond"
	case 1509:
		description = "Robcad"
	case 1510:
		description = "Midland Valley Exploration Ltd. Lic. Man."
	case 1511:
		description = "3l-l1"
	case 1512:
		description = "Microsoft's Windows Internet Name Service"
	case 1513:
		description = "Fujitsu Systems Business of America"
	case 1514:
		description = "Fujitsu Systems Business of America"
	case 1515:
		description = "ifor-protocol"
	case 1516:
		description = "Virtual Places Audio data"
	case 1517:
		description = "Virtual Places Audio control"
	case 1518:
		description = "Virtual Places Video data"
	case 1519:
		description = "Virtual Places Video control"
	case 1520:
		description = "atm zip office"
	case 1521:
		description = "nCube License Manager"
	case 1522:
		description = "Ricardo North America License Manager"
	case 1523:
		description = "cichild"
	case 1524:
		description = "ingres"
	case 1525:
		description = "oracle"
	case 1526:
		description = "Prospero Data Access Prot non-priv"
	case 1527:
		description = "oracle"
	case 1528:
		description = "Not Only a Routeing Protocol"
	case 1529:
		description = "oracle"
	case 1530:
		description = "rap-service"
	case 1531:
		description = "rap-listen"
	case 1532:
		description = "miroconnect"
	case 1533:
		description = "Virtual Places Software"
	case 1534:
		description = "micromuse-lm"
	case 1535:
		description = "ampr-info"
	case 1536:
		description = "ampr-inter"
	case 1537:
		description = "isi-lm"
	case 1538:
		description = "3ds-lm"
	case 1539:
		description = "Intellistor License Manager"
	case 1540:
		description = "rds"
	case 1541:
		description = "rds2"
	case 1542:
		description = "gridgen-elmd"
	case 1543:
		description = "simba-cs"
	case 1544:
		description = "aspeclmd"
	case 1545:
		description = "vistium-share"
	case 1546:
		description = "abbaccuray"
	case 1547:
		description = "laplink"
	case 1548:
		description = "Axon License Manager"
	case 1549:
		description = "Shiva Hose"
	case 1550:
		description = "Image Storage license manager 3M Company"
	case 1551:
		description = "HECMTL-DB"
	case 1552:
		description = "pciarray"
	case 1553:
		description = "sna-cs"
	case 1554:
		description = "CACI Products Company License Manager"
	case 1555:
		description = "livelan"
	case 1556:
		description = "VERITAS Private Branch Exchange"
	case 1557:
		description = "ArborText License Manager"
	case 1558:
		description = "xingmpeg"
	case 1559:
		description = "web2host"
	case 1560:
		description = "ASCI-RemoteSHADOW"
	case 1561:
		description = "facilityview"
	case 1562:
		description = "pconnectmgr"
	case 1563:
		description = "Cadabra License Manager"
	case 1564:
		description = "Pay-Per-View"
	case 1565:
		description = "WinDD"
	case 1566:
		description = "CORELVIDEO"
	case 1567:
		description = "jlicelmd"
	case 1568:
		description = "tsspmap"
	case 1569:
		description = "ets"
	case 1570:
		description = "orbixd"
	case 1571:
		description = "Oracle Remote Data Base"
	case 1572:
		description = "Chipcom License Manager"
	case 1573:
		description = "itscomm-ns"
	case 1574:
		description = "mvel-lm"
	case 1575:
		description = "oraclenames"
	case 1576:
		description = "Moldflow License Manager"
	case 1577:
		description = "hypercube-lm"
	case 1578:
		description = "Jacobus License Manager"
	case 1579:
		description = "ioc-sea-lm"
	case 1580:
		description = "tn-tl-r1"
	case 1581:
		description = "MIL-2045-47001"
	case 1582:
		description = "MSIMS"
	case 1583:
		description = "simbaexpress"
	case 1584:
		description = "tn-tl-fd2"
	case 1585:
		description = "intv"
	case 1586:
		description = "ibm-abtact"
	case 1587:
		description = "pra_elmd"
	case 1588:
		description = "triquest-lm"
	case 1589:
		description = "VQP"
	case 1590:
		description = "gemini-lm"
	case 1591:
		description = "ncpm-pm"
	case 1592:
		description = "commonspace"
	case 1593:
		description = "mainsoft-lm"
	case 1594:
		description = "sixtrak"
	case 1595:
		description = "radio"
	case 1596:
		description = "radio-sm"
	case 1597:
		description = "orbplus-iiop"
	case 1598:
		description = "picknfs"
	case 1599:
		description = "simbaservices"
	case 1600:
		description = "issd"
	case 1601:
		description = "aas"
	case 1602:
		description = "inspect"
	case 1603:
		description = "pickodbc"
	case 1604:
		description = "icabrowser"
	case 1605:
		description = "Salutation Manager (Salutation Protocol)"
	case 1606:
		description = "Salutation Manager (SLM-API)"
	case 1607:
		description = "stt"
	case 1608:
		description = "Smart Corp. License Manager"
	case 1609:
		description = "isysg-lm"
	case 1610:
		description = "taurus-wh"
	case 1611:
		description = "Inter Library Loan"
	case 1612:
		description = "NetBill Transaction Server"
	case 1613:
		description = "NetBill Key Repository"
	case 1614:
		description = "NetBill Credential Server"
	case 1615:
		description = "NetBill Authorization Server"
	case 1616:
		description = "NetBill Product Server"
	case 1617:
		description = "Nimrod Inter-Agent Communication"
	case 1618:
		description = "skytelnet"
	case 1619:
		description = "xs-openstorage"
	case 1620:
		description = "faxportwinport"
	case 1621:
		description = "softdataphone"
	case 1622:
		description = "ontime"
	case 1623:
		description = "jaleosnd"
	case 1624:
		description = "udp-sr-port"
	case 1625:
		description = "svs-omagent"
	case 1626:
		description = "Shockwave"
	case 1627:
		description = "T.128 Gateway"
	case 1628:
		description = "LonTalk normal"
	case 1629:
		description = "LonTalk urgent"
	case 1630:
		description = "Oracle Net8 Cman"
	case 1631:
		description = "Visit view"
	case 1632:
		description = "PAMMRATC"
	case 1633:
		description = "PAMMRPC"
	case 1634:
		description = "Log On America Probe"
	case 1635:
		description = "EDB Server 1"
	case 1636:
		description = "ISP shared public data control"
	case 1637:
		description = "ISP shared local data control"
	case 1638:
		description = "ISP shared management control"
	case 1639:
		description = "cert-initiator"
	case 1640:
		description = "cert-responder"
	case 1641:
		description = "InVision"
	case 1642:
		description = "isis-am"
	case 1643:
		description = "isis-ambc"
	case 1644:
		description = "Satellite-data Acquisition System 4"
	case 1645:
		description = "SightLine"
	case 1646:
		description = "sa-msg-port"
	case 1647:
		description = "rsap"
	case 1648:
		description = "concurrent-lm"
	case 1649:
		description = "kermit"
	case 1650:
		description = "nkd"
	case 1651:
		description = "shiva_confsrvr"
	case 1652:
		description = "xnmp"
	case 1653:
		description = "alphatech-lm"
	case 1654:
		description = "stargatealerts"
	case 1655:
		description = "dec-mbadmin"
	case 1656:
		description = "dec-mbadmin-h"
	case 1657:
		description = "fujitsu-mmpdc"
	case 1658:
		description = "sixnetudr"
	case 1659:
		description = "Silicon Grail License Manager"
	case 1660:
		description = "skip-mc-gikreq"
	case 1661:
		description = "netview-aix-1"
	case 1662:
		description = "netview-aix-2"
	case 1663:
		description = "netview-aix-3"
	case 1664:
		description = "netview-aix-4"
	case 1665:
		description = "netview-aix-5"
	case 1666:
		description = "netview-aix-6"
	case 1667:
		description = "netview-aix-7"
	case 1668:
		description = "netview-aix-8"
	case 1669:
		description = "netview-aix-9"
	case 1670:
		description = "netview-aix-10"
	case 1671:
		description = "netview-aix-11"
	case 1672:
		description = "netview-aix-12"
	case 1673:
		description = "Intel Proshare Multicast"
	case 1674:
		description = "Intel Proshare Multicast"
	case 1675:
		description = "Pacific Data Products"
	case 1676:
		description = "netcomm1"
	case 1677:
		description = "groupwise"
	case 1678:
		description = "prolink"
	case 1679:
		description = "darcorp-lm"
	case 1680:
		description = "microcom-sbp"
	case 1681:
		description = "sd-elmd"
	case 1682:
		description = "lanyon-lantern"
	case 1683:
		description = "ncpm-hip"
	case 1684:
		description = "SnareSecure"
	case 1685:
		description = "n2nremote"
	case 1686:
		description = "cvmon"
	case 1687:
		description = "nsjtp-ctrl"
	case 1688:
		description = "nsjtp-data"
	case 1689:
		description = "firefox"
	case 1690:
		description = "ng-umds"
	case 1691:
		description = "empire-empuma"
	case 1692:
		description = "sstsys-lm"
	case 1693:
		description = "rrirtr"
	case 1694:
		description = "rrimwm"
	case 1695:
		description = "rrilwm"
	case 1696:
		description = "rrifmm"
	case 1697:
		description = "rrisat"
	case 1698:
		description = "RSVP-ENCAPSULATION-1"
	case 1699:
		description = "RSVP-ENCAPSULATION-2"
	case 1700:
		description = "mps-raft"
	case 1701:
		description = "l2f"
	case 1702:
		description = "deskshare"
	case 1703:
		description = "hb-engine"
	case 1704:
		description = "bcs-broker"
	case 1705:
		description = "slingshot"
	case 1706:
		description = "jetform"
	case 1707:
		description = "vdmplay"
	case 1708:
		description = "gat-lmd"
	case 1709:
		description = "centra"
	case 1710:
		description = "impera"
	case 1711:
		description = "pptconference"
	case 1712:
		description = "resource monitoring service"
	case 1713:
		description = "ConferenceTalk"
	case 1714:
		description = "sesi-lm"
	case 1715:
		description = "houdini-lm"
	case 1716:
		description = "xmsg"
	case 1717:
		description = "fj-hdnet"
	case 1718:
		description = "H.323 Multicast Gatekeeper Discover"
	case 1719:
		description = "H.323 Unicast Gatekeeper Signaling"
	case 1720:
		description = "H.323 Call Control"
	case 1721:
		description = "caicci"
	case 1722:
		description = "HKS License Manager"
	case 1723:
		description = "pptp"
	case 1724:
		description = "csbphonemaster"
	case 1725:
		description = "iden-ralp"
	case 1726:
		description = "IBERIAGAMES"
	case 1727:
		description = "winddx"
	case 1728:
		description = "TELINDUS"
	case 1729:
		description = "CityNL License Management"
	case 1730:
		description = "roketz"
	case 1731:
		description = "MSICCP"
	case 1732:
		description = "proxim"
	case 1733:
		description = "SIMS - SIIPAT"
	case 1734:
		description = "Camber Corporation License Management"
	case 1735:
		description = "PrivateChat"
	case 1736:
		description = "street-stream"
	case 1737:
		description = "ultimad"
	case 1738:
		description = "GameGen1"
	case 1739:
		description = "webaccess"
	case 1740:
		description = "encore"
	case 1741:
		description = "cisco-net-mgmt"
	case 1742:
		description = "3Com-nsd"
	case 1743:
		description = "Cinema Graphics License Manager"
	case 1744:
		description = "ncpm-ft"
	case 1745:
		description = "remote-winsock"
	case 1746:
		description = "ftrapid-1"
	case 1747:
		description = "ftrapid-2"
	case 1748:
		description = "oracle-em1"
	case 1749:
		description = "aspen-services"
	case 1750:
		description = "Simple Socket Library's PortMaster"
	case 1751:
		description = "SwiftNet"
	case 1752:
		description = "Leap of Faith Research License Manager"
	case 1753:
		description = "Predatar Comms Service"
	case 1754:
		description = "oracle-em2"
	case 1755:
		description = "ms-streaming"
	case 1756:
		description = "capfast-lmd"
	case 1757:
		description = "cnhrp"
	case 1758:
		description = "tftp-mcast"
	case 1759:
		description = "SPSS License Manager"
	case 1760:
		description = "www-ldap-gw"
	case 1761:
		description = "cft-0"
	case 1762:
		description = "cft-1"
	case 1763:
		description = "cft-2"
	case 1764:
		description = "cft-3"
	case 1765:
		description = "cft-4"
	case 1766:
		description = "cft-5"
	case 1767:
		description = "cft-6"
	case 1768:
		description = "cft-7"
	case 1769:
		description = "bmc-net-adm"
	case 1770:
		description = "bmc-net-svc"
	case 1771:
		description = "vaultbase"
	case 1772:
		description = "EssWeb Gateway"
	case 1773:
		description = "KMSControl"
	case 1774:
		description = "global-dtserv"
	case 1775:
		description = "data interchange between visual processing containers"
	case 1776:
		description = "Federal Emergency Management Information System"
	case 1777:
		description = "powerguardian"
	case 1778:
		description = "prodigy-internet"
	case 1779:
		description = "pharmasoft"
	case 1780:
		description = "dpkeyserv"
	case 1781:
		description = "answersoft-lm"
	case 1782:
		description = "hp-hcip"
	case 1783:
		description = "Decomissioned Port 04/14/00"
	case 1784:
		description = "Finle License Manager"
	case 1785:
		description = "Wind River Systems License Manager"
	case 1786:
		description = "funk-logger"
	case 1787:
		description = "funk-license"
	case 1788:
		description = "psmond"
	case 1789:
		description = "hello"
	case 1790:
		description = "Narrative Media Streaming Protocol"
	case 1791:
		description = "EA1"
	case 1792:
		description = "ibm-dt-2"
	case 1793:
		description = "rsc-robot"
	case 1794:
		description = "cera-bcm"
	case 1795:
		description = "dpi-proxy"
	case 1796:
		description = "Vocaltec Server Administration"
	case 1797:
		description = "UMA"
	case 1798:
		description = "Event Transfer Protocol"
	case 1799:
		description = "NETRISK"
	case 1800:
		description = "ANSYS-License manager"
	case 1801:
		description = "Microsoft Message Que"
	case 1802:
		description = "ConComp1"
	case 1803:
		description = "HP-HCIP-GWY"
	case 1804:
		description = "ENL"
	case 1805:
		description = "ENL-Name"
	case 1806:
		description = "Musiconline"
	case 1807:
		description = "Fujitsu Hot Standby Protocol"
	case 1808:
		description = "Oracle-VP2"
	case 1809:
		description = "Oracle-VP1"
	case 1810:
		description = "Jerand License Manager"
	case 1811:
		description = "Scientia-SDB"
	case 1812:
		description = "RADIUS"
	case 1813:
		description = "RADIUS Accounting"
	case 1814:
		description = "TDP Suite"
	case 1815:
		description = "Manufacturing messaging protocol for factory transmission"
	case 1816:
		description = "HARP"
	case 1817:
		description = "RKB-OSCS"
	case 1818:
		description = "Enhanced Trivial File Transfer Protocol"
	case 1819:
		description = "Plato License Manager"
	case 1820:
		description = "mcagent"
	case 1821:
		description = "donnyworld"
	case 1822:
		description = "es-elmd"
	case 1823:
		description = "Unisys Natural Language License Manager"
	case 1824:
		description = "metrics-pas"
	case 1825:
		description = "DirecPC Video"
	case 1826:
		description = "ARDT"
	case 1827:
		description = "ASI"
	case 1828:
		description = "itm-mcell-u"
	case 1829:
		description = "Optika eMedia"
	case 1830:
		description = "Oracle Net8 CMan Admin"
	case 1831:
		description = "Myrtle"
	case 1832:
		description = "ThoughtTreasure"
	case 1833:
		description = "udpradio"
	case 1834:
		description = "ARDUS Unicast"
	case 1835:
		description = "ARDUS Multicast"
	case 1836:
		description = "ste-smsc"
	case 1837:
		description = "csoft1"
	case 1838:
		description = "TALNET"
	case 1839:
		description = "netopia-vo1"
	case 1840:
		description = "netopia-vo2"
	case 1841:
		description = "netopia-vo3"
	case 1842:
		description = "netopia-vo4"
	case 1843:
		description = "netopia-vo5"
	case 1844:
		description = "DirecPC-DLL"
	case 1845:
		description = "altalink"
	case 1846:
		description = "Tunstall PNC"
	case 1847:
		description = "SLP Notification"
	case 1848:
		description = "fjdocdist"
	case 1849:
		description = "ALPHA-SMS"
	case 1850:
		description = "GSI"
	case 1851:
		description = "ctcd"
	case 1852:
		description = "Virtual Time"
	case 1853:
		description = "VIDS-AVTP"
	case 1854:
		description = "Buddy Draw"
	case 1855:
		description = "Fiorano RtrSvc"
	case 1856:
		description = "Fiorano MsgSvc"
	case 1857:
		description = "DataCaptor"
	case 1858:
		description = "PrivateArk"
	case 1859:
		description = "Gamma Fetcher Server"
	case 1860:
		description = "SunSCALAR Services"
	case 1861:
		description = "LeCroy VICP"
	case 1862:
		description = "MySQL Cluster Manager Agent"
	case 1863:
		description = "MSNP"
	case 1864:
		description = "Paradym 31 Port"
	case 1865:
		description = "ENTP"
	case 1866:
		description = "swrmi"
	case 1867:
		description = "UDRIVE"
	case 1868:
		description = "VizibleBrowser"
	case 1869:
		description = "TransAct"
	case 1870:
		description = "SunSCALAR DNS Service"
	case 1871:
		description = "Cano Central 0"
	case 1872:
		description = "Cano Central 1"
	case 1873:
		description = "Fjmpjps"
	case 1874:
		description = "Fjswapsnp"
	case 1875:
		description = "westell stats"
	case 1876:
		description = "ewcappsrv"
	case 1877:
		description = "hp-webqosdb"
	case 1878:
		description = "drmsmc"
	case 1879:
		description = "NettGain NMS"
	case 1880:
		description = "Gilat VSAT Control"
	case 1881:
		description = "IBM WebSphere MQ Everyplace"
	case 1882:
		description = "CA eTrust Common Services"
	case 1883:
		description = "Message Queuing Telemetry Transport Protocol"
	case 1884:
		description = "Internet Distance Map Svc"
	case 1885:
		description = "Veritas Trap Server"
	case 1886:
		description = "Leonardo over IP"
	case 1887:
		description = "FileX Listening Port"
	case 1888:
		description = "NC Config Port"
	case 1889:
		description = "Unify Web Adapter Service"
	case 1890:
		description = "wilkenListener"
	case 1891:
		description = "ChildKey Notification"
	case 1892:
		description = "ChildKey Control"
	case 1893:
		description = "ELAD Protocol"
	case 1894:
		description = "O2Server Port"
	case 1896:
		description = "b-novative license server"
	case 1897:
		description = "MetaAgent"
	case 1898:
		description = "Cymtec secure management"
	case 1899:
		description = "MC2Studios"
	case 1900:
		description = "SSDP"
	case 1901:
		description = "Fujitsu ICL Terminal Emulator Program A"
	case 1902:
		description = "Fujitsu ICL Terminal Emulator Program B"
	case 1903:
		description = "Local Link Name Resolution"
	case 1904:
		description = "Fujitsu ICL Terminal Emulator Program C"
	case 1905:
		description = "Secure UP.Link Gateway Protocol"
	case 1906:
		description = "TPortMapperReq"
	case 1907:
		description = "IntraSTAR"
	case 1908:
		description = "Dawn"
	case 1909:
		description = "Global World Link"
	case 1910:
		description = "UltraBac Software communications port"
	case 1911:
		description = "Starlight Networks Multimedia Transport Protocol"
	case 1912:
		description = "rhp-iibp"
	case 1913:
		description = "armadp"
	case 1914:
		description = "Elm-Momentum"
	case 1915:
		description = "FACELINK"
	case 1916:
		description = "Persoft Persona"
	case 1917:
		description = "nOAgent"
	case 1918:
		description = "IBM Tivole Directory Service - NDS"
	case 1919:
		description = "IBM Tivoli Directory Service - DCH"
	case 1920:
		description = "IBM Tivoli Directory Service - FERRET"
	case 1921:
		description = "NoAdmin"
	case 1922:
		description = "Tapestry"
	case 1923:
		description = "SPICE"
	case 1924:
		description = "XIIP"
	case 1925:
		description = "Surrogate Discovery Port"
	case 1926:
		description = "Evolution Game Server"
	case 1927:
		description = "Videte CIPC Port"
	case 1928:
		description = "Expnd Maui Srvr Dscovr"
	case 1929:
		description = "Bandwiz System - Server"
	case 1930:
		description = "Drive AppServer"
	case 1931:
		description = "AMD SCHED"
	case 1932:
		description = "CTT Broker"
	case 1933:
		description = "IBM LM MT Agent"
	case 1934:
		description = "IBM LM Appl Agent"
	case 1935:
		description = "Macromedia Flash Communications Server MX"
	case 1936:
		description = "JetCmeServer Server Port"
	case 1937:
		description = "JetVWay Server Port"
	case 1938:
		description = "JetVWay Client Port"
	case 1939:
		description = "JetVision Server Port"
	case 1940:
		description = "JetVision Client Port"
	case 1941:
		description = "DIC-Aida"
	case 1942:
		description = "Real Enterprise Service"
	case 1943:
		description = "Beeyond Media"
	case 1944:
		description = "close-combat"
	case 1945:
		description = "dialogic-elmd"
	case 1946:
		description = "tekpls"
	case 1947:
		description = "SentinelSRM"
	case 1948:
		description = "eye2eye"
	case 1949:
		description = "ISMA Easdaq Live"
	case 1950:
		description = "ISMA Easdaq Test"
	case 1951:
		description = "bcs-lmserver"
	case 1952:
		description = "mpnjsc"
	case 1953:
		description = "Rapid Base"
	case 1954:
		description = "ABR-API (diskbridge)"
	case 1955:
		description = "ABR-Secure Data (diskbridge)"
	case 1956:
		description = "Vertel VMF DS"
	case 1957:
		description = "unix-status"
	case 1958:
		description = "CA Administration Daemon"
	case 1959:
		description = "SIMP Channel"
	case 1960:
		description = "Merit DAC NASmanager"
	case 1961:
		description = "BTS APPSERVER"
	case 1962:
		description = "BIAP-MP"
	case 1963:
		description = "WebMachine"
	case 1964:
		description = "SOLID E ENGINE"
	case 1965:
		description = "Tivoli NPM"
	case 1966:
		description = "Slush"
	case 1967:
		description = "SNS Quote"
	case 1968:
		description = "LIPSinc"
	case 1969:
		description = "LIPSinc 1"
	case 1970:
		description = "NetOp Remote Control"
	case 1971:
		description = "NetOp School"
	case 1972:
		description = "Cache"
	case 1973:
		description = "Data Link Switching Remote Access Protocol"
	case 1974:
		description = "DRP"
	case 1975:
		description = "TCO Flash Agent"
	case 1976:
		description = "TCO Reg Agent"
	case 1977:
		description = "TCO Address Book"
	case 1978:
		description = "UniSQL"
	case 1979:
		description = "UniSQL Java"
	case 1980:
		description = "PearlDoc XACT"
	case 1981:
		description = "p2pQ"
	case 1982:
		description = "Evidentiary Timestamp"
	case 1983:
		description = "Loophole Test Protocol"
	case 1984:
		description = "BB"
	case 1985:
		description = "Hot Standby Router Protocol"
	case 1986:
		description = "cisco license management"
	case 1987:
		description = "cisco RSRB Priority 1 port"
	case 1988:
		description = "cisco RSRB Priority 2 port"
	case 1989:
		description = "cisco RSRB Priority 3 port"
	case 1990:
		description = "cisco STUN Priority 1 port"
	case 1991:
		description = "cisco STUN Priority 2 port"
	case 1992:
		description = "cisco STUN Priority 3 port"
	case 1993:
		description = "cisco SNMP TCP port"
	case 1994:
		description = "cisco serial tunnel port"
	case 1995:
		description = "cisco perf port"
	case 1996:
		description = "cisco Remote SRB port"
	case 1997:
		description = "cisco Gateway Discovery Protocol"
	case 1998:
		description = "cisco X.25 service (XOT)"
	case 1999:
		description = "cisco identification port"
	case 2000:
		description = "Cisco SCCP"
	case 2001:
		description = "curry"
	case 2003:
		description = "Brutus Server"
	case 2004:
		description = "CCWS mm conf"
	case 2006:
		description = "raid"
	case 2010:
		description = "pipe_server."
	case 2011:
		description = "raid"
	case 2029:
		description = "Hot Standby Router Protocol IPv6"
	case 2031:
		description = "mobrien-chat"
	case 2036:
		description = "Ethernet WS DP network"
	case 2037:
		description = "APplus Application Server"
	case 2039:
		description = "Prizma Monitoring Service"
	case 2042:
		description = "isis"
	case 2043:
		description = "isis-bcast"
	case 2049:
		description = "Network File System"
	case 2050:
		description = "Avaya EMB Config Port"
	case 2051:
		description = "EPNSDP"
	case 2052:
		description = "clearVisn Services Port"
	case 2053:
		description = "Lot105 DSuper Updates"
	case 2054:
		description = "Weblogin Port"
	case 2055:
		description = "Iliad-Odyssey Protocol"
	case 2056:
		description = "OmniSky Port"
	case 2057:
		description = "Rich Content Protocol"
	case 2058:
		description = "NewWaveSearchables RMI"
	case 2059:
		description = "BMC Messaging Service"
	case 2060:
		description = "Telenium Daemon IF"
	case 2061:
		description = "NetMount"
	case 2062:
		description = "ICG SWP Port"
	case 2063:
		description = "ICG Bridge Port"
	case 2064:
		description = "ICG IP Relay Port"
	case 2065:
		description = "Data Link Switch Read Port Number"
	case 2066:
		description = "AVM USB Remote Architecture"
	case 2067:
		description = "Data Link Switch Write Port Number"
	case 2068:
		description = "Avocent AuthSrv Protocol"
	case 2069:
		description = "HTTP Event Port"
	case 2070:
		description = "AH and ESP Encapsulated in UDP packet"
	case 2071:
		description = "Axon Control Protocol"
	case 2072:
		description = "GlobeCast mSync"
	case 2073:
		description = "DataReel Database Socket"
	case 2074:
		description = "Vertel VMF SA"
	case 2075:
		description = "Newlix ServerWare Engine"
	case 2076:
		description = "Newlix JSPConfig"
	case 2077:
		description = "Old Tivoli Storage Manager"
	case 2078:
		description = "IBM Total Productivity Center Server"
	case 2079:
		description = "IDWARE Router Port"
	case 2080:
		description = "Autodesk NLM (FLEXlm)"
	case 2081:
		description = "KME PRINTER TRAP PORT"
	case 2082:
		description = "Infowave Mobility Server"
	case 2083:
		description = "Secure Radius Service"
	case 2084:
		description = "SunCluster Geographic"
	case 2085:
		description = "ADA Control"
	case 2086:
		description = "GNUnet"
	case 2087:
		description = "ELI - Event Logging Integration"
	case 2088:
		description = "IP Busy Lamp Field"
	case 2089:
		description = "Security Encapsulation Protocol - SEP"
	case 2090:
		description = "Load Report Protocol"
	case 2091:
		description = "PRP"
	case 2092:
		description = "Descent 3"
	case 2093:
		description = "NBX CC"
	case 2094:
		description = "NBX AU"
	case 2095:
		description = "NBX SER"
	case 2096:
		description = "NBX DIR"
	case 2097:
		description = "Jet Form Preview"
	case 2098:
		description = "Dialog Port"
	case 2099:
		description = "H.225.0 Annex G Signalling"
	case 2100:
		description = "Amiga Network Filesystem"
	case 2101:
		description = "rtcm-sc104"
	case 2102:
		description = "Zephyr server"
	case 2103:
		description = "Zephyr serv-hm connection"
	case 2104:
		description = "Zephyr hostmanager"
	case 2105:
		description = "MiniPay"
	case 2106:
		description = "MZAP"
	case 2107:
		description = "BinTec Admin"
	case 2108:
		description = "Comcam"
	case 2109:
		description = "Ergolight"
	case 2110:
		description = "UMSP"
	case 2111:
		description = "OPNET Dynamic Sampling Agent Transaction Protocol"
	case 2112:
		description = "Idonix MetaNet"
	case 2113:
		description = "HSL StoRM"
	case 2114:
		description = "Classical Music Meta-Data Access and Enhancement"
	case 2115:
		description = "Key Distribution Manager"
	case 2116:
		description = "CCOWCMR"
	case 2117:
		description = "MENTACLIENT"
	case 2118:
		description = "MENTASERVER"
	case 2119:
		description = "GSIGATEKEEPER"
	case 2120:
		description = "Quick Eagle Networks CP"
	case 2121:
		description = "SCIENTIA-SSDB"
	case 2122:
		description = "CauPC Remote Control"
	case 2123:
		description = "GTP-Control Plane (3GPP)"
	case 2124:
		description = "ELATELINK"
	case 2125:
		description = "LOCKSTEP"
	case 2126:
		description = "PktCable-COPS"
	case 2127:
		description = "INDEX-PC-WB"
	case 2128:
		description = "Net Steward Control"
	case 2129:
		description = "cs-live.com"
	case 2130:
		description = "XDS"
	case 2131:
		description = "Avantageb2b"
	case 2132:
		description = "SoleraTec End Point Map"
	case 2133:
		description = "ZYMED-ZPP"
	case 2134:
		description = "AVENUE"
	case 2135:
		description = "Grid Resource Information Server"
	case 2136:
		description = "APPWORXSRV"
	case 2137:
		description = "CONNECT"
	case 2138:
		description = "UNBIND-CLUSTER"
	case 2139:
		description = "IAS-AUTH"
	case 2140:
		description = "IAS-REG"
	case 2141:
		description = "IAS-ADMIND"
	case 2142:
		description = "TDM OVER IP"
	case 2143:
		description = "Live Vault Job Control"
	case 2144:
		description = "Live Vault Fast Object Transfer"
	case 2145:
		description = "Live Vault Remote Diagnostic Console Support"
	case 2146:
		description = "Live Vault Admin Event Notification"
	case 2147:
		description = "Live Vault Authentication"
	case 2148:
		description = "VERITAS UNIVERSAL COMMUNICATION LAYER"
	case 2149:
		description = "ACPTSYS"
	case 2150:
		description = "DYNAMIC3D"
	case 2151:
		description = "DOCENT"
	case 2152:
		description = "GTP-User Plane (3GPP)"
	case 2153:
		description = "Control Protocol"
	case 2154:
		description = "Standard Protocol"
	case 2155:
		description = "Bridge Protocol"
	case 2156:
		description = "Talari Reliable Protocol"
	case 2157:
		description = "Xerox Network Document Scan Protocol"
	case 2158:
		description = "TouchNetPlus Service"
	case 2159:
		description = "GDB Remote Debug Port"
	case 2160:
		description = "APC 2160"
	case 2161:
		description = "APC 2161"
	case 2162:
		description = "Navisphere"
	case 2163:
		description = "Navisphere Secure"
	case 2164:
		description = "Dynamic DNS Version 3"
	case 2165:
		description = "X-Bone API"
	case 2166:
		description = "iwserver"
	case 2167:
		description = "Raw Async Serial Link"
	case 2168:
		description = "easy-soft Multiplexer"
	case 2169:
		description = "Backbone for Academic Information Notification (BRAIN)"
	case 2170:
		description = "EyeTV Server Port"
	case 2171:
		description = "MS Firewall Storage"
	case 2172:
		description = "MS Firewall SecureStorage"
	case 2173:
		description = "MS Firewall Replication"
	case 2174:
		description = "MS Firewall Intra Array"
	case 2175:
		description = "Microsoft Desktop AirSync Protocol"
	case 2176:
		description = "Microsoft ActiveSync Remote API"
	case 2177:
		description = "qWAVE Bandwidth Estimate"
	case 2178:
		description = "Peer Services for BITS"
	case 2179:
		description = "Microsoft RDP for virtual machines"
	case 2180:
		description = "Millicent Vendor Gateway Server"
	case 2181:
		description = "eforward"
	case 2182:
		description = "CGN status"
	case 2183:
		description = "Code Green configuration"
	case 2184:
		description = "NVD User"
	case 2185:
		description = "OnBase Distributed Disk Services"
	case 2186:
		description = "Guy-Tek Automated Update Applications"
	case 2187:
		description = "Sepehr System Management Control"
	case 2188:
		description = "Radware Resource Pool Manager"
	case 2189:
		description = "Secure Radware Resource Pool Manager"
	case 2190:
		description = "TiVoConnect Beacon"
	case 2191:
		description = "TvBus Messaging"
	case 2192:
		description = "ASDIS software management"
	case 2193:
		description = "Dr.Web Enterprise Management Service"
	case 2197:
		description = "MNP data exchange"
	case 2198:
		description = "OneHome Remote Access"
	case 2199:
		description = "OneHome Service Port"
	case 2201:
		description = "Advanced Training System Program"
	case 2202:
		description = "Int. Multimedia Teleconferencing Cosortium"
	case 2203:
		description = "b2 Runtime Protocol"
	case 2204:
		description = "b2 License Server"
	case 2205:
		description = "Java Presentation Server"
	case 2206:
		description = "HP OpenCall bus"
	case 2207:
		description = "HP Status and Services"
	case 2208:
		description = "HP I/O Backend"
	case 2209:
		description = "HP RIM for Files Portal Service"
	case 2210:
		description = "NOAAPORT Broadcast Network"
	case 2211:
		description = "EMWIN"
	case 2212:
		description = "LeeCO POS Server Service"
	case 2213:
		description = "Kali"
	case 2214:
		description = "RDQ Protocol Interface"
	case 2215:
		description = "IPCore.co.za GPRS"
	case 2216:
		description = "VTU data service"
	case 2217:
		description = "GoToDevice Device Management"
	case 2218:
		description = "Bounzza IRC Proxy"
	case 2219:
		description = "NetIQ NCAP Protocol"
	case 2220:
		description = "NetIQ End2End"
	case 2221:
		description = "EtherNet/IP over TLS"
	case 2222:
		description = "EtherNet/IP I/O"
	case 2223:
		description = "Rockwell CSP2"
	case 2224:
		description = "Easy Flexible Internet/Multiplayer Games"
	case 2225:
		description = "Resource Connection Initiation Protocol"
	case 2226:
		description = "Digital Instinct DRM"
	case 2227:
		description = "DI Messaging Service"
	case 2228:
		description = "eHome Message Server"
	case 2229:
		description = "DataLens Service"
	case 2230:
		description = "MetaSoft Job Queue Administration Service"
	case 2231:
		description = "WiMAX ASN Control Plane Protocol"
	case 2232:
		description = "IVS Video default"
	case 2233:
		description = "INFOCRYPT"
	case 2234:
		description = "DirectPlay"
	case 2235:
		description = "Sercomm-WLink"
	case 2236:
		description = "Nani"
	case 2237:
		description = "Optech Port1 License Manager"
	case 2238:
		description = "AVIVA SNA SERVER"
	case 2239:
		description = "Image Query"
	case 2240:
		description = "RECIPe"
	case 2241:
		description = "IVS Daemon"
	case 2242:
		description = "Folio Remote Server"
	case 2243:
		description = "Magicom Protocol"
	case 2244:
		description = "NMS Server"
	case 2245:
		description = "HaO"
	case 2246:
		description = "PacketCable MTA Addr Map"
	case 2247:
		description = "Antidote Deployment Manager Service"
	case 2248:
		description = "User Management Service"
	case 2249:
		description = "RISO File Manager Protocol"
	case 2250:
		description = "remote-collab"
	case 2251:
		description = "Distributed Framework Port"
	case 2252:
		description = "NJENET using SSL"
	case 2253:
		description = "DTV Channel Request"
	case 2254:
		description = "Seismic P.O.C. Port"
	case 2255:
		description = "VRTP - ViRtue Transfer Protocol"
	case 2256:
		description = "PCC MFP"
	case 2257:
		description = "simple text/file transfer"
	case 2258:
		description = "Rotorcraft Communications Test System"
	case 2259:
		description = "BIF identifiers resolution service"
	case 2260:
		description = "APC 2260"
	case 2261:
		description = "CoMotion Master Server"
	case 2262:
		description = "CoMotion Backup Server"
	case 2263:
		description = "ECweb Configuration Service"
	case 2264:
		description = "Audio Precision Apx500 API Port 1"
	case 2265:
		description = "Audio Precision Apx500 API Port 2"
	case 2266:
		description = "M-files Server"
	case 2267:
		description = "OntoBroker"
	case 2268:
		description = "AMT"
	case 2269:
		description = "MIKEY"
	case 2270:
		description = "starSchool"
	case 2271:
		description = "Secure Meeting Maker Scheduling"
	case 2272:
		description = "Meeting Maker Scheduling"
	case 2273:
		description = "MySQL Instance Manager"
	case 2274:
		description = "PCTTunneller"
	case 2275:
		description = "iBridge Conferencing"
	case 2276:
		description = "iBridge Management"
	case 2277:
		description = "Bt device control proxy"
	case 2278:
		description = "Simple Stacked Sequences Database"
	case 2279:
		description = "xmquery"
	case 2280:
		description = "LNVPOLLER"
	case 2281:
		description = "LNVCONSOLE"
	case 2282:
		description = "LNVALARM"
	case 2283:
		description = "LNVSTATUS"
	case 2284:
		description = "LNVMAPS"
	case 2285:
		description = "LNVMAILMON"
	case 2286:
		description = "NAS-Metering"
	case 2287:
		description = "DNA"
	case 2288:
		description = "NETML"
	case 2289:
		description = "Lookup dict server"
	case 2290:
		description = "Sonus Logging Services"
	case 2291:
		description = "EPSON Advanced Printer Share Protocol"
	case 2292:
		description = "Sonus Element Management Services"
	case 2293:
		description = "Network Platform Debug Manager"
	case 2294:
		description = "Konshus License Manager (FLEX)"
	case 2295:
		description = "Advant License Manager"
	case 2296:
		description = "Theta License Manager (Rainbow)"
	case 2297:
		description = "D2K DataMover 1"
	case 2298:
		description = "D2K DataMover 2"
	case 2299:
		description = "PC Telecommute"
	case 2300:
		description = "CVMMON"
	case 2301:
		description = "Compaq HTTP"
	case 2302:
		description = "Bindery Support"
	case 2303:
		description = "Proxy Gateway"
	case 2304:
		description = "Attachmate UTS"
	case 2305:
		description = "MT ScaleServer"
	case 2306:
		description = "TAPPI BoxNet"
	case 2307:
		description = "pehelp"
	case 2308:
		description = "sdhelp"
	case 2309:
		description = "SD Server"
	case 2310:
		description = "SD Client"
	case 2311:
		description = "Message Service"
	case 2312:
		description = "WANScaler Communication Service"
	case 2313:
		description = "IAPP (Inter Access Point Protocol)"
	case 2314:
		description = "CR WebSystems"
	case 2315:
		description = "Precise Sft."
	case 2316:
		description = "SENT License Manager"
	case 2317:
		description = "Attachmate G32"
	case 2318:
		description = "Cadence Control"
	case 2319:
		description = "InfoLibria"
	case 2320:
		description = "Siebel NS"
	case 2321:
		description = "RDLAP"
	case 2322:
		description = "ofsd"
	case 2323:
		description = "3d-nfsd"
	case 2324:
		description = "Cosmocall"
	case 2325:
		description = "ANSYS Licensing Interconnect"
	case 2326:
		description = "IDCP"
	case 2327:
		description = "xingcsm"
	case 2328:
		description = "Netrix SFTM"
	case 2329:
		description = "NVD"
	case 2330:
		description = "TSCCHAT"
	case 2331:
		description = "AGENTVIEW"
	case 2332:
		description = "RCC Host"
	case 2333:
		description = "SNAPP"
	case 2334:
		description = "ACE Client Auth"
	case 2335:
		description = "ACE Proxy"
	case 2336:
		description = "Apple UG Control"
	case 2337:
		description = "ideesrv"
	case 2338:
		description = "Norton Lambert"
	case 2339:
		description = "3Com WebView"
	case 2340:
		description = "WRS Registry"
	case 2341:
		description = "XIO Status"
	case 2342:
		description = "Seagate Manage Exec"
	case 2343:
		description = "nati logos"
	case 2344:
		description = "fcmsys"
	case 2345:
		description = "dbm"
	case 2346:
		description = "Game Connection Port"
	case 2347:
		description = "Game Announcement and Location"
	case 2348:
		description = "Information to query for game status"
	case 2349:
		description = "Diagnostics Port"
	case 2350:
		description = "Pharos Booking Server"
	case 2351:
		description = "psrserver"
	case 2352:
		description = "pslserver"
	case 2353:
		description = "pspserver"
	case 2354:
		description = "psprserver"
	case 2355:
		description = "psdbserver"
	case 2356:
		description = "GXT License Managemant"
	case 2357:
		description = "UniHub Server"
	case 2358:
		description = "Futrix"
	case 2359:
		description = "FlukeServer"
	case 2360:
		description = "NexstorIndLtd"
	case 2361:
		description = "TL1"
	case 2362:
		description = "digiman"
	case 2363:
		description = "Media Central NFSD"
	case 2364:
		description = "OI-2000"
	case 2365:
		description = "dbref"
	case 2366:
		description = "qip-login"
	case 2367:
		description = "Service Control"
	case 2368:
		description = "OpenTable"
	case 2369:
		description = "Blockchain Identifier InFrastructure P2P"
	case 2370:
		description = "L3-HBMon"
	case 2371:
		description = "Remote Device Access"
	case 2372:
		description = "LanMessenger"
	case 2373:
		description = "Remograph License Manager"
	case 2374:
		description = "Hydra RPC"
	case 2375:
		description = "Docker REST API (plain text)"
	case 2376:
		description = "Docker REST API (ssl)"
	case 2377:
		description = "RPC interface for Docker Swarm"
	case 2378:
		description = "DALI lighting control"
	case 2379:
		description = "etcd client communication"
	case 2380:
		description = "etcd server to server communication"
	case 2381:
		description = "Compaq HTTPS"
	case 2382:
		description = "Microsoft OLAP"
	case 2383:
		description = "Microsoft OLAP"
	case 2384:
		description = "SD-REQUEST"
	case 2385:
		description = "SD-DATA"
	case 2386:
		description = "Virtual Tape"
	case 2387:
		description = "VSAM Redirector"
	case 2388:
		description = "MYNAH AutoStart"
	case 2389:
		description = "OpenView Session Mgr"
	case 2390:
		description = "RSMTP"
	case 2391:
		description = "3COM Net Management"
	case 2392:
		description = "Tactical Auth"
	case 2393:
		description = "MS OLAP 1"
	case 2394:
		description = "MS OLAP 2"
	case 2395:
		description = "LAN900 Remote"
	case 2396:
		description = "Wusage"
	case 2397:
		description = "NCL"
	case 2398:
		description = "Orbiter"
	case 2399:
		description = "FileMaker"
	case 2400:
		description = "OpEquus Server"
	case 2401:
		description = "cvspserver"
	case 2402:
		description = "TaskMaster 2000 Server"
	case 2403:
		description = "TaskMaster 2000 Web"
	case 2404:
		description = "IEC 60870-5-104 process control over IP"
	case 2405:
		description = "TRC Netpoll"
	case 2406:
		description = "JediServer"
	case 2407:
		description = "Orion"
	case 2408:
		description = "CloudFlare Railgun Web Acceleration Protocol"
	case 2409:
		description = "SNS Protocol"
	case 2410:
		description = "VRTS Registry"
	case 2411:
		description = "Netwave AP Management"
	case 2412:
		description = "CDN"
	case 2413:
		description = "orion-rmi-reg"
	case 2414:
		description = "Beeyond"
	case 2415:
		description = "Codima Remote Transaction Protocol"
	case 2416:
		description = "RMT Server"
	case 2417:
		description = "Composit Server"
	case 2418:
		description = "cas"
	case 2419:
		description = "Attachmate S2S"
	case 2420:
		description = "DSL Remote Management"
	case 2421:
		description = "G-Talk"
	case 2422:
		description = "CRMSBITS"
	case 2423:
		description = "RNRP"
	case 2424:
		description = "KOFAX-SVR"
	case 2425:
		description = "Fujitsu App Manager"
	case 2426:
		description = "VeloCloud MultiPath Protocol"
	case 2427:
		description = "Media Gateway Control Protocol Gateway"
	case 2428:
		description = "One Way Trip Time"
	case 2429:
		description = "FT-ROLE"
	case 2430:
		description = "venus"
	case 2431:
		description = "venus-se"
	case 2432:
		description = "codasrv"
	case 2433:
		description = "codasrv-se"
	case 2434:
		description = "pxc-epmap"
	case 2435:
		description = "OptiLogic"
	case 2436:
		description = "TOP/X"
	case 2437:
		description = "UniControl"
	case 2438:
		description = "MSP"
	case 2439:
		description = "SybaseDBSynch"
	case 2440:
		description = "Spearway Lockers"
	case 2441:
		description = "Pervasive I*net Data Server"
	case 2442:
		description = "Netangel"
	case 2443:
		description = "PowerClient Central Storage Facility"
	case 2444:
		description = "BT PP2 Sectrans"
	case 2445:
		description = "DTN1"
	case 2446:
		description = "bues_service"
	case 2447:
		description = "OpenView NNM daemon"
	case 2448:
		description = "hpppsvr"
	case 2449:
		description = "RATL"
	case 2450:
		description = "netadmin"
	case 2451:
		description = "netchat"
	case 2452:
		description = "SnifferClient"
	case 2453:
		description = "madge ltd"
	case 2454:
		description = "IndX-DDS"
	case 2455:
		description = "WAGO-IO-SYSTEM"
	case 2456:
		description = "altav-remmgt"
	case 2457:
		description = "Rapido_IP"
	case 2458:
		description = "griffin"
	case 2459:
		description = "Community"
	case 2460:
		description = "ms-theater"
	case 2461:
		description = "qadmifoper"
	case 2462:
		description = "qadmifevent"
	case 2463:
		description = "LSI RAID Management"
	case 2464:
		description = "DirecPC SI"
	case 2465:
		description = "Load Balance Management"
	case 2466:
		description = "Load Balance Forwarding"
	case 2467:
		description = "High Criteria"
	case 2468:
		description = "qip_msgd"
	case 2469:
		description = "MTI-TCS-COMM"
	case 2470:
		description = "taskman port"
	case 2471:
		description = "SeaODBC"
	case 2472:
		description = "C3"
	case 2473:
		description = "Aker-cdp"
	case 2474:
		description = "Vital Analysis"
	case 2475:
		description = "ACE Server"
	case 2476:
		description = "ACE Server Propagation"
	case 2477:
		description = "SecurSight Certificate Valifation Service"
	case 2478:
		description = "SecurSight Authentication Server (SSL)"
	case 2479:
		description = "SecurSight Event Logging Server (SSL)"
	case 2480:
		description = "Informatica PowerExchange Listener"
	case 2481:
		description = "Oracle GIOP"
	case 2482:
		description = "Oracle GIOP SSL"
	case 2483:
		description = "Oracle TTC"
	case 2484:
		description = "Oracle TTC SSL"
	case 2485:
		description = "Net Objects1"
	case 2486:
		description = "Net Objects2"
	case 2487:
		description = "Policy Notice Service"
	case 2488:
		description = "Moy Corporation"
	case 2489:
		description = "TSILB"
	case 2490:
		description = "qip_qdhcp"
	case 2491:
		description = "Conclave CPP"
	case 2492:
		description = "GROOVE"
	case 2493:
		description = "Talarian MQS"
	case 2494:
		description = "BMC AR"
	case 2495:
		description = "Fast Remote Services"
	case 2496:
		description = "DIRGIS"
	case 2497:
		description = "Quad DB"
	case 2498:
		description = "ODN-CasTraq"
	case 2499:
		description = "UniControl"
	case 2500:
		description = "Resource Tracking system server"
	case 2501:
		description = "Resource Tracking system client"
	case 2502:
		description = "Kentrox Protocol"
	case 2503:
		description = "NMS-DPNSS"
	case 2504:
		description = "WLBS"
	case 2505:
		description = "PowerPlay Control"
	case 2506:
		description = "jbroker"
	case 2507:
		description = "spock"
	case 2508:
		description = "JDataStore"
	case 2509:
		description = "fjmpss"
	case 2510:
		description = "fjappmgrbulk"
	case 2511:
		description = "Metastorm"
	case 2512:
		description = "Citrix IMA"
	case 2513:
		description = "Citrix ADMIN"
	case 2514:
		description = "Facsys NTP"
	case 2515:
		description = "Facsys Router"
	case 2516:
		description = "Main Control"
	case 2517:
		description = "H.323 Annex E Call Control Signalling Transport"
	case 2518:
		description = "Willy"
	case 2519:
		description = "globmsgsvc"
	case 2520:
		description = "Pervasive Listener"
	case 2521:
		description = "Adaptec Manager"
	case 2522:
		description = "WinDb"
	case 2523:
		description = "Qke LLC V.3"
	case 2524:
		description = "Optiwave License Management"
	case 2525:
		description = "MS V-Worlds"
	case 2526:
		description = "EMA License Manager"
	case 2527:
		description = "IQ Server"
	case 2528:
		description = "NCR CCL"
	case 2529:
		description = "UTS FTP"
	case 2530:
		description = "VR Commerce"
	case 2531:
		description = "ITO-E GUI"
	case 2532:
		description = "OVTOPMD"
	case 2533:
		description = "SnifferServer"
	case 2534:
		description = "Combox Web Access"
	case 2535:
		description = "MADCAP"
	case 2536:
		description = "btpp2audctr1"
	case 2537:
		description = "Upgrade Protocol"
	case 2538:
		description = "vnwk-prapi"
	case 2539:
		description = "VSI Admin"
	case 2540:
		description = "LonWorks"
	case 2541:
		description = "LonWorks2"
	case 2542:
		description = "uDraw(Graph)"
	case 2543:
		description = "REFTEK"
	case 2544:
		description = "Management Daemon Refresh"
	case 2545:
		description = "sis-emt"
	case 2546:
		description = "vytalvaultbrtp"
	case 2547:
		description = "vytalvaultvsmp"
	case 2548:
		description = "vytalvaultpipe"
	case 2549:
		description = "IPASS"
	case 2550:
		description = "ADS"
	case 2551:
		description = "ISG UDA Server"
	case 2552:
		description = "Call Logging"
	case 2553:
		description = "efidiningport"
	case 2554:
		description = "VCnet-Link v10"
	case 2555:
		description = "Compaq WCP"
	case 2556:
		description = "nicetec-nmsvc"
	case 2557:
		description = "nicetec-mgmt"
	case 2558:
		description = "PCLE Multi Media"
	case 2559:
		description = "LSTP"
	case 2560:
		description = "labrat"
	case 2561:
		description = "MosaixCC"
	case 2562:
		description = "Delibo"
	case 2563:
		description = "CTI Redwood"
	case 2564:
		description = "HP 3000 NS/VT block mode telnet"
	case 2565:
		description = "Coordinator Server"
	case 2566:
		description = "pcs-pcw"
	case 2567:
		description = "Cisco Line Protocol"
	case 2568:
		description = "SPAM TRAP"
	case 2569:
		description = "Sonus Call Signal"
	case 2570:
		description = "HS Port"
	case 2571:
		description = "CECSVC"
	case 2572:
		description = "IBP"
	case 2573:
		description = "Trust Establish"
	case 2574:
		description = "Blockade BPSP"
	case 2575:
		description = "HL7"
	case 2576:
		description = "TCL Pro Debugger"
	case 2577:
		description = "Scriptics Lsrvr"
	case 2578:
		description = "RVS ISDN DCP"
	case 2579:
		description = "mpfoncl"
	case 2580:
		description = "Tributary"
	case 2581:
		description = "ARGIS TE"
	case 2582:
		description = "ARGIS DS"
	case 2583:
		description = "MON"
	case 2584:
		description = "cyaserv"
	case 2585:
		description = "NETX Server"
	case 2586:
		description = "NETX Agent"
	case 2587:
		description = "MASC"
	case 2588:
		description = "Privilege"
	case 2589:
		description = "quartus tcl"
	case 2590:
		description = "idotdist"
	case 2591:
		description = "Maytag Shuffle"
	case 2592:
		description = "netrek"
	case 2593:
		description = "MNS Mail Notice Service"
	case 2594:
		description = "Data Base Server"
	case 2595:
		description = "World Fusion 1"
	case 2596:
		description = "World Fusion 2"
	case 2597:
		description = "Homestead Glory"
	case 2598:
		description = "Citrix MA Client"
	case 2599:
		description = "Snap Discovery"
	case 2600:
		description = "HPSTGMGR"
	case 2601:
		description = "discp client"
	case 2602:
		description = "discp server"
	case 2603:
		description = "Service Meter"
	case 2604:
		description = "NSC CCS"
	case 2605:
		description = "NSC POSA"
	case 2606:
		description = "Dell Netmon"
	case 2607:
		description = "Dell Connection"
	case 2608:
		description = "Wag Service"
	case 2609:
		description = "System Monitor"
	case 2610:
		description = "VersaTek"
	case 2611:
		description = "LIONHEAD"
	case 2612:
		description = "Qpasa Agent"
	case 2613:
		description = "SMNTUBootstrap"
	case 2614:
		description = "Never Offline"
	case 2615:
		description = "firepower"
	case 2616:
		description = "appswitch-emp"
	case 2617:
		description = "Clinical Context Managers"
	case 2618:
		description = "Priority E-Com"
	case 2619:
		description = "bruce"
	case 2620:
		description = "LPSRecommender"
	case 2621:
		description = "Miles Apart Jukebox Server"
	case 2622:
		description = "MetricaDBC"
	case 2623:
		description = "LMDP"
	case 2624:
		description = "Aria"
	case 2625:
		description = "Blwnkl Port"
	case 2626:
		description = "gbjd816"
	case 2627:
		description = "Moshe Beeri"
	case 2628:
		description = "DICT"
	case 2629:
		description = "Sitara Server"
	case 2630:
		description = "Sitara Management"
	case 2631:
		description = "Sitara Dir"
	case 2632:
		description = "IRdg Post"
	case 2633:
		description = "InterIntelli"
	case 2634:
		description = "PK Electronics"
	case 2635:
		description = "Back Burner"
	case 2636:
		description = "Solve"
	case 2637:
		description = "Import Document Service"
	case 2638:
		description = "Sybase Anywhere"
	case 2639:
		description = "AMInet"
	case 2640:
		description = "Alcorn McBride Inc protocol used for device control"
	case 2641:
		description = "HDL Server"
	case 2642:
		description = "Tragic"
	case 2643:
		description = "GTE-SAMP"
	case 2644:
		description = "Travsoft IPX Tunnel"
	case 2645:
		description = "Novell IPX CMD"
	case 2646:
		description = "AND License Manager"
	case 2647:
		description = "SyncServer"
	case 2648:
		description = "Upsnotifyprot"
	case 2649:
		description = "VPSIPPORT"
	case 2650:
		description = "eristwoguns"
	case 2651:
		description = "EBInSite"
	case 2652:
		description = "InterPathPanel"
	case 2653:
		description = "Sonus"
	case 2654:
		description = "Corel VNC Admin"
	case 2655:
		description = "UNIX Nt Glue"
	case 2656:
		description = "Kana"
	case 2657:
		description = "SNS Dispatcher"
	case 2658:
		description = "SNS Admin"
	case 2659:
		description = "SNS Query"
	case 2660:
		description = "GC Monitor"
	case 2661:
		description = "OLHOST"
	case 2662:
		description = "BinTec-CAPI"
	case 2663:
		description = "BinTec-TAPI"
	case 2664:
		description = "Patrol for MQ GM"
	case 2665:
		description = "Patrol for MQ NM"
	case 2666:
		description = "extensis"
	case 2667:
		description = "Alarm Clock Server"
	case 2668:
		description = "Alarm Clock Client"
	case 2669:
		description = "TOAD"
	case 2670:
		description = "TVE Announce"
	case 2671:
		description = "newlixreg"
	case 2672:
		description = "nhserver"
	case 2673:
		description = "First Call 42"
	case 2674:
		description = "ewnn"
	case 2675:
		description = "TTC ETAP"
	case 2676:
		description = "SIMSLink"
	case 2677:
		description = "Gadget Gate 1 Way"
	case 2678:
		description = "Gadget Gate 2 Way"
	case 2679:
		description = "Sync Server SSL"
	case 2680:
		description = "pxc-sapxom"
	case 2681:
		description = "mpnjsomb"
	case 2682:
		description = "Removed"
	case 2683:
		description = "NCDLoadBalance"
	case 2684:
		description = "mpnjsosv"
	case 2685:
		description = "mpnjsocl"
	case 2686:
		description = "mpnjsomg"
	case 2687:
		description = "pq-lic-mgmt"
	case 2688:
		description = "md-cf-http"
	case 2689:
		description = "FastLynx"
	case 2690:
		description = "HP NNM Embedded Database"
	case 2691:
		description = "ITInternet ISM Server"
	case 2692:
		description = "Admins LMS"
	case 2694:
		description = "pwrsevent"
	case 2695:
		description = "VSPREAD"
	case 2696:
		description = "Unify Admin"
	case 2697:
		description = "Oce SNMP Trap Port"
	case 2698:
		description = "MCK-IVPIP"
	case 2699:
		description = "Csoft Plus Client"
	case 2700:
		description = "tqdata"
	case 2701:
		description = "SMS RCINFO"
	case 2702:
		description = "SMS XFER"
	case 2703:
		description = "SMS CHAT"
	case 2704:
		description = "SMS REMCTRL"
	case 2705:
		description = "SDS Admin"
	case 2706:
		description = "NCD Mirroring"
	case 2707:
		description = "EMCSYMAPIPORT"
	case 2708:
		description = "Banyan-Net"
	case 2709:
		description = "Supermon"
	case 2710:
		description = "SSO Service"
	case 2711:
		description = "SSO Control"
	case 2712:
		description = "Axapta Object Communication Protocol"
	case 2713:
		description = "Raven Trinity Broker Service"
	case 2714:
		description = "Raven Trinity Data Mover"
	case 2715:
		description = "HPSTGMGR2"
	case 2716:
		description = "Inova IP Disco"
	case 2717:
		description = "PN REQUESTER"
	case 2718:
		description = "PN REQUESTER 2"
	case 2719:
		description = "Scan & Change"
	case 2720:
		description = "wkars"
	case 2721:
		description = "Smart Diagnose"
	case 2722:
		description = "Proactive Server"
	case 2723:
		description = "WatchDog NT Protocol"
	case 2724:
		description = "qotps"
	case 2725:
		description = "MSOLAP PTP2"
	case 2726:
		description = "TAMS"
	case 2727:
		description = "Media Gateway Control Protocol Call Agent"
	case 2728:
		description = "SQDR"
	case 2729:
		description = "TCIM Control"
	case 2730:
		description = "NEC RaidPlus"
	case 2731:
		description = "Fyre Messanger"
	case 2732:
		description = "G5M"
	case 2733:
		description = "Signet CTF"
	case 2734:
		description = "CCS Software"
	case 2735:
		description = "NetIQ Monitor Console"
	case 2736:
		description = "RADWIZ NMS SRV"
	case 2737:
		description = "SRP Feedback"
	case 2738:
		description = "NDL TCP-OSI Gateway"
	case 2739:
		description = "TN Timing"
	case 2740:
		description = "Alarm"
	case 2741:
		description = "TSB"
	case 2742:
		description = "TSB2"
	case 2743:
		description = "murx"
	case 2744:
		description = "honyaku"
	case 2745:
		description = "URBISNET"
	case 2746:
		description = "CPUDPENCAP"
	case 2752:
		description = "RSISYS ACCESS"
	case 2753:
		description = "de-spot"
	case 2754:
		description = "APOLLO CC"
	case 2755:
		description = "Express Pay"
	case 2756:
		description = "simplement-tie"
	case 2757:
		description = "CNRP"
	case 2758:
		description = "APOLLO Status"
	case 2759:
		description = "APOLLO GMS"
	case 2760:
		description = "Saba MS"
	case 2761:
		description = "DICOM ISCL"
	case 2762:
		description = "DICOM TLS"
	case 2763:
		description = "Desktop DNA"
	case 2764:
		description = "Data Insurance"
	case 2765:
		description = "qip-audup"
	case 2766:
		description = "Compaq SCP"
	case 2767:
		description = "UADTC"
	case 2768:
		description = "UACS"
	case 2769:
		description = "eXcE"
	case 2770:
		description = "Veronica"
	case 2771:
		description = "Vergence CM"
	case 2772:
		description = "auris"
	case 2773:
		description = "RBackup Remote Backup"
	case 2774:
		description = "RBackup Remote Backup"
	case 2775:
		description = "SMPP"
	case 2776:
		description = "Ridgeway Systems & Software"
	case 2777:
		description = "Ridgeway Systems & Software"
	case 2778:
		description = "Gwen-Sonya"
	case 2779:
		description = "LBC Sync"
	case 2780:
		description = "LBC Control"
	case 2781:
		description = "whosells"
	case 2782:
		description = "everydayrc"
	case 2783:
		description = "AISES"
	case 2784:
		description = "world wide web - development"
	case 2785:
		description = "aic-np"
	case 2786:
		description = "aic-oncrpc - Destiny MCD database"
	case 2787:
		description = "piccolo - Cornerstone Software"
	case 2788:
		description = "NetWare Loadable Module - Seagate Software"
	case 2789:
		description = "Media Agent"
	case 2790:
		description = "PLG Proxy"
	case 2791:
		description = "MT Port Registrator"
	case 2792:
		description = "f5-globalsite"
	case 2793:
		description = "initlsmsad"
	case 2794:
		description = "Uniform Resource Platform"
	case 2795:
		description = "LiveStats"
	case 2796:
		description = "ac-tech"
	case 2797:
		description = "esp-encap"
	case 2798:
		description = "TMESIS-UPShot"
	case 2799:
		description = "ICON Discover"
	case 2800:
		description = "ACC RAID"
	case 2801:
		description = "IGCP"
	case 2802:
		description = "Veritas"
	case 2803:
		description = "btprjctrl"
	case 2804:
		description = "March Networks Digital Video Recorders"
	case 2805:
		description = "WTA WSP-S"
	case 2806:
		description = "cspuni"
	case 2807:
		description = "cspmulti"
	case 2808:
		description = "J-LAN-P"
	case 2809:
		description = "CORBA LOC"
	case 2810:
		description = "Active Net Steward"
	case 2811:
		description = "GSI FTP"
	case 2812:
		description = "atmtcp"
	case 2813:
		description = "llm-pass"
	case 2814:
		description = "llm-csv"
	case 2815:
		description = "LBC Measurement"
	case 2816:
		description = "LBC Watchdog"
	case 2817:
		description = "NMSig Port"
	case 2818:
		description = "rmlnk"
	case 2819:
		description = "FC Fault Notification"
	case 2820:
		description = "UniVision"
	case 2821:
		description = "VERITAS Authentication Service"
	case 2822:
		description = "ka0wuc"
	case 2823:
		description = "CQG Net/LAN"
	case 2824:
		description = "CQG Net/LAN 1"
	case 2826:
		description = "slc systemlog"
	case 2827:
		description = "slc ctrlrloops"
	case 2828:
		description = "ITM License Manager"
	case 2829:
		description = "silkp1"
	case 2830:
		description = "silkp2"
	case 2831:
		description = "silkp3"
	case 2832:
		description = "silkp4"
	case 2833:
		description = "glishd"
	case 2834:
		description = "EVTP"
	case 2835:
		description = "EVTP-DATA"
	case 2836:
		description = "catalyst"
	case 2837:
		description = "Repliweb"
	case 2838:
		description = "Starbot"
	case 2839:
		description = "NMSigPort"
	case 2840:
		description = "l3-exprt"
	case 2841:
		description = "l3-ranger"
	case 2842:
		description = "l3-hawk"
	case 2843:
		description = "PDnet"
	case 2844:
		description = "BPCP POLL"
	case 2845:
		description = "BPCP TRAP"
	case 2846:
		description = "AIMPP Hello"
	case 2847:
		description = "AIMPP Port Req"
	case 2848:
		description = "AMT-BLC-PORT"
	case 2849:
		description = "FXP"
	case 2850:
		description = "MetaConsole"
	case 2851:
		description = "webemshttp"
	case 2852:
		description = "bears-01"
	case 2853:
		description = "ISPipes"
	case 2854:
		description = "InfoMover"
	case 2855:
		description = "MSRP over TCP"
	case 2856:
		description = "cesdinv"
	case 2857:
		description = "SimCtIP"
	case 2858:
		description = "ECNP"
	case 2859:
		description = "Active Memory"
	case 2860:
		description = "Dialpad Voice 1"
	case 2861:
		description = "Dialpad Voice 2"
	case 2862:
		description = "TTG Protocol"
	case 2863:
		description = "Sonar Data"
	case 2864:
		description = "main 5001 cmd"
	case 2865:
		description = "pit-vpn"
	case 2866:
		description = "iwlistener"
	case 2867:
		description = "esps-portal"
	case 2868:
		description = "Norman Proprietaqry Events Protocol"
	case 2869:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 2870:
		description = "daishi"
	case 2871:
		description = "MSI Select Play"
	case 2872:
		description = "RADIX"
	case 2873:
		description = "PubSub Realtime Telemetry Protocol"
	case 2874:
		description = "DX Message Base Transport Protocol"
	case 2875:
		description = "DX Message Base Transport Protocol"
	case 2876:
		description = "SPS Tunnel"
	case 2877:
		description = "BLUELANCE"
	case 2878:
		description = "AAP"
	case 2879:
		description = "ucentric-ds"
	case 2880:
		description = "Synapse Transport"
	case 2881:
		description = "NDSP"
	case 2882:
		description = "NDTP"
	case 2883:
		description = "NDNP"
	case 2884:
		description = "Flash Msg"
	case 2885:
		description = "TopFlow"
	case 2886:
		description = "RESPONSELOGIC"
	case 2887:
		description = "aironet"
	case 2888:
		description = "SPCSDLOBBY"
	case 2889:
		description = "RSOM"
	case 2890:
		description = "CSPCLMULTI"
	case 2891:
		description = "CINEGRFX-ELMD License Manager"
	case 2892:
		description = "SNIFFERDATA"
	case 2893:
		description = "VSECONNECTOR"
	case 2894:
		description = "ABACUS-REMOTE"
	case 2895:
		description = "NATUS LINK"
	case 2896:
		description = "ECOVISIONG6-1"
	case 2897:
		description = "Citrix RTMP"
	case 2898:
		description = "APPLIANCE-CFG"
	case 2899:
		description = "POWERGEMPLUS"
	case 2900:
		description = "QUICKSUITE"
	case 2901:
		description = "ALLSTORCNS"
	case 2902:
		description = "NET ASPI"
	case 2903:
		description = "SUITCASE"
	case 2904:
		description = "M2UA"
	case 2905:
		description = "M3UA"
	case 2906:
		description = "CALLER9"
	case 2907:
		description = "WEBMETHODS B2B"
	case 2908:
		description = "mao"
	case 2909:
		description = "Funk Dialout"
	case 2910:
		description = "TDAccess"
	case 2911:
		description = "Blockade"
	case 2912:
		description = "Epicon"
	case 2913:
		description = "Booster Ware"
	case 2914:
		description = "Game Lobby"
	case 2915:
		description = "TK Socket"
	case 2916:
		description = "Elvin Server"
	case 2917:
		description = "Elvin Client"
	case 2918:
		description = "Kasten Chase Pad"
	case 2919:
		description = "roboER"
	case 2920:
		description = "roboEDA"
	case 2921:
		description = "CESD Contents Delivery Management"
	case 2922:
		description = "CESD Contents Delivery Data Transfer"
	case 2923:
		description = "WTA-WSP-WTP-S"
	case 2924:
		description = "PRECISE-VIP"
	case 2926:
		description = "MOBILE-FILE-DL"
	case 2927:
		description = "UNIMOBILECTRL"
	case 2928:
		description = "REDSTONE-CPSS"
	case 2929:
		description = "AMX-WEBADMIN"
	case 2930:
		description = "AMX-WEBLINX"
	case 2931:
		description = "Circle-X"
	case 2932:
		description = "INCP"
	case 2933:
		description = "4-TIER OPM GW"
	case 2934:
		description = "4-TIER OPM CLI"
	case 2935:
		description = "QTP"
	case 2936:
		description = "OTPatch"
	case 2937:
		description = "PNACONSULT-LM"
	case 2938:
		description = "SM-PAS-1"
	case 2939:
		description = "SM-PAS-2"
	case 2940:
		description = "SM-PAS-3"
	case 2941:
		description = "SM-PAS-4"
	case 2942:
		description = "SM-PAS-5"
	case 2943:
		description = "TTNRepository"
	case 2944:
		description = "Megaco-H.248 text"
	case 2945:
		description = "Megaco/H.248 binary"
	case 2946:
		description = "FJSVmpor"
	case 2947:
		description = "GPS Daemon request/response protocol"
	case 2948:
		description = "WAP PUSH"
	case 2949:
		description = "WAP PUSH SECURE"
	case 2950:
		description = "ESIP"
	case 2951:
		description = "OTTP"
	case 2952:
		description = "MPFWSAS"
	case 2953:
		description = "OVALARMSRV"
	case 2954:
		description = "OVALARMSRV-CMD"
	case 2955:
		description = "CSNOTIFY"
	case 2956:
		description = "OVRIMOSDBMAN"
	case 2957:
		description = "JAMCT5"
	case 2958:
		description = "JAMCT6"
	case 2959:
		description = "RMOPAGT"
	case 2960:
		description = "DFOXSERVER"
	case 2961:
		description = "BOLDSOFT-LM"
	case 2962:
		description = "IPH-POLICY-CLI"
	case 2963:
		description = "IPH-POLICY-ADM"
	case 2964:
		description = "BULLANT SRAP"
	case 2965:
		description = "BULLANT RAP"
	case 2966:
		description = "IDP-INFOTRIEVE"
	case 2967:
		description = "SSC-AGENT"
	case 2968:
		description = "ENPP"
	case 2969:
		description = "ESSP"
	case 2970:
		description = "INDEX-NET"
	case 2971:
		description = "NetClip clipboard daemon"
	case 2972:
		description = "PMSM Webrctl"
	case 2973:
		description = "SV Networks"
	case 2974:
		description = "Signal"
	case 2975:
		description = "Fujitsu Configuration Management Service"
	case 2976:
		description = "CNS Server Port"
	case 2977:
		description = "TTCs Enterprise Test Access Protocol - NS"
	case 2978:
		description = "TTCs Enterprise Test Access Protocol - DS"
	case 2979:
		description = "H.263 Video Streaming"
	case 2980:
		description = "Instant Messaging Service"
	case 2981:
		description = "MYLXAMPORT"
	case 2982:
		description = "IWB-WHITEBOARD"
	case 2983:
		description = "NETPLAN"
	case 2984:
		description = "HPIDSADMIN"
	case 2985:
		description = "HPIDSAGENT"
	case 2986:
		description = "STONEFALLS"
	case 2987:
		description = "identify"
	case 2988:
		description = "HIPPA Reporting Protocol"
	case 2989:
		description = "ZARKOV Intelligent Agent Communication"
	case 2990:
		description = "BOSCAP"
	case 2991:
		description = "WKSTN-MON"
	case 2992:
		description = "Avenyo Server"
	case 2993:
		description = "VERITAS VIS1"
	case 2994:
		description = "VERITAS VIS2"
	case 2995:
		description = "IDRS"
	case 2996:
		description = "vsixml"
	case 2997:
		description = "REBOL"
	case 2998:
		description = "Real Secure"
	case 3000:
		description = "RemoteWare Client"
	case 3001:
		description = "OrigoDB Server Native Interface"
	case 3002:
		description = "RemoteWare Server"
	case 3003:
		description = "CGMS"
	case 3004:
		description = "Csoft Agent"
	case 3005:
		description = "Genius License Manager"
	case 3006:
		description = "Instant Internet Admin"
	case 3007:
		description = "Lotus Mail Tracking Agent Protocol"
	case 3008:
		description = "Midnight Technologies"
	case 3009:
		description = "PXC-NTFY"
	case 3010:
		description = "Telerate Workstation"
	case 3011:
		description = "Trusted Web"
	case 3012:
		description = "Trusted Web Client"
	case 3013:
		description = "Gilat Sky Surfer"
	case 3014:
		description = "Broker Service"
	case 3015:
		description = "NATI DSTP"
	case 3016:
		description = "Notify Server"
	case 3017:
		description = "Event Listener"
	case 3018:
		description = "Service Registry"
	case 3019:
		description = "Resource Manager"
	case 3020:
		description = "CIFS"
	case 3021:
		description = "AGRI Server"
	case 3022:
		description = "CSREGAGENT"
	case 3023:
		description = "magicnotes"
	case 3024:
		description = "NDS_SSO"
	case 3025:
		description = "Arepa Raft"
	case 3026:
		description = "AGRI Gateway"
	case 3027:
		description = "LiebDevMgmt_C"
	case 3028:
		description = "LiebDevMgmt_DM"
	case 3029:
		description = "LiebDevMgmt_A"
	case 3030:
		description = "Arepa Cas"
	case 3031:
		description = "Remote AppleEvents/PPC Toolbox"
	case 3032:
		description = "Redwood Chat"
	case 3033:
		description = "PDB"
	case 3034:
		description = "Osmosis / Helix (R) AEEA Port"
	case 3035:
		description = "FJSV gssagt"
	case 3036:
		description = "Hagel DUMP"
	case 3037:
		description = "HP SAN Mgmt"
	case 3038:
		description = "Santak UPS"
	case 3039:
		description = "Cogitate"
	case 3040:
		description = "Tomato Springs"
	case 3041:
		description = "di-traceware"
	case 3042:
		description = "journee"
	case 3043:
		description = "Broadcast Routing Protocol"
	case 3044:
		description = "EndPoint Protocol"
	case 3045:
		description = "ResponseNet"
	case 3046:
		description = "di-ase"
	case 3047:
		description = "Fast Security HL Server"
	case 3048:
		description = "Sierra Net PC Trader"
	case 3049:
		description = "NSWS"
	case 3050:
		description = "gds_db"
	case 3051:
		description = "Galaxy Server"
	case 3052:
		description = "APC 3052"
	case 3053:
		description = "dsom-server"
	case 3054:
		description = "AMT CNF PROT"
	case 3055:
		description = "Policy Server"
	case 3056:
		description = "CDL Server"
	case 3057:
		description = "GoAhead FldUp"
	case 3058:
		description = "videobeans"
	case 3059:
		description = "qsoft"
	case 3060:
		description = "interserver"
	case 3061:
		description = "cautcpd"
	case 3062:
		description = "ncacn-ip-tcp"
	case 3063:
		description = "ncadg-ip-udp"
	case 3064:
		description = "Remote Port Redirector"
	case 3065:
		description = "slinterbase"
	case 3066:
		description = "NETATTACHSDMP"
	case 3067:
		description = "FJHPJP"
	case 3068:
		description = "ls3 Broadcast"
	case 3069:
		description = "ls3"
	case 3070:
		description = "MGXSWITCH"
	case 3071:
		description = "Crossplatform replication protocol"
	case 3072:
		description = "ContinuStor Monitor Port"
	case 3073:
		description = "Very simple chatroom prot"
	case 3074:
		description = "Xbox game port"
	case 3075:
		description = "Orbix 2000 Locator"
	case 3076:
		description = "Orbix 2000 Config"
	case 3077:
		description = "Orbix 2000 Locator SSL"
	case 3078:
		description = "Orbix 2000 Locator SSL"
	case 3079:
		description = "LV Front Panel"
	case 3080:
		description = "stm_pproc"
	case 3081:
		description = "TL1-LV"
	case 3082:
		description = "TL1-RAW"
	case 3083:
		description = "TL1-TELNET"
	case 3084:
		description = "ITM-MCCS"
	case 3085:
		description = "PCIHReq"
	case 3086:
		description = "JDL-DBKitchen"
	case 3087:
		description = "Asoki SMA"
	case 3088:
		description = "eXtensible Data Transfer Protocol"
	case 3089:
		description = "ParaTek Agent Linking"
	case 3090:
		description = "Senforce Session Services"
	case 3091:
		description = "1Ci Server Management"
	case 3093:
		description = "Jiiva RapidMQ Center"
	case 3094:
		description = "Jiiva RapidMQ Registry"
	case 3095:
		description = "Panasas rendezvous port"
	case 3096:
		description = "Active Print Server Port"
	case 3097:
		description = "ITU-T Q.1902.1/Q.2150.3"
	case 3098:
		description = "Universal Message Manager"
	case 3099:
		description = "CHIPSY Machine Daemon"
	case 3100:
		description = "OpCon/xps"
	case 3101:
		description = "HP PolicyXpert PIB Server"
	case 3102:
		description = "SoftlinK Slave Mon Port"
	case 3103:
		description = "Autocue SMI Protocol"
	case 3104:
		description = "Autocue Logger Protocol"
	case 3105:
		description = "Cardbox"
	case 3106:
		description = "Cardbox HTTP"
	case 3107:
		description = "Business protocol"
	case 3108:
		description = "Geolocate protocol"
	case 3109:
		description = "Personnel protocol"
	case 3110:
		description = "simulator control port"
	case 3111:
		description = "Web Synchronous Services"
	case 3112:
		description = "KDE System Guard"
	case 3113:
		description = "CS-Authenticate Svr Port"
	case 3114:
		description = "CCM AutoDiscover"
	case 3115:
		description = "MCTET Master"
	case 3116:
		description = "MCTET Gateway"
	case 3117:
		description = "MCTET Jserv"
	case 3118:
		description = "PKAgent"
	case 3119:
		description = "D2000 Kernel Port"
	case 3120:
		description = "D2000 Webserver Port"
	case 3121:
		description = "pacemaker remote"
	case 3122:
		description = "MTI VTR Emulator port"
	case 3123:
		description = "EDI Translation Protocol"
	case 3124:
		description = "Beacon Port"
	case 3125:
		description = "A13-AN Interface"
	case 3127:
		description = "CTX Bridge Port"
	case 3128:
		description = "Active API Server Port"
	case 3129:
		description = "NetPort Discovery Port"
	case 3130:
		description = "ICPv2"
	case 3131:
		description = "Net Book Mark"
	case 3132:
		description = "Microsoft Business Rule Engine Update Service"
	case 3133:
		description = "Prism Deploy User Port"
	case 3134:
		description = "Extensible Code Protocol"
	case 3135:
		description = "PeerBook Port"
	case 3136:
		description = "Grub Server Port"
	case 3137:
		description = "rtnt-1 data packets"
	case 3138:
		description = "rtnt-2 data packets"
	case 3139:
		description = "Incognito Rendez-Vous"
	case 3140:
		description = "Arilia Multiplexor"
	case 3141:
		description = "VMODEM"
	case 3142:
		description = "RDC WH EOS"
	case 3143:
		description = "Sea View"
	case 3144:
		description = "Tarantella"
	case 3145:
		description = "CSI-LFAP"
	case 3146:
		description = "bears-02"
	case 3147:
		description = "RFIO"
	case 3148:
		description = "NetMike Game Administrator"
	case 3149:
		description = "NetMike Game Server"
	case 3150:
		description = "NetMike Assessor Administrator"
	case 3151:
		description = "NetMike Assessor"
	case 3152:
		description = "FeiTian Port"
	case 3153:
		description = "S8Cargo Client Port"
	case 3154:
		description = "ON RMI Registry"
	case 3155:
		description = "JpegMpeg Port"
	case 3156:
		description = "Indura Collector"
	case 3157:
		description = "LSA Communicator"
	case 3158:
		description = "SmashTV Protocol"
	case 3159:
		description = "NavegaWeb Tarification"
	case 3160:
		description = "TIP Application Server"
	case 3161:
		description = "DOC1 License Manager"
	case 3162:
		description = "SFLM"
	case 3163:
		description = "RES-SAP"
	case 3164:
		description = "IMPRS"
	case 3165:
		description = "Newgenpay Engine Service"
	case 3166:
		description = "Quest Spotlight Out-Of-Process Collector"
	case 3167:
		description = "Now Contact Public Server"
	case 3168:
		description = "Now Up-to-Date Public Server"
	case 3169:
		description = "SERVERVIEW-AS"
	case 3170:
		description = "SERVERVIEW-ASN"
	case 3171:
		description = "SERVERVIEW-GF"
	case 3172:
		description = "SERVERVIEW-RM"
	case 3173:
		description = "SERVERVIEW-ICC"
	case 3174:
		description = "ARMI Server"
	case 3175:
		description = "T1_E1_Over_IP"
	case 3176:
		description = "ARS Master"
	case 3177:
		description = "Phonex Protocol"
	case 3178:
		description = "Radiance UltraEdge Port"
	case 3179:
		description = "H2GF W.2m Handover prot."
	case 3180:
		description = "Millicent Broker Server"
	case 3181:
		description = "BMC Patrol Agent"
	case 3182:
		description = "BMC Patrol Rendezvous"
	case 3183:
		description = "COPS/TLS"
	case 3184:
		description = "ApogeeX Port"
	case 3185:
		description = "SuSE Meta PPPD"
	case 3186:
		description = "IIW Monitor User Port"
	case 3187:
		description = "Open Design Listen Port"
	case 3188:
		description = "Broadcom Port"
	case 3189:
		description = "Pinnacle Sys InfEx Port"
	case 3190:
		description = "ConServR Proxy"
	case 3191:
		description = "ConServR SSL Proxy"
	case 3192:
		description = "FireMon Revision Control"
	case 3193:
		description = "SpanDataPort"
	case 3194:
		description = "Rockstorm MAG protocol"
	case 3195:
		description = "Network Control Unit"
	case 3196:
		description = "Network Control Unit"
	case 3197:
		description = "Embrace Device Protocol Server"
	case 3198:
		description = "Embrace Device Protocol Client"
	case 3199:
		description = "DMOD WorkSpace"
	case 3200:
		description = "Press-sense Tick Port"
	case 3201:
		description = "CPQ-TaskSmart"
	case 3202:
		description = "IntraIntra"
	case 3203:
		description = "Network Watcher Monitor"
	case 3204:
		description = "Network Watcher DB Access"
	case 3205:
		description = "iSNS Server Port"
	case 3206:
		description = "IronMail POP Proxy"
	case 3207:
		description = "Veritas Authentication Port"
	case 3208:
		description = "PFU PR Callback"
	case 3209:
		description = "HP OpenView Network Path Engine Server"
	case 3210:
		description = "Flamenco Networks Proxy"
	case 3211:
		description = "Avocent Secure Management"
	case 3212:
		description = "Survey Instrument"
	case 3213:
		description = "NEON 24X7 Mission Control"
	case 3214:
		description = "JMQ Daemon Port 1"
	case 3215:
		description = "JMQ Daemon Port 2"
	case 3216:
		description = "Ferrari electronic FOAM"
	case 3217:
		description = "Unified IP & Telecom Environment"
	case 3218:
		description = "EMC SmartPackets"
	case 3219:
		description = "WMS Messenger"
	case 3220:
		description = "XML NM over SSL"
	case 3221:
		description = "XML NM over TCP"
	case 3222:
		description = "Gateway Load Balancing Pr"
	case 3223:
		description = "DIGIVOTE (R) Vote-Server"
	case 3224:
		description = "AES Discovery Port"
	case 3225:
		description = "FCIP"
	case 3226:
		description = "ISI Industry Software IRP"
	case 3227:
		description = "DiamondWave NMS Server"
	case 3228:
		description = "DiamondWave MSG Server"
	case 3229:
		description = "Global CD Port"
	case 3230:
		description = "Software Distributor Port"
	case 3231:
		description = "VidiGo communication (previous was: Delta Solutions Direct)"
	case 3232:
		description = "MDT port"
	case 3233:
		description = "WhiskerControl main port"
	case 3234:
		description = "Alchemy Server"
	case 3235:
		description = "MDAP port"
	case 3236:
		description = "appareNet Test Server"
	case 3237:
		description = "appareNet Test Packet Sequencer"
	case 3238:
		description = "appareNet Analysis Server"
	case 3239:
		description = "appareNet User Interface"
	case 3240:
		description = "Trio Motion Control Port"
	case 3241:
		description = "SysOrb Monitoring Server"
	case 3242:
		description = "Session Description ID"
	case 3243:
		description = "Timelot Port"
	case 3244:
		description = "OneSAF"
	case 3245:
		description = "VIEO Fabric Executive"
	case 3246:
		description = "DVT SYSTEM PORT"
	case 3247:
		description = "DVT DATA LINK"
	case 3248:
		description = "PROCOS LM"
	case 3249:
		description = "State Sync Protocol"
	case 3250:
		description = "HMS hicp port"
	case 3251:
		description = "Sys Scanner"
	case 3252:
		description = "DHE port"
	case 3253:
		description = "PDA Data"
	case 3254:
		description = "PDA System"
	case 3255:
		description = "Semaphore Connection Port"
	case 3256:
		description = "Compaq RPM Agent Port"
	case 3257:
		description = "Compaq RPM Server Port"
	case 3258:
		description = "Ivecon Server Port"
	case 3259:
		description = "Epson Network Common Devi"
	case 3260:
		description = "iSCSI port"
	case 3261:
		description = "winShadow"
	case 3262:
		description = "NECP"
	case 3263:
		description = "E-Color Enterprise Imager"
	case 3264:
		description = "cc:mail/lotus"
	case 3265:
		description = "Altav Tunnel"
	case 3266:
		description = "NS CFG Server"
	case 3267:
		description = "IBM Dial Out"
	case 3268:
		description = "Microsoft Global Catalog"
	case 3269:
		description = "Microsoft Global Catalog with LDAP/SSL"
	case 3270:
		description = "Verismart"
	case 3271:
		description = "CSoft Prev Port"
	case 3272:
		description = "Fujitsu User Manager"
	case 3273:
		description = "Simple Extensible Multiplexed Protocol"
	case 3274:
		description = "Ordinox Server"
	case 3275:
		description = "SAMD"
	case 3276:
		description = "Maxim ASICs"
	case 3277:
		description = "AWG Proxy"
	case 3278:
		description = "LKCM Server"
	case 3279:
		description = "admind"
	case 3280:
		description = "VS Server"
	case 3281:
		description = "SYSOPT"
	case 3282:
		description = "Datusorb"
	case 3283:
		description = "Net Assistant"
	case 3284:
		description = "4Talk"
	case 3285:
		description = "Plato"
	case 3286:
		description = "E-Net"
	case 3287:
		description = "DIRECTVDATA"
	case 3288:
		description = "COPS"
	case 3289:
		description = "dasHost.exe / ENPC"
	case 3290:
		description = "CAPS LOGISTICS TOOLKIT - LM"
	case 3291:
		description = "S A Holditch & Associates - LM"
	case 3292:
		description = "Cart O Rama"
	case 3293:
		description = "fg-fps"
	case 3294:
		description = "fg-gip"
	case 3295:
		description = "Dynamic IP Lookup"
	case 3296:
		description = "Rib License Manager"
	case 3297:
		description = "Cytel License Manager"
	case 3298:
		description = "DeskView"
	case 3299:
		description = "pdrncs"
	case 3300:
		description = "Ceph monitor"
	case 3301:
		description = "Tarantool in-memory computing platform"
	case 3302:
		description = "MCS Fastmail"
	case 3303:
		description = "OP Session Client"
	case 3304:
		description = "OP Session Server"
	case 3305:
		description = "ODETTE-FTP"
	case 3306:
		description = "MySQL"
	case 3307:
		description = "OP Session Proxy"
	case 3308:
		description = "TNS Server"
	case 3309:
		description = "TNS ADV"
	case 3310:
		description = "Dyna Access"
	case 3311:
		description = "MCNS Tel Ret"
	case 3312:
		description = "Application Management Server"
	case 3313:
		description = "Unify Object Broker"
	case 3314:
		description = "Unify Object Host"
	case 3315:
		description = "CDID"
	case 3316:
		description = "AICC/CMI"
	case 3317:
		description = "VSAI PORT"
	case 3318:
		description = "Swith to Swith Routing Information Protocol"
	case 3319:
		description = "SDT License Manager"
	case 3320:
		description = "Office Link 2000"
	case 3321:
		description = "VNSSTR"
	case 3326:
		description = "SFTU"
	case 3327:
		description = "BBARS"
	case 3328:
		description = "Eaglepoint License Manager"
	case 3329:
		description = "HP Device Disc"
	case 3330:
		description = "MCS Calypso ICF"
	case 3331:
		description = "MCS Messaging"
	case 3332:
		description = "MCS Mail Server"
	case 3333:
		description = "DEC Notes"
	case 3334:
		description = "Direct TV Webcasting"
	case 3335:
		description = "Direct TV Software Updates"
	case 3336:
		description = "Direct TV Tickers"
	case 3337:
		description = "Direct TV Data Catalog"
	case 3338:
		description = "OMF data b"
	case 3339:
		description = "OMF data l"
	case 3340:
		description = "OMF data m"
	case 3341:
		description = "OMF data h"
	case 3342:
		description = "WebTIE"
	case 3343:
		description = "MS Cluster Net"
	case 3344:
		description = "BNT Manager"
	case 3345:
		description = "Influence"
	case 3346:
		description = "Trnsprnt Proxy"
	case 3347:
		description = "Phoenix RPC"
	case 3348:
		description = "Pangolin Laser"
	case 3349:
		description = "Chevin Services"
	case 3350:
		description = "FINDVIATV"
	case 3351:
		description = "Btrieve port"
	case 3352:
		description = "Scalable SQL"
	case 3353:
		description = "FATPIPE"
	case 3354:
		description = "SUITJD"
	case 3355:
		description = "Ordinox Dbase"
	case 3356:
		description = "UPNOTIFYPS"
	case 3357:
		description = "Adtech Test IP"
	case 3358:
		description = "Mp Sys Rmsvr"
	case 3359:
		description = "WG NetForce"
	case 3360:
		description = "KV Server"
	case 3361:
		description = "KV Agent"
	case 3362:
		description = "DJ ILM"
	case 3363:
		description = "NATI Vi Server"
	case 3364:
		description = "Creative Server"
	case 3365:
		description = "Content Server"
	case 3366:
		description = "Creative Partner"
	case 3372:
		description = "TIP 2"
	case 3373:
		description = "Lavenir License Manager"
	case 3374:
		description = "Cluster Disc"
	case 3375:
		description = "VSNM Agent"
	case 3376:
		description = "CD Broker"
	case 3377:
		description = "Cogsys Network License Manager"
	case 3378:
		description = "WSICOPY"
	case 3379:
		description = "SOCORFS"
	case 3380:
		description = "SNS Channels"
	case 3381:
		description = "Geneous"
	case 3382:
		description = "Fujitsu Network Enhanced Antitheft function"
	case 3383:
		description = "Enterprise Software Products License Manager"
	case 3384:
		description = "Hardware Management"
	case 3385:
		description = "qnxnetman"
	case 3386:
		description = "GPRS SIG"
	case 3387:
		description = "Back Room Net"
	case 3388:
		description = "CB Server"
	case 3389:
		description = "MS WBT Server"
	case 3390:
		description = "Distributed Service Coordinator"
	case 3391:
		description = "SAVANT"
	case 3392:
		description = "EFI License Management"
	case 3393:
		description = "D2K Tapestry Client to Server"
	case 3394:
		description = "D2K Tapestry Server to Server"
	case 3395:
		description = "Dyna License Manager (Elam)"
	case 3396:
		description = "Printer Agent"
	case 3397:
		description = "Cloanto License Manager"
	case 3398:
		description = "Mercantile"
	case 3399:
		description = "CSMS"
	case 3400:
		description = "CSMS2"
	case 3401:
		description = "filecast"
	case 3402:
		description = "FXa Engine Network Port"
	case 3403:
		description = "De-registered"
	case 3404:
		description = "Removed"
	case 3405:
		description = "Nokia Announcement ch 1"
	case 3406:
		description = "Nokia Announcement ch 2"
	case 3407:
		description = "LDAP admin server port"
	case 3408:
		description = "BES Api Port"
	case 3409:
		description = "NetworkLens Event Port"
	case 3410:
		description = "NetworkLens SSL Event"
	case 3411:
		description = "BioLink Authenteon server"
	case 3412:
		description = "xmlBlaster"
	case 3413:
		description = "SpecView Networking"
	case 3414:
		description = "BroadCloud WIP Port"
	case 3415:
		description = "BCI Name Service"
	case 3416:
		description = "AirMobile IS Command Port"
	case 3417:
		description = "ConServR file translation"
	case 3418:
		description = "Remote nmap"
	case 3419:
		description = "ISogon SoftAudit"
	case 3420:
		description = "iFCP User Port"
	case 3421:
		description = "Bull Apprise portmapper"
	case 3422:
		description = "Remote USB System Port"
	case 3423:
		description = "xTrade Reliable Messaging"
	case 3424:
		description = "xTrade over TLS/SSL"
	case 3425:
		description = "AGPS Access Port"
	case 3426:
		description = "Arkivio Storage Protocol"
	case 3427:
		description = "WebSphere SNMP"
	case 3428:
		description = "2Wire CSS"
	case 3429:
		description = "GCSP user port"
	case 3430:
		description = "Scott Studios Dispatch"
	case 3431:
		description = "Active License Server Port"
	case 3432:
		description = "Secure Device Protocol"
	case 3433:
		description = "OPNET Service Management Platform"
	case 3434:
		description = "OpenCM Server"
	case 3435:
		description = "Pacom Security User Port"
	case 3436:
		description = "GuardControl Exchange Protocol"
	case 3437:
		description = "Autocue Directory Service"
	case 3438:
		description = "Spiralcraft Admin"
	case 3439:
		description = "HRI Interface Port"
	case 3440:
		description = "Net Steward Mgmt Console"
	case 3441:
		description = "OC Connect Client"
	case 3442:
		description = "OC Connect Server"
	case 3443:
		description = "OpenView Network Node Manager WEB Server"
	case 3444:
		description = "Denali Server"
	case 3445:
		description = "Media Object Network Protocol"
	case 3446:
		description = "3Com FAX RPC port"
	case 3447:
		description = "DirectNet IM System"
	case 3448:
		description = "Discovery and Net Config"
	case 3449:
		description = "HotU Chat"
	case 3450:
		description = "CAStorProxy"
	case 3451:
		description = "ASAM Services"
	case 3452:
		description = "SABP-Signalling Protocol"
	case 3453:
		description = "PSC Update"
	case 3454:
		description = "Apple Remote Access Protocol"
	case 3455:
		description = "RSVP Port"
	case 3456:
		description = "VAT default data"
	case 3457:
		description = "VAT default control"
	case 3458:
		description = "D3WinOSFI"
	case 3459:
		description = "TIP Integral"
	case 3460:
		description = "EDM Manger"
	case 3461:
		description = "EDM Stager"
	case 3462:
		description = "EDM STD Notify"
	case 3463:
		description = "EDM ADM Notify"
	case 3464:
		description = "EDM MGR Sync"
	case 3465:
		description = "EDM MGR Cntrl"
	case 3466:
		description = "WORKFLOW"
	case 3467:
		description = "RCST"
	case 3468:
		description = "TTCM Remote Controll"
	case 3469:
		description = "Pluribus"
	case 3470:
		description = "jt400"
	case 3471:
		description = "jt400-ssl"
	case 3472:
		description = "JAUGS N-G Remotec 1"
	case 3473:
		description = "JAUGS N-G Remotec 2"
	case 3474:
		description = "TSP Automation"
	case 3475:
		description = "Genisar Comm Port"
	case 3476:
		description = "NVIDIA Mgmt Protocol"
	case 3477:
		description = "eComm link port"
	case 3478:
		description = "Session Traversal Utilities for NAT (STUN) port"
	case 3479:
		description = "2Wire RPC"
	case 3480:
		description = "Secure Virtual Workspace"
	case 3481:
		description = "CleanerLive remote ctrl"
	case 3482:
		description = "Vulture Monitoring System"
	case 3483:
		description = "Slim Devices Protocol"
	case 3484:
		description = "GBS SnapTalk Protocol"
	case 3485:
		description = "CelaTalk"
	case 3486:
		description = "IFSF Heartbeat Port"
	case 3487:
		description = "LISA Transfer Channel"
	case 3488:
		description = "FS Remote Host Server"
	case 3489:
		description = "DTP/DIA"
	case 3490:
		description = "Colubris Management Port"
	case 3491:
		description = "SWR Port"
	case 3492:
		description = "TVDUM Tray Port"
	case 3493:
		description = "Network UPS Tools"
	case 3494:
		description = "IBM 3494"
	case 3495:
		description = "securitylayer over tcp"
	case 3496:
		description = "securitylayer over tls"
	case 3497:
		description = "ipEther232Port"
	case 3498:
		description = "DASHPAS user port"
	case 3499:
		description = "SccIP Media"
	case 3500:
		description = "RTMP Port"
	case 3501:
		description = "iSoft-P2P"
	case 3502:
		description = "Avocent Install Discovery"
	case 3503:
		description = "MPLS LSP-echo Port"
	case 3504:
		description = "IronStorm game server"
	case 3505:
		description = "CCM communications port"
	case 3506:
		description = "APC 3506"
	case 3507:
		description = "Nesh Broker Port"
	case 3508:
		description = "Interaction Web"
	case 3509:
		description = "Virtual Token SSL Port"
	case 3510:
		description = "XSS Port"
	case 3511:
		description = "WebMail/2"
	case 3512:
		description = "Aztec Distribution Port"
	case 3513:
		description = "Adaptec Remote Protocol"
	case 3514:
		description = "MUST Peer to Peer"
	case 3515:
		description = "MUST Backplane"
	case 3516:
		description = "Smartcard Port"
	case 3517:
		description = "IEEE 802.11 WLANs WG IAPP"
	case 3518:
		description = "Artifact Message Server"
	case 3519:
		description = "Netvion Messenger Port"
	case 3520:
		description = "Netvion Galileo Log Port"
	case 3521:
		description = "Telequip Labs MC3SS"
	case 3522:
		description = "DO over NSSocketPort"
	case 3523:
		description = "Odeum Serverlink"
	case 3524:
		description = "ECM Server port"
	case 3525:
		description = "EIS Server port"
	case 3526:
		description = "starQuiz Port"
	case 3527:
		description = "VERITAS Backup Exec Server"
	case 3528:
		description = "JBoss IIOP"
	case 3529:
		description = "JBoss IIOP/SSL"
	case 3530:
		description = "Grid Friendly"
	case 3531:
		description = "Joltid"
	case 3532:
		description = "Raven Remote Management Control"
	case 3533:
		description = "Raven Remote Management Data"
	case 3534:
		description = "URL Daemon Port"
	case 3535:
		description = "MS-LA"
	case 3536:
		description = "SNAC"
	case 3537:
		description = "Remote NI-VISA port"
	case 3538:
		description = "IBM Directory Server"
	case 3539:
		description = "IBM Directory Server SSL"
	case 3540:
		description = "PNRP User Port"
	case 3541:
		description = "VoiSpeed Port"
	case 3542:
		description = "HA cluster monitor"
	case 3543:
		description = "qftest Lookup Port"
	case 3544:
		description = "Teredo Port"
	case 3545:
		description = "CAMAC equipment"
	case 3547:
		description = "Symantec SIM"
	case 3548:
		description = "Interworld"
	case 3549:
		description = "Tellumat MDR NMS"
	case 3550:
		description = "Secure SMPP"
	case 3551:
		description = "Apcupsd Information Port"
	case 3552:
		description = "TeamAgenda Server Port"
	case 3553:
		description = "Red Box Recorder ADP"
	case 3554:
		description = "Quest Notification Server"
	case 3555:
		description = "Vipul's Razor"
	case 3556:
		description = "Sky Transport Protocol"
	case 3557:
		description = "PersonalOS Comm Port"
	case 3558:
		description = "MCP user port"
	case 3559:
		description = "CCTV control port"
	case 3560:
		description = "INIServe port"
	case 3561:
		description = "BMC-OneKey"
	case 3562:
		description = "SDBProxy"
	case 3563:
		description = "Watcom Debug"
	case 3564:
		description = "Electromed SIM port"
	case 3565:
		description = "M2PA"
	case 3566:
		description = "Quest Data Hub"
	case 3567:
		description = "DOF Protocol Stack"
	case 3568:
		description = "DOF Secure Tunnel"
	case 3569:
		description = "Meinberg Control Service"
	case 3570:
		description = "MCC Web Server Port"
	case 3571:
		description = "MegaRAID Server Port"
	case 3572:
		description = "Registration Server Port"
	case 3573:
		description = "Advantage Group UPS Suite"
	case 3574:
		description = "DMAF Server"
	case 3575:
		description = "Coalsere CCM Port"
	case 3576:
		description = "Coalsere CMC Port"
	case 3577:
		description = "Configuration Port"
	case 3578:
		description = "Data Port"
	case 3579:
		description = "Tarantella Load Balancing"
	case 3580:
		description = "NATI-ServiceLocator"
	case 3581:
		description = "Ascent Capture Licensing"
	case 3582:
		description = "PEG PRESS Server"
	case 3583:
		description = "CANEX Watch System"
	case 3584:
		description = "U-DBase Access Protocol"
	case 3585:
		description = "Emprise License Server"
	case 3586:
		description = "License Server Console"
	case 3587:
		description = "Peer to Peer Grouping"
	case 3588:
		description = "Sentinel Server"
	case 3589:
		description = "isomair"
	case 3590:
		description = "WV CSP SMS Binding"
	case 3591:
		description = "LOCANIS G-TRACK Server"
	case 3592:
		description = "LOCANIS G-TRACK NE Port"
	case 3593:
		description = "BP Model Debugger"
	case 3594:
		description = "MediaSpace"
	case 3595:
		description = "ShareApp"
	case 3596:
		description = "Illusion Wireless MMOG"
	case 3597:
		description = "A14 (AN-to-SC/MM)"
	case 3598:
		description = "A15 (AN-to-AN)"
	case 3599:
		description = "Quasar Accounting Server"
	case 3600:
		description = "text relay-answer"
	case 3601:
		description = "Visinet Gui"
	case 3602:
		description = "InfiniSwitch Mgr Client"
	case 3603:
		description = "Integrated Rcvr Control"
	case 3604:
		description = "BMC JMX Port"
	case 3605:
		description = "ComCam IO Port"
	case 3606:
		description = "Splitlock Server"
	case 3607:
		description = "Precise I3"
	case 3608:
		description = "Trendchip control protocol"
	case 3609:
		description = "CPDI PIDAS Connection Mon"
	case 3610:
		description = "ECHONET"
	case 3611:
		description = "Six Degrees Port"
	case 3612:
		description = "Micro Focus Data Protector"
	case 3613:
		description = "Alaris Device Discovery"
	case 3614:
		description = "Satchwell Sigma"
	case 3615:
		description = "Start Messaging Network"
	case 3616:
		description = "cd3o Control Protocol"
	case 3617:
		description = "ATI SHARP Logic Engine"
	case 3618:
		description = "AAIR-Network 1"
	case 3619:
		description = "AAIR-Network 2"
	case 3620:
		description = "EPSON Projector Control Port"
	case 3621:
		description = "EPSON Network Screen Port"
	case 3622:
		description = "FF LAN Redundancy Port"
	case 3623:
		description = "HAIPIS Dynamic Discovery"
	case 3624:
		description = "Distributed Upgrade Port"
	case 3625:
		description = "Volley"
	case 3626:
		description = "bvControl Daemon"
	case 3627:
		description = "Jam Server Port"
	case 3628:
		description = "EPT Machine Interface"
	case 3629:
		description = "ESC/VP.net"
	case 3630:
		description = "C&S Remote Database Port"
	case 3631:
		description = "C&S Web Services Port"
	case 3632:
		description = "distributed compiler"
	case 3633:
		description = "Wyrnix AIS port"
	case 3634:
		description = "hNTSP Library Manager"
	case 3635:
		description = "Simple Distributed Objects"
	case 3636:
		description = "SerVistaITSM"
	case 3637:
		description = "Customer Service Port"
	case 3638:
		description = "EHP Backup Protocol"
	case 3639:
		description = "Extensible Automation"
	case 3640:
		description = "Netplay Port 1"
	case 3641:
		description = "Netplay Port 2"
	case 3642:
		description = "Juxml Replication port"
	case 3643:
		description = "AudioJuggler"
	case 3644:
		description = "ssowatch"
	case 3645:
		description = "Cyc"
	case 3646:
		description = "XSS Server Port"
	case 3647:
		description = "Splitlock Gateway"
	case 3648:
		description = "Fujitsu Cooperation Port"
	case 3649:
		description = "Nishioka Miyuki Msg Protocol"
	case 3650:
		description = "PRISMIQ VOD plug-in"
	case 3651:
		description = "XRPC Registry"
	case 3652:
		description = "VxCR NBU Default Port"
	case 3653:
		description = "Tunnel Setup Protocol"
	case 3654:
		description = "VAP RealTime Messenger"
	case 3655:
		description = "ActiveBatch Exec Agent"
	case 3656:
		description = "ActiveBatch Job Scheduler"
	case 3657:
		description = "ImmediaNet Beacon"
	case 3658:
		description = "PlayStation AMS (Secure)"
	case 3659:
		description = "Apple SASL"
	case 3660:
		description = "IBM Tivoli Directory Service using SSL"
	case 3661:
		description = "IBM Tivoli Directory Service using SSL"
	case 3662:
		description = "pserver"
	case 3663:
		description = "DIRECWAY Tunnel Protocol"
	case 3664:
		description = "UPS Engine Port"
	case 3665:
		description = "Enterprise Engine Port"
	case 3666:
		description = "IBM EServer PAP"
	case 3667:
		description = "IBM Information Exchange"
	case 3668:
		description = "Dell Remote Management"
	case 3669:
		description = "CA SAN Switch Management"
	case 3670:
		description = "SMILE TCP/UDP Interface"
	case 3671:
		description = "e Field Control (EIBnet)"
	case 3672:
		description = "LispWorks ORB"
	case 3673:
		description = "Openview Media Vault GUI"
	case 3674:
		description = "WinINSTALL IPC Port"
	case 3675:
		description = "CallTrax Data Port"
	case 3676:
		description = "VisualAge Pacbase server"
	case 3677:
		description = "RoverLog IPC"
	case 3678:
		description = "DataGuardianLT"
	case 3679:
		description = "Newton Dock"
	case 3680:
		description = "NPDS Tracker"
	case 3681:
		description = "BTS X73 Port"
	case 3682:
		description = "EMC SmartPackets-MAPI"
	case 3683:
		description = "BMC EDV/EA"
	case 3684:
		description = "FAXstfX"
	case 3685:
		description = "DS Expert Agent"
	case 3686:
		description = "Trivial Network Management"
	case 3687:
		description = "simple-push"
	case 3688:
		description = "simple-push Secure"
	case 3689:
		description = "Digital Audio Access Protocol (iTunes)"
	case 3690:
		description = "Subversion"
	case 3691:
		description = "Magaya Network Port"
	case 3692:
		description = "Brimstone IntelSync"
	case 3693:
		description = "Emergency Automatic Structure Lockdown System"
	case 3695:
		description = "BMC Data Collection"
	case 3696:
		description = "Telnet Com Port Control"
	case 3697:
		description = "NavisWorks License System"
	case 3698:
		description = "SAGECTLPANEL"
	case 3699:
		description = "Internet Call Waiting"
	case 3700:
		description = "LRS NetPage"
	case 3701:
		description = "NetCelera"
	case 3702:
		description = "Web Service Discovery"
	case 3703:
		description = "Adobe Server 3"
	case 3704:
		description = "Adobe Server 4"
	case 3705:
		description = "Adobe Server 5"
	case 3706:
		description = "Real-Time Event Port"
	case 3707:
		description = "Real-Time Event Secure Port"
	case 3708:
		description = "Sun App Svr - Naming"
	case 3709:
		description = "CA-IDMS Server"
	case 3710:
		description = "PortGate Authentication"
	case 3711:
		description = "EBD Server 2"
	case 3712:
		description = "Sentinel Enterprise"
	case 3713:
		description = "TFTP over TLS"
	case 3714:
		description = "DELOS Direct Messaging"
	case 3715:
		description = "Anoto Rendezvous Port"
	case 3716:
		description = "WV CSP SMS CIR Channel"
	case 3717:
		description = "WV CSP UDP/IP CIR Channel"
	case 3718:
		description = "OPUS Server Port"
	case 3719:
		description = "iTel Server Port"
	case 3720:
		description = "UF Astro. Instr. Services"
	case 3721:
		description = "Xsync"
	case 3722:
		description = "Xserve RAID"
	case 3723:
		description = "Sychron Service Daemon"
	case 3724:
		description = "World of Warcraft"
	case 3725:
		description = "Netia NA-ER Port"
	case 3726:
		description = "Xyratex Array Manager"
	case 3727:
		description = "Ericsson Mobile Data Unit"
	case 3728:
		description = "Ericsson Web on Air"
	case 3729:
		description = "Fireking Audit Port"
	case 3730:
		description = "Client Control"
	case 3731:
		description = "Service Manager"
	case 3732:
		description = "Mobile Wnn"
	case 3733:
		description = "Multipuesto Msg Port"
	case 3734:
		description = "Synel Data Collection Port"
	case 3735:
		description = "Password Distribution"
	case 3736:
		description = "RealSpace RMI"
	case 3737:
		description = "XPanel Daemon"
	case 3738:
		description = "versaTalk Server Port"
	case 3739:
		description = "Launchbird LicenseManager"
	case 3740:
		description = "Heartbeat Protocol"
	case 3741:
		description = "WysDM Agent"
	case 3742:
		description = "CST - Configuration & Service Tracker"
	case 3743:
		description = "IP Control Systems Ltd."
	case 3744:
		description = "SASG"
	case 3745:
		description = "GWRTC Call Port"
	case 3746:
		description = "LXPRO.COM LinkTest"
	case 3747:
		description = "LXPRO.COM LinkTest SSL"
	case 3748:
		description = "webData"
	case 3749:
		description = "CimTrak"
	case 3750:
		description = "CBOS/IP ncapsalation port"
	case 3751:
		description = "CommLinx GPRS Cube"
	case 3752:
		description = "Vigil-IP RemoteAgent"
	case 3753:
		description = "NattyServer Port"
	case 3754:
		description = "TimesTen Broker Port"
	case 3755:
		description = "SAS Remote Help Server"
	case 3756:
		description = "Canon CAPT Port"
	case 3757:
		description = "GRF Server Port"
	case 3758:
		description = "apw RMI registry"
	case 3759:
		description = "Exapt License Manager"
	case 3760:
		description = "adTEmpus Client"
	case 3761:
		description = "gsakmp port"
	case 3762:
		description = "GBS SnapMail Protocol"
	case 3763:
		description = "XO Wave Control Port"
	case 3764:
		description = "MNI Protected Routing"
	case 3765:
		description = "Remote Traceroute"
	case 3766:
		description = "SSL e-watch sitewatch server"
	case 3767:
		description = "ListMGR Port"
	case 3768:
		description = "rblcheckd server daemon"
	case 3769:
		description = "HAIPE Network Keying"
	case 3770:
		description = "Cinderella Collaboration"
	case 3771:
		description = "RTP Paging Port"
	case 3772:
		description = "Chantry Tunnel Protocol"
	case 3773:
		description = "ctdhercules"
	case 3774:
		description = "ZICOM"
	case 3775:
		description = "ISPM Manager Port"
	case 3776:
		description = "Device Provisioning Port"
	case 3777:
		description = "Jibe EdgeBurst"
	case 3778:
		description = "Cutler-Hammer IT Port"
	case 3779:
		description = "Cognima Replication"
	case 3780:
		description = "Nuzzler Network Protocol"
	case 3781:
		description = "ABCvoice server port"
	case 3782:
		description = "Secure ISO TP0 port"
	case 3783:
		description = "Impact Mgr./PEM Gateway"
	case 3784:
		description = "BFD Control Protocol"
	case 3785:
		description = "BFD Echo Protocol"
	case 3786:
		description = "VSW Upstrigger port"
	case 3787:
		description = "Fintrx"
	case 3788:
		description = "SPACEWAY Routing port"
	case 3789:
		description = "RemoteDeploy Administration Port [July 2003]"
	case 3790:
		description = "QuickBooks RDS"
	case 3791:
		description = "TV NetworkVideo Data port"
	case 3792:
		description = "e-Watch Corporation SiteWatch"
	case 3793:
		description = "DataCore Software"
	case 3794:
		description = "JAUS Robots"
	case 3795:
		description = "myBLAST Mekentosj port"
	case 3796:
		description = "Spaceway Dialer"
	case 3797:
		description = "idps"
	case 3798:
		description = "Minilock"
	case 3799:
		description = "RADIUS Dynamic Authorization"
	case 3800:
		description = "Print Services Interface"
	case 3801:
		description = "ibm manager service"
	case 3802:
		description = "VHD"
	case 3803:
		description = "SoniqSync"
	case 3804:
		description = "Harman IQNet Port"
	case 3805:
		description = "ThorGuard Server Port"
	case 3806:
		description = "Remote System Manager"
	case 3807:
		description = "SpuGNA Communication Port"
	case 3808:
		description = "Sun App Svr-IIOPClntAuth"
	case 3809:
		description = "Java Desktop System Configuration Agent"
	case 3810:
		description = "WLAN AS server"
	case 3811:
		description = "AMP"
	case 3812:
		description = "netO WOL Server"
	case 3813:
		description = "Rhapsody Interface Protocol"
	case 3814:
		description = "netO DCS"
	case 3815:
		description = "LANsurveyor XML"
	case 3816:
		description = "Sun Local Patch Server"
	case 3817:
		description = "Yosemite Tech Tapeware"
	case 3818:
		description = "Crinis Heartbeat"
	case 3819:
		description = "EPL Sequ Layer Protocol"
	case 3820:
		description = "Siemens AuD SCP"
	case 3821:
		description = "ATSC PMCP Standard"
	case 3822:
		description = "Compute Pool Discovery"
	case 3823:
		description = "Compute Pool Conduit"
	case 3824:
		description = "Compute Pool Policy"
	case 3825:
		description = "Antera FlowFusion Process Simulation"
	case 3826:
		description = "WarMUX game server"
	case 3827:
		description = "Netadmin Systems MPI service"
	case 3828:
		description = "Netadmin Systems Event Handler"
	case 3829:
		description = "Netadmin Systems Event Handler External"
	case 3830:
		description = "Cerner System Management Agent"
	case 3831:
		description = "Docsvault Application Service"
	case 3832:
		description = "xxNETserver"
	case 3833:
		description = "AIPN LS Authentication"
	case 3834:
		description = "Spectar Data Stream Service"
	case 3835:
		description = "Spectar Database Rights Service"
	case 3836:
		description = "MARKEM NEXTGEN DCP"
	case 3837:
		description = "MARKEM Auto-Discovery"
	case 3838:
		description = "Scito Object Server"
	case 3839:
		description = "AMX Resource Management Suite"
	case 3840:
		description = "www.FlirtMitMir.de"
	case 3841:
		description = "ShipRush Database Server"
	case 3842:
		description = "NHCI status port"
	case 3843:
		description = "Quest Common Agent"
	case 3844:
		description = "RNM"
	case 3845:
		description = "V-ONE Single Port Proxy"
	case 3846:
		description = "Astare Network PCP"
	case 3847:
		description = "MS Firewall Control"
	case 3848:
		description = "IT Environmental Monitor"
	case 3849:
		description = "SPACEWAY DNS Preload"
	case 3850:
		description = "QTMS Bootstrap Protocol"
	case 3851:
		description = "SpectraTalk Port"
	case 3852:
		description = "SSE App Configuration"
	case 3853:
		description = "SONY scanning protocol"
	case 3854:
		description = "Stryker Comm Port"
	case 3855:
		description = "OpenTRAC"
	case 3856:
		description = "INFORMER"
	case 3857:
		description = "Trap Port"
	case 3858:
		description = "Trap Port MOM"
	case 3859:
		description = "Navini Port"
	case 3860:
		description = "Server/Application State Protocol (SASP)"
	case 3861:
		description = "winShadow Host Discovery"
	case 3862:
		description = "GIGA-POCKET"
	case 3863:
		description = "asap"
	case 3864:
		description = "asap-sctp/tls"
	case 3865:
		description = "xpl automation protocol"
	case 3866:
		description = "Sun SDViz DZDAEMON Port"
	case 3867:
		description = "Sun SDViz DZOGLSERVER Port"
	case 3868:
		description = "DIAMETER"
	case 3869:
		description = "hp OVSAM MgmtServer Disco"
	case 3870:
		description = "hp OVSAM HostAgent Disco"
	case 3871:
		description = "Avocent DS Authorization"
	case 3872:
		description = "OEM Agent"
	case 3873:
		description = "fagordnc"
	case 3874:
		description = "SixXS Configuration"
	case 3875:
		description = "PNBSCADA"
	case 3876:
		description = "DirectoryLockdown Agent"
	case 3877:
		description = "XMPCR Interface Port"
	case 3878:
		description = "FotoG CAD interface"
	case 3879:
		description = "appss license manager"
	case 3880:
		description = "IGRS"
	case 3881:
		description = "Data Acquisition and Control"
	case 3882:
		description = "DTS Service Port"
	case 3883:
		description = "VR Peripheral Network"
	case 3884:
		description = "SofTrack Metering"
	case 3885:
		description = "TopFlow SSL"
	case 3886:
		description = "NEI management port"
	case 3887:
		description = "Ciphire Data Transport"
	case 3888:
		description = "Ciphire Services"
	case 3889:
		description = "D and V Tester Control Port"
	case 3890:
		description = "Niche Data Server Connect"
	case 3891:
		description = "Oracle RTC-PM port"
	case 3892:
		description = "PCC-image-port"
	case 3893:
		description = "CGI StarAPI Server"
	case 3894:
		description = "SyAM Agent Port"
	case 3895:
		description = "SyAm SMC Service Port"
	case 3896:
		description = "Simple Distributed Objects over TLS"
	case 3897:
		description = "Simple Distributed Objects over SSH"
	case 3898:
		description = "IAS"
	case 3899:
		description = "ITV Port"
	case 3900:
		description = "Unidata UDT OS"
	case 3901:
		description = "NIM Service Handler"
	case 3902:
		description = "NIMsh Auxiliary Port"
	case 3903:
		description = "CharsetMGR"
	case 3904:
		description = "Arnet Omnilink Port"
	case 3905:
		description = "Mailbox Update (MUPDATE) protocol"
	case 3906:
		description = "TopoVista elevation data"
	case 3907:
		description = "Imoguia Port"
	case 3908:
		description = "HP Procurve NetManagement"
	case 3909:
		description = "SurfControl CPA"
	case 3910:
		description = "Printer Request Port"
	case 3911:
		description = "Printer Status Port"
	case 3912:
		description = "Global Maintech Stars"
	case 3913:
		description = "ListCREATOR Port"
	case 3914:
		description = "ListCREATOR Port 2"
	case 3915:
		description = "Auto-Graphics Cataloging"
	case 3916:
		description = "WysDM Controller"
	case 3917:
		description = "AFT multiplex port"
	case 3918:
		description = "PacketCableMultimediaCOPS"
	case 3919:
		description = "HyperIP"
	case 3920:
		description = "Exasoft IP Port"
	case 3921:
		description = "Herodotus Net"
	case 3922:
		description = "Soronti Update Port"
	case 3923:
		description = "Symbian Service Broker"
	case 3924:
		description = "MPL_GPRS_PORT"
	case 3925:
		description = "Zoran Media Port"
	case 3926:
		description = "WINPort"
	case 3927:
		description = "ScsTsr"
	case 3928:
		description = "PXE NetBoot Manager"
	case 3929:
		description = "AMS Port"
	case 3930:
		description = "Syam Web Server Port"
	case 3931:
		description = "MSR Plugin Port"
	case 3932:
		description = "Dynamic Site System"
	case 3933:
		description = "PL/B App Server User Port"
	case 3934:
		description = "PL/B File Manager Port"
	case 3935:
		description = "SDP Port Mapper Protocol"
	case 3936:
		description = "Mailprox"
	case 3937:
		description = "DVB Service Discovery"
	case 3938:
		description = "Oracle dbControl Agent po"
	case 3939:
		description = "Anti-virus Application Management Port"
	case 3940:
		description = "XeCP Node Service"
	case 3941:
		description = "Home Portal Web Server"
	case 3942:
		description = "satellite distribution"
	case 3943:
		description = "TetraNode Ip Gateway"
	case 3944:
		description = "S-Ops Management"
	case 3945:
		description = "EMCADS Server Port"
	case 3946:
		description = "BackupEDGE Server"
	case 3947:
		description = "Connect and Control Protocol for Consumer"
	case 3948:
		description = "Anton Paar Device Administration Protocol"
	case 3949:
		description = "Dynamic Routing Information Protocol"
	case 3950:
		description = "Name Munging"
	case 3951:
		description = "PWG IPP Facsimile"
	case 3952:
		description = "I3 Session Manager"
	case 3953:
		description = "Eydeas XMLink Connect"
	case 3954:
		description = "AD Replication RPC"
	case 3955:
		description = "p2pCommunity"
	case 3956:
		description = "GigE Vision Control"
	case 3957:
		description = "MQEnterprise Broker"
	case 3958:
		description = "MQEnterprise Agent"
	case 3959:
		description = "Tree Hopper Networking"
	case 3960:
		description = "Bess Peer Assessment"
	case 3961:
		description = "ProAxess Server"
	case 3962:
		description = "SBI Agent Protocol"
	case 3963:
		description = "Teran Hybrid Routing Protocol"
	case 3964:
		description = "SASG GPRS"
	case 3965:
		description = "Avanti IP to NCPE API"
	case 3966:
		description = "BuildForge Lock Manager"
	case 3967:
		description = "PPS Message Service"
	case 3968:
		description = "iAnywhere DBNS"
	case 3969:
		description = "Landmark Messages"
	case 3970:
		description = "LANrev Agent"
	case 3971:
		description = "LANrev Server"
	case 3972:
		description = "ict-control Protocol"
	case 3973:
		description = "ConnectShip Progistics"
	case 3974:
		description = "Remote Applicant Tracking Service"
	case 3975:
		description = "Air Shot"
	case 3976:
		description = "Server Automation Agent"
	case 3977:
		description = "Opsware Manager"
	case 3978:
		description = "Secured Configuration Server"
	case 3979:
		description = "Smith Micro Wide Area Network Service"
	case 3981:
		description = "Starfish System Admin"
	case 3982:
		description = "ESRI Image Server"
	case 3983:
		description = "ESRI Image Service"
	case 3984:
		description = "MAPPER network node manager"
	case 3985:
		description = "MAPPER TCP/IP server"
	case 3986:
		description = "MAPPER workstation server"
	case 3987:
		description = "Centerline"
	case 3988:
		description = "DCS Configuration Port"
	case 3989:
		description = "BindView-Query Engine"
	case 3990:
		description = "BindView-IS"
	case 3991:
		description = "BindView-SMCServer"
	case 3992:
		description = "BindView-DirectoryServer"
	case 3993:
		description = "BindView-Agent"
	case 3995:
		description = "ISS Management Svcs SSL"
	case 3996:
		description = "abcsoftware-01"
	case 3997:
		description = "aes_db"
	case 3998:
		description = "Distributed Nagios Executor Service"
	case 3999:
		description = "Norman distributes scanning service"
	case 4000:
		description = "Terabase"
	case 4001:
		description = "NewOak"
	case 4002:
		description = "pxc-spvr-ft"
	case 4003:
		description = "pxc-splr-ft"
	case 4004:
		description = "pxc-roid"
	case 4005:
		description = "pxc-pin"
	case 4006:
		description = "pxc-spvr"
	case 4007:
		description = "pxc-splr"
	case 4008:
		description = "NetCheque accounting"
	case 4009:
		description = "Chimera HWM"
	case 4010:
		description = "Samsung Unidex"
	case 4011:
		description = "Alternate Service Boot"
	case 4012:
		description = "PDA Gate"
	case 4013:
		description = "ACL Manager"
	case 4014:
		description = "TAICLOCK"
	case 4015:
		description = "Talarian Mcast"
	case 4016:
		description = "Talarian Mcast"
	case 4017:
		description = "Talarian Mcast"
	case 4018:
		description = "Talarian Mcast"
	case 4019:
		description = "Talarian Mcast"
	case 4020:
		description = "TRAP Port"
	case 4021:
		description = "Nexus Portal"
	case 4022:
		description = "DNOX"
	case 4023:
		description = "ESNM Zoning Port"
	case 4024:
		description = "TNP1 User Port"
	case 4025:
		description = "Partition Image Port"
	case 4026:
		description = "Graphical Debug Server"
	case 4027:
		description = "bitxpress"
	case 4028:
		description = "DTServer Port"
	case 4029:
		description = "IP Q signaling protocol"
	case 4030:
		description = "Accell/JSP Daemon Port"
	case 4031:
		description = "UUCP over SSL"
	case 4032:
		description = "VERITAS Authorization Service"
	case 4033:
		description = "SANavigator Peer Port"
	case 4034:
		description = "Ubiquinox Daemon"
	case 4035:
		description = "WAP Push OTA-HTTP port"
	case 4036:
		description = "WAP Push OTA-HTTP secure"
	case 4037:
		description = "RaveHD network control"
	case 4038:
		description = "Fazzt Point-To-Point"
	case 4039:
		description = "Fazzt Administration"
	case 4040:
		description = "Yo.net main service"
	case 4041:
		description = "Rocketeer-Houston"
	case 4042:
		description = "LDXP"
	case 4043:
		description = "Neighbour Identity Resolution"
	case 4044:
		description = "Location Tracking Protocol"
	case 4045:
		description = "Network Paging Protocol"
	case 4046:
		description = "Accounting Protocol"
	case 4047:
		description = "Context Transfer Protocol"
	case 4049:
		description = "Wide Area File Services"
	case 4050:
		description = "Wide Area File Services"
	case 4051:
		description = "Cisco Peer to Peer Distribution Protocol"
	case 4052:
		description = "VoiceConnect Interact"
	case 4053:
		description = "CosmoCall Universe Communications Port 1"
	case 4054:
		description = "CosmoCall Universe Communications Port 2"
	case 4055:
		description = "CosmoCall Universe Communications Port 3"
	case 4056:
		description = "Location Message Service"
	case 4057:
		description = "Servigistics WFM server"
	case 4058:
		description = "Kingfisher protocol"
	case 4059:
		description = "DLMS/COSEM"
	case 4060:
		description = "DSMETER Inter-Agent Transfer Channel"
	case 4061:
		description = "Ice Location Service (TCP)"
	case 4062:
		description = "Ice Location Service (SSL)"
	case 4063:
		description = "Ice Firewall Traversal Service (TCP)"
	case 4064:
		description = "Ice Firewall Traversal Service (SSL)"
	case 4065:
		description = "Avanti Common Data"
	case 4066:
		description = "Performance Measurement and Analysis"
	case 4067:
		description = "Information Distribution Protocol"
	case 4068:
		description = "IP Fleet Broadcast"
	case 4069:
		description = "Minger Email Address Validation Service"
	case 4070:
		description = "Trivial IP Encryption (TrIPE)"
	case 4071:
		description = "Automatically Incremental Backup"
	case 4072:
		description = "Zieto Socket Communications"
	case 4073:
		description = "Interactive Remote Application Pairing Protocol"
	case 4074:
		description = "Cequint City ID UI trigger"
	case 4075:
		description = "ISC Alarm Message Service"
	case 4076:
		description = "Seraph DCS"
	case 4077:
		description = "Ascom IP Alarming"
	case 4078:
		description = "Coordinated Security Service Protocol"
	case 4079:
		description = "SANtools Diagnostic Server"
	case 4080:
		description = "Lorica inside facing"
	case 4081:
		description = "Lorica inside facing (SSL)"
	case 4082:
		description = "Lorica outside facing"
	case 4083:
		description = "Lorica outside facing (SSL)"
	case 4084:
		description = "Fortisphere VM Service"
	case 4085:
		description = "EZNews Newsroom Message Service"
	case 4086:
		description = "Firewall/NAT state table synchronization"
	case 4087:
		description = "APplus Service"
	case 4088:
		description = "Noah Printing Service Protocol"
	case 4089:
		description = "OpenCORE Remote Control Service"
	case 4090:
		description = "OMA BCAST Service Guide"
	case 4091:
		description = "EminentWare Installer"
	case 4092:
		description = "EminentWare DGS"
	case 4093:
		description = "Pvx Plus CS Host"
	case 4094:
		description = "sysrq daemon"
	case 4095:
		description = "xtgui information service"
	case 4096:
		description = "BRE (Bridge Relay Element)"
	case 4097:
		description = "Patrol View"
	case 4098:
		description = "drmsfsd"
	case 4099:
		description = "DPCP"
	case 4100:
		description = "IGo Incognito Data Port"
	case 4101:
		description = "Braille protocol"
	case 4102:
		description = "Braille protocol"
	case 4103:
		description = "Braille protocol"
	case 4104:
		description = "Braille protocol"
	case 4105:
		description = "Shofar"
	case 4106:
		description = "Synchronite"
	case 4107:
		description = "JDL Accounting LAN Service"
	case 4108:
		description = "ACCEL"
	case 4109:
		description = "Instantiated Zero-control Messaging"
	case 4110:
		description = "G2 RFID Tag Telemetry Data"
	case 4111:
		description = "Xgrid"
	case 4112:
		description = "Apple VPN Server Reporting Protocol"
	case 4113:
		description = "AIPN LS Registration"
	case 4114:
		description = "JomaMQMonitor"
	case 4115:
		description = "CDS Transfer Agent"
	case 4116:
		description = "smartcard-TLS"
	case 4117:
		description = "Hillr Connection Manager"
	case 4118:
		description = "Netadmin Systems NETscript service"
	case 4119:
		description = "Assuria Log Manager"
	case 4120:
		description = "MiniRem Remote Telemetry and Control"
	case 4121:
		description = "e-Builder Application Communication"
	case 4122:
		description = "Fiber Patrol Alarm Service"
	case 4123:
		description = "Z-Wave Protocol"
	case 4124:
		description = "Rohill TetraNode Ip Gateway v2"
	case 4125:
		description = "Opsview Envoy"
	case 4126:
		description = "Data Domain Replication Service"
	case 4127:
		description = "NetUniKeyServer"
	case 4128:
		description = "NuFW decision delegation protocol"
	case 4129:
		description = "NuFW authentication protocol"
	case 4130:
		description = "FRONET message protocol"
	case 4131:
		description = "Global Maintech Stars"
	case 4132:
		description = "NUTS Daemon"
	case 4133:
		description = "NUTS Bootp Server"
	case 4134:
		description = "NIFTY-Serve HMI protocol"
	case 4135:
		description = "Classic Line Database Server Attach"
	case 4136:
		description = "Classic Line Database Server Request"
	case 4137:
		description = "Classic Line Database Server Remote"
	case 4138:
		description = "nettest"
	case 4139:
		description = "Imperfect Networks Server"
	case 4140:
		description = "Cedros Fraud Detection System"
	case 4141:
		description = "Workflow Server"
	case 4142:
		description = "Document Server"
	case 4143:
		description = "Document Replication"
	case 4145:
		description = "VVR Control"
	case 4146:
		description = "TGCConnect Beacon"
	case 4147:
		description = "Multum Service Manager"
	case 4148:
		description = "HHB Handheld Client"
	case 4149:
		description = "A10 GSLB Service"
	case 4150:
		description = "PowerAlert Network Shutdown Agent"
	case 4151:
		description = "Men & Mice Remote Control"
	case 4152:
		description = "iDigTech Multiplex"
	case 4153:
		description = "MBL Remote Battery Monitoring"
	case 4154:
		description = "atlinks device discovery"
	case 4155:
		description = "Bazaar version control system"
	case 4156:
		description = "STAT Results"
	case 4157:
		description = "STAT Scanner Control"
	case 4158:
		description = "STAT Command Center"
	case 4159:
		description = "Network Security Service"
	case 4160:
		description = "Jini Discovery"
	case 4161:
		description = "OMS Contact"
	case 4162:
		description = "OMS Topology"
	case 4163:
		description = "Silver Peak Peer Protocol"
	case 4164:
		description = "Silver Peak Communication Protocol"
	case 4165:
		description = "ArcLink over Ethernet"
	case 4166:
		description = "Joost Peer to Peer Protocol"
	case 4167:
		description = "DeskDirect Global Network"
	case 4168:
		description = "PrintSoft License Server"
	case 4169:
		description = "Internet ADT Discovery Protocol"
	case 4170:
		description = "SMPTE Content Synchonization Protocol"
	case 4171:
		description = "Maxlogic Supervisor Communication"
	case 4172:
		description = "PC over IP"
	case 4173:
		description = "MMA Device Discovery"
	case 4174:
		description = "StorMagic Cluster Services"
	case 4175:
		description = "Brocade Cluster Communication Protocol"
	case 4176:
		description = "Translattice Cluster IPC Proxy"
	case 4177:
		description = "Wello P2P pubsub service"
	case 4178:
		description = "StorMan"
	case 4179:
		description = "Maxum Services"
	case 4180:
		description = "HTTPX"
	case 4181:
		description = "MacBak"
	case 4182:
		description = "Production Company Pro TCP Service"
	case 4183:
		description = "CyborgNet communications protocol"
	case 4184:
		description = "UNIVERSE SUITE MESSAGE SERVICE"
	case 4185:
		description = "Woven Control Plane Protocol"
	case 4186:
		description = "Box Backup Store Service"
	case 4187:
		description = "Cascade Proxy"
	case 4188:
		description = "Vatata Peer to Peer Protocol"
	case 4189:
		description = "Path Computation Element Communication Protocol"
	case 4190:
		description = "ManageSieve Protocol"
	case 4191:
		description = "Dual Stack MIPv6 NAT Traversal"
	case 4192:
		description = "azeti blinddate"
	case 4193:
		description = "PxPlus remote file srvr"
	case 4194:
		description = "Security Protocol and Data Model"
	case 4195:
		description = "AWS protocol for cloud remoting solution"
	case 4197:
		description = "Harman HControl Protocol"
	case 4199:
		description = "EIMS ADMIN"
	case 4300:
		description = "Corel CCam"
	case 4301:
		description = "Diagnostic Data"
	case 4302:
		description = "Diagnostic Data Control"
	case 4303:
		description = "Simple Railroad Command Protocol"
	case 4304:
		description = "One-Wire Filesystem Server"
	case 4305:
		description = "better approach to mobile ad-hoc networking"
	case 4306:
		description = "Hellgate London"
	case 4307:
		description = "TrueConf Videoconference Service"
	case 4308:
		description = "CompX-LockView"
	case 4309:
		description = "Exsequi Appliance Discovery"
	case 4310:
		description = "Mir-RT exchange service"
	case 4311:
		description = "P6R Secure Server Management Console"
	case 4312:
		description = "Parascale Membership Manager"
	case 4313:
		description = "PERRLA User Services"
	case 4314:
		description = "ChoiceView Agent"
	case 4316:
		description = "ChoiceView Client"
	case 4317:
		description = "OpenTelemetry Protocol"
	case 4319:
		description = "Fox SkyTale encrypted communication"
	case 4320:
		description = "FDT Remote Categorization Protocol"
	case 4321:
		description = "Remote Who Is"
	case 4322:
		description = "TRIM Event Service"
	case 4323:
		description = "TRIM ICE Service"
	case 4325:
		description = "Cadcorp GeognoSIS Administrator"
	case 4326:
		description = "Cadcorp GeognoSIS"
	case 4327:
		description = "Jaxer Web Protocol"
	case 4328:
		description = "Jaxer Manager Command Protocol"
	case 4329:
		description = "PubliQare Distributed Environment Synchronisation Engine"
	case 4330:
		description = "DEY Storage Administration REST API"
	case 4331:
		description = "ktickets REST API for event management and ticketing"
	case 4332:
		description = "Getty Images FOCUS service"
	case 4333:
		description = "ArrowHead Service Protocol (AHSP)"
	case 4334:
		description = "NETCONF Call Home (SSH)"
	case 4335:
		description = "NETCONF Call Home (TLS)"
	case 4336:
		description = "RESTCONF Call Home (TLS)"
	case 4340:
		description = "Gaia Connector Protocol"
	case 4341:
		description = "LISP Data Packets"
	case 4342:
		description = "LISP Control Packets"
	case 4343:
		description = "UNICALL"
	case 4344:
		description = "VinaInstall"
	case 4345:
		description = "Macro 4 Network AS"
	case 4346:
		description = "ELAN LM"
	case 4347:
		description = "LAN Surveyor"
	case 4348:
		description = "ITOSE"
	case 4349:
		description = "File System Port Map"
	case 4350:
		description = "Net Device"
	case 4351:
		description = "PLCY Net Services"
	case 4352:
		description = "Projector Link"
	case 4353:
		description = "F5 iQuery"
	case 4354:
		description = "QSNet Transmitter"
	case 4355:
		description = "QSNet Workstation"
	case 4356:
		description = "QSNet Assistant"
	case 4357:
		description = "QSNet Conductor"
	case 4358:
		description = "QSNet Nucleus"
	case 4359:
		description = "OMA BCAST Long-Term Key Messages"
	case 4360:
		description = "Matrix VNet Communication Protocol"
	case 4361:
		description = "NavCom Discovery and Control Port"
	case 4362:
		description = "AFORE vNode Discovery protocol"
	case 4366:
		description = "ShadowStream System"
	case 4368:
		description = "WeatherBrief Direct"
	case 4369:
		description = "Erlang Port Mapper Daemon"
	case 4370:
		description = "ELPRO V2 Protocol Tunnel"
	case 4371:
		description = "LAN2CAN Control"
	case 4372:
		description = "LAN2CAN Data"
	case 4373:
		description = "Remote Authenticated Command Service"
	case 4374:
		description = "PSI Push-to-Talk Protocol"
	case 4375:
		description = "Toltec EasyShare"
	case 4376:
		description = "BioAPI Interworking"
	case 4377:
		description = "Cambridge Pixel SPx Server"
	case 4378:
		description = "Cambridge Pixel SPx Display"
	case 4379:
		description = "CTDB"
	case 4389:
		description = "Xandros Community Management Service"
	case 4390:
		description = "Physical Access Control"
	case 4391:
		description = "American Printware IMServer Protocol"
	case 4392:
		description = "American Printware RXServer Protocol"
	case 4393:
		description = "American Printware RXSpooler Protocol"
	case 4394:
		description = "American Printware Discovery"
	case 4395:
		description = "OmniVision communication for Virtual environments"
	case 4396:
		description = "Fly Object Space"
	case 4400:
		description = "ASIGRA Services"
	case 4401:
		description = "ASIGRA Televaulting DS-System Service"
	case 4402:
		description = "ASIGRA Televaulting DS-Client Service"
	case 4403:
		description = "ASIGRA Televaulting DS-Client Monitoring/Management"
	case 4404:
		description = "ASIGRA Televaulting DS-System Monitoring/Management"
	case 4405:
		description = "ASIGRA Televaulting Message Level Restore service"
	case 4406:
		description = "ASIGRA Televaulting DS-Sleeper Service"
	case 4407:
		description = "Network Access Control Agent"
	case 4408:
		description = "SLS Technology Control Centre"
	case 4409:
		description = "Net-Cabinet comunication"
	case 4410:
		description = "RIB iTWO Application Server"
	case 4411:
		description = "Found Messaging Protocol"
	case 4412:
		description = "SmallChat"
	case 4413:
		description = "AVI Systems NMS"
	case 4414:
		description = "Updog Monitoring and Status Framework"
	case 4415:
		description = "Brocade Virtual Router Request"
	case 4416:
		description = "PJJ Media Player"
	case 4417:
		description = "Workflow Director Communication"
	case 4418:
		description = "AXYS communication protocol"
	case 4419:
		description = "Colnod Binary Protocol"
	case 4420:
		description = "NVM Express over Fabrics storage access"
	case 4421:
		description = "Multi-Platform Remote Management for Cloud Infrastructure"
	case 4422:
		description = "TSEP Installation Service Protocol"
	case 4423:
		description = "thingkit secure mesh"
	case 4425:
		description = "NetROCKEY6 SMART Plus Service"
	case 4426:
		description = "SMARTS Beacon Port"
	case 4427:
		description = "Drizzle database server"
	case 4428:
		description = "OMV-Investigation Server-Client"
	case 4429:
		description = "OMV Investigation Agent-Server"
	case 4430:
		description = "REAL SQL Server"
	case 4431:
		description = "adWISE Pipe"
	case 4432:
		description = "L-ACOUSTICS management"
	case 4433:
		description = "Versile Object Protocol"
	case 4441:
		description = "Netblox Protocol"
	case 4442:
		description = "Saris"
	case 4443:
		description = "Pharos"
	case 4444:
		description = "NV Video default"
	case 4445:
		description = "UPNOTIFYP"
	case 4446:
		description = "N1-FWP"
	case 4447:
		description = "N1-RMGMT"
	case 4448:
		description = "ASC Licence Manager"
	case 4449:
		description = "PrivateWire"
	case 4450:
		description = "Common ASCII Messaging Protocol"
	case 4451:
		description = "CTI System Msg"
	case 4452:
		description = "CTI Program Load"
	case 4453:
		description = "NSS Alert Manager"
	case 4454:
		description = "NSS Agent Manager"
	case 4455:
		description = "PR Chat User"
	case 4456:
		description = "PR Chat Server"
	case 4457:
		description = "PR Register"
	case 4458:
		description = "Matrix Configuration Protocol"
	case 4460:
		description = "Network Time Security Key Establishment"
	case 4484:
		description = "hpssmgmt service"
	case 4485:
		description = "Assyst Data Repository Service"
	case 4486:
		description = "Integrated Client Message Service"
	case 4487:
		description = "Protocol for Remote Execution over TCP"
	case 4488:
		description = "Apple Wide Area Connectivity Service ICE Bootstrap"
	case 4500:
		description = "IPsec NAT-Traversal"
	case 4502:
		description = "A25 (FAP-FGW)"
	case 4503:
		description = "M-Bus-OMS over UDP"
	case 4534:
		description = "Armagetron Advanced Game Server"
	case 4535:
		description = "Event Heap Server"
	case 4536:
		description = "Event Heap Server SSL"
	case 4537:
		description = "WSS Security Service"
	case 4538:
		description = "Software Data Exchange Gateway"
	case 4545:
		description = "WorldScores"
	case 4546:
		description = "SF License Manager (Sentinel)"
	case 4547:
		description = "Lanner License Manager"
	case 4548:
		description = "Synchromesh"
	case 4549:
		description = "Aegate PMR Service"
	case 4550:
		description = "Perman I Interbase Server"
	case 4551:
		description = "MIH Services"
	case 4552:
		description = "Men and Mice Monitoring"
	case 4553:
		description = "ICS host services"
	case 4554:
		description = "MS FRS Replication"
	case 4555:
		description = "RSIP Port"
	case 4556:
		description = "DTN Bundle CL Protocol"
	case 4557:
		description = "Marathon everRun Quorum Service Server"
	case 4558:
		description = "Marathon everRun Quorum Service Manager"
	case 4559:
		description = "HylaFAX"
	case 4563:
		description = "Amahi Anywhere"
	case 4566:
		description = "Kids Watch Time Control Service"
	case 4567:
		description = "TRAM"
	case 4568:
		description = "BMC Reporting"
	case 4569:
		description = "Inter-Asterisk eXchange"
	case 4570:
		description = "Oracle Communications Suite update"
	case 4573:
		description = "Custom backup system"
	case 4590:
		description = "RID over HTTP/TLS"
	case 4591:
		description = "HRPD L3T (AT-AN)"
	case 4592:
		description = "HRPD-ITH (AT-AN)"
	case 4593:
		description = "IPT (ANRI-ANRI)"
	case 4594:
		description = "IAS-Session (ANRI-ANRI)"
	case 4595:
		description = "IAS-Paging (ANRI-ANRI)"
	case 4596:
		description = "IAS-Neighbor (ANRI-ANRI)"
	case 4597:
		description = "A21 (AN-1xBS)"
	case 4598:
		description = "A16 (AN-AN)"
	case 4599:
		description = "A17 (AN-AN)"
	case 4600:
		description = "Piranha1"
	case 4601:
		description = "Piranha2"
	case 4602:
		description = "EAX MTS Server"
	case 4603:
		description = "Men & Mice Upgrade Agent"
	case 4604:
		description = "Identity Registration Protocol"
	case 4605:
		description = "Direct End to End Secure Chat Protocol"
	case 4606:
		description = "Secure ID to IP registration and lookup"
	case 4621:
		description = "Single port remote radio VOIP stream"
	case 4646:
		description = "DOTS Signal Channel Protocol"
	case 4658:
		description = "PlayStation2 App Port"
	case 4659:
		description = "PlayStation2 Lobby Port"
	case 4660:
		description = "smaclmgr"
	case 4661:
		description = "Kar2ouche Peer location service"
	case 4662:
		description = "OrbitNet Message Service"
	case 4663:
		description = "Note It! Message Service"
	case 4664:
		description = "Rimage Messaging Server"
	case 4665:
		description = "Container Client Message Service"
	case 4666:
		description = "E-Port Message Service"
	case 4667:
		description = "MMA Comm Services"
	case 4668:
		description = "MMA EDS Service"
	case 4669:
		description = "E-Port Data Service"
	case 4670:
		description = "Light packets transfer protocol"
	case 4671:
		description = "Bull RSF action server"
	case 4672:
		description = "remote file access server"
	case 4673:
		description = "CXWS Operations"
	case 4674:
		description = "AppIQ Agent Management"
	case 4675:
		description = "BIAP Device Status"
	case 4676:
		description = "BIAP Generic Alert"
	case 4677:
		description = "Business Continuity Servi"
	case 4678:
		description = "boundary traversal"
	case 4679:
		description = "MGE UPS Supervision"
	case 4680:
		description = "MGE UPS Management"
	case 4681:
		description = "Parliant Telephony System"
	case 4682:
		description = "finisar"
	case 4683:
		description = "Spike Clipboard Service"
	case 4684:
		description = "RFID Reader Protocol 1.0"
	case 4685:
		description = "Autopac Protocol"
	case 4686:
		description = "Manina Service Protocol"
	case 4687:
		description = "Network Scanner Tool FTP"
	case 4688:
		description = "Mobile P2P Service"
	case 4689:
		description = "Altova DatabaseCentral"
	case 4690:
		description = "Prelude IDS message proto"
	case 4691:
		description = "monotone Netsync Protocol"
	case 4692:
		description = "Conspiracy messaging"
	case 4700:
		description = "NetXMS Agent"
	case 4701:
		description = "NetXMS Management"
	case 4702:
		description = "NetXMS Server Synchronization"
	case 4703:
		description = "Network Performance Quality Evaluation"
	case 4704:
		description = "Assuria Insider"
	case 4711:
		description = "Trinity Trust Network Node Communication"
	case 4725:
		description = "TruckStar Service"
	case 4726:
		description = "A26 (FAP-FGW)"
	case 4727:
		description = "F-Link Client Information Service"
	case 4728:
		description = "CA Port Multiplexer"
	case 4729:
		description = "GSM Interface Tap"
	case 4730:
		description = "Gearman Job Queue System"
	case 4731:
		description = "Remote Capture Protocol"
	case 4732:
		description = "OHM server trigger"
	case 4733:
		description = "RES Orchestration Catalog Services"
	case 4737:
		description = "IPDR/SP"
	case 4738:
		description = "SoleraTec Locator"
	case 4739:
		description = "IP Flow Info Export"
	case 4740:
		description = "ipfix protocol over TLS"
	case 4741:
		description = "Luminizer Manager"
	case 4742:
		description = "SICCT"
	case 4743:
		description = "openhpi HPI service"
	case 4744:
		description = "Internet File Synchronization Protocol"
	case 4745:
		description = "Funambol Mobile Push"
	case 4746:
		description = "IntelliAdmin Discovery"
	case 4747:
		description = "peer-to-peer file exchange protocol"
	case 4749:
		description = "Profile for Mac"
	case 4750:
		description = "Simple Service Auto Discovery"
	case 4751:
		description = "Simple Policy Control Protocol"
	case 4752:
		description = "Simple Network Audio Protocol"
	case 4753:
		description = "Simple Invocation of Methods Over Network (SIMON)"
	case 4754:
		description = "GRE-in-UDP Encapsulation"
	case 4755:
		description = "GRE-in-UDP Encapsulation with DTLS"
	case 4756:
		description = "Reticle Decision Center"
	case 4774:
		description = "Converge RPC"
	case 4784:
		description = "BFD Multihop Control"
	case 4785:
		description = "Cisco Nexus Control Protocol"
	case 4786:
		description = "Smart Install Service"
	case 4787:
		description = "Service Insertion Architecture (SIA) Control-Plane"
	case 4788:
		description = "eXtensible Messaging Client Protocol"
	case 4789:
		description = "Virtual eXtensible Local Area Network (VXLAN)"
	case 4790:
		description = "VXLAN"
	case 4791:
		description = "IP Routable RocE"
	case 4792:
		description = "IP Routable Unified Bus"
	case 4793:
		description = "Ultra Ethernet Transport"
	case 4800:
		description = "Icona Instant Messenging System"
	case 4801:
		description = "Icona Web Embedded Chat"
	case 4802:
		description = "Icona License System Server"
	case 4803:
		description = "Notateit Messaging"
	case 4804:
		description = "AJA ntv4 Video System Discovery"
	case 4827:
		description = "HTCP"
	case 4837:
		description = "Varadero-0"
	case 4838:
		description = "Varadero-1"
	case 4839:
		description = "Varadero-2"
	case 4840:
		description = "OPC UA Connection Protocol"
	case 4841:
		description = "QUOSA Virtual Library Service"
	case 4842:
		description = "nCode ICE-flow Library AppServer"
	case 4843:
		description = "OPC UA TCP Protocol over TLS/SSL"
	case 4844:
		description = "nCode ICE-flow Library LogServer"
	case 4845:
		description = "WordCruncher Remote Library Service"
	case 4846:
		description = "Contamac ICM Service"
	case 4847:
		description = "Web Fresh Communication"
	case 4848:
		description = "App Server - Admin HTTP"
	case 4849:
		description = "App Server - Admin HTTPS"
	case 4850:
		description = "Sun App Server - NA"
	case 4851:
		description = "Apache Derby Replication"
	case 4867:
		description = "Unify Debugger"
	case 4868:
		description = "Photon Relay"
	case 4869:
		description = "Photon Relay Debug"
	case 4870:
		description = "Citcom Tracking Service"
	case 4871:
		description = "Wired"
	case 4876:
		description = "Tritium CAN Bus Bridge Service"
	case 4877:
		description = "Lighting Management Control System"
	case 4878:
		description = "Agilent Instrument Discovery"
	case 4879:
		description = "WSDL Event Receiver"
	case 4880:
		description = "IVI High-Speed LAN Instrument Protocol"
	case 4881:
		description = "SOCP Time Synchronization Protocol"
	case 4882:
		description = "SOCP Control Protocol"
	case 4883:
		description = "Meier-Phelps License Server"
	case 4884:
		description = "HiveStor Distributed File System"
	case 4885:
		description = "ABBS"
	case 4888:
		description = "xcap code analysis portal public user access"
	case 4889:
		description = "xcap code analysis portal cluster control and administration"
	case 4894:
		description = "LysKOM Protocol A"
	case 4899:
		description = "RAdmin Port"
	case 4900:
		description = "HFSQL Client/Server Database Engine"
	case 4901:
		description = "FileLocator Remote Search Agent"
	case 4902:
		description = "magicCONROL RF and Data Interface"
	case 4912:
		description = "Technicolor LUT Access Protocol"
	case 4913:
		description = "LUTher Control Protocol"
	case 4914:
		description = "Bones Remote Control"
	case 4915:
		description = "Fibics Remote Control Service"
	case 4936:
		description = "Signal protocol port for autonomic networking"
	case 4937:
		description = "ATSC-M/H Service Signaling Channel"
	case 4940:
		description = "Equitrac Office"
	case 4941:
		description = "Equitrac Office"
	case 4942:
		description = "Equitrac Office"
	case 4949:
		description = "Munin Graphing Framework"
	case 4950:
		description = "Sybase Server Monitor"
	case 4951:
		description = "PWG WIMS"
	case 4952:
		description = "SAG Directory Server"
	case 4953:
		description = "Synchronization Arbiter"
	case 4969:
		description = "CCSS QMessageMonitor"
	case 4970:
		description = "CCSS QSystemMonitor"
	case 4971:
		description = "BackUp and Restore Program"
	case 4980:
		description = "Citrix Virtual Path"
	case 4984:
		description = "WebYast"
	case 4985:
		description = "GER HC Standard"
	case 4986:
		description = "Model Railway Interface Program"
	case 4987:
		description = "SMAR Ethernet Port 1"
	case 4988:
		description = "SMAR Ethernet Port 2"
	case 4989:
		description = "Parallel for GAUSS (tm)"
	case 4990:
		description = "BusySync Calendar Synch. Protocol"
	case 4991:
		description = "VITA Radio Transport"
	case 4999:
		description = "HFSQL Client/Server Database Engine Manager"
	case 5002:
		description = "radio free ethernet"
	case 5003:
		description = "FileMaker"
	case 5004:
		description = "RTP media data"
	case 5005:
		description = "RTP control protocol"
	case 5006:
		description = "wsm server"
	case 5007:
		description = "wsm server ssl"
	case 5008:
		description = "Synapsis EDGE"
	case 5009:
		description = "Microsoft Windows Filesystem"
	case 5010:
		description = "TelepathStart"
	case 5011:
		description = "TelepathAttack"
	case 5012:
		description = "NetOnTap Service"
	case 5013:
		description = "FileMaker"
	case 5014:
		description = "Overlay Network Protocol"
	case 5015:
		description = "FileMaker"
	case 5020:
		description = "zenginkyo-1"
	case 5021:
		description = "zenginkyo-2"
	case 5022:
		description = "mice server"
	case 5023:
		description = "Htuil Server for PLD2"
	case 5024:
		description = "SCPI-TELNET"
	case 5025:
		description = "SCPI-RAW"
	case 5026:
		description = "Storix I/O daemon (data)"
	case 5027:
		description = "Storix I/O daemon (stat)"
	case 5028:
		description = "Quiqum Virtual Relais"
	case 5029:
		description = "Infobright Database Server"
	case 5031:
		description = "Direct Message Protocol"
	case 5032:
		description = "SignaCert Enterprise Trust Server Agent"
	case 5033:
		description = "Janstor Secure Data"
	case 5034:
		description = "Janstor Status"
	case 5042:
		description = "asnaacceler8db"
	case 5043:
		description = "ShopWorX Administration"
	case 5044:
		description = "LXI Event Service"
	case 5045:
		description = "Open Settlement Protocol"
	case 5046:
		description = "Vishay PM UDP Service"
	case 5047:
		description = "iSCAPE Data Broadcasting"
	case 5048:
		description = "Texai Message Service"
	case 5049:
		description = "iVocalize Web Conference"
	case 5050:
		description = "multimedia conference control tool"
	case 5051:
		description = "ITA Agent"
	case 5052:
		description = "ITA Manager"
	case 5053:
		description = "RLM Discovery Server"
	case 5054:
		description = "RLM administrative interface"
	case 5055:
		description = "UNOT"
	case 5056:
		description = "Intecom Pointspan 1"
	case 5057:
		description = "Intecom Pointspan 2"
	case 5058:
		description = "Locus Discovery"
	case 5059:
		description = "SIP Directory Services"
	case 5060:
		description = "SIP"
	case 5061:
		description = "SIP-TLS"
	case 5062:
		description = "Localisation access"
	case 5063:
		description = "centrify secure RPC"
	case 5064:
		description = "Channel Access 1"
	case 5065:
		description = "Channel Access 2"
	case 5066:
		description = "STANAG-5066-SUBNET-INTF"
	case 5067:
		description = "Authentx Service"
	case 5068:
		description = "Bitforest Data Service"
	case 5069:
		description = "I/Net 2000-NPR"
	case 5070:
		description = "VersaTrans Server Agent Service"
	case 5071:
		description = "PowerSchool"
	case 5072:
		description = "Anything In Anything"
	case 5073:
		description = "Advantage Group Port Mgr"
	case 5074:
		description = "ALES Query"
	case 5075:
		description = "Experimental Physics and Industrial Control System"
	case 5078:
		description = "PixelPusher pixel data"
	case 5079:
		description = "Cambridge Pixel SPx Reports"
	case 5080:
		description = "OnScreen Data Collection Service"
	case 5081:
		description = "SDL - Ent Trans Server"
	case 5082:
		description = "Qpur Communication Protocol"
	case 5083:
		description = "Qpur File Protocol"
	case 5084:
		description = "EPCglobal Low-Level Reader Protocol"
	case 5085:
		description = "EPCglobal Encrypted LLRP"
	case 5086:
		description = "Aprigo Collection Service"
	case 5087:
		description = "BIOTIC"
	case 5090:
		description = "Candidate AR"
	case 5091:
		description = "Context Transfer Protocol"
	case 5092:
		description = "Magpie Binary"
	case 5093:
		description = "Sentinel LM"
	case 5094:
		description = "HART-IP"
	case 5099:
		description = "SentLM Srv2Srv"
	case 5100:
		description = "Socalia service mux"
	case 5101:
		description = "Talarian"
	case 5102:
		description = "Oracle OMS non-secure"
	case 5103:
		description = "Actifio C2C"
	case 5104:
		description = "TinyMessage"
	case 5105:
		description = "Hughes Association Protocol"
	case 5106:
		description = "Actifio UDS Agent"
	case 5107:
		description = "Disk to Disk replication between Actifio Clusters"
	case 5111:
		description = "TAEP AS service"
	case 5112:
		description = "PeerMe Msg Cmd Service"
	case 5114:
		description = "Enterprise Vault Services"
	case 5115:
		description = "Symantec Autobuild Service"
	case 5116:
		description = "EPSON Projecter Image Transfer"
	case 5117:
		description = "GradeCam Image Processing"
	case 5120:
		description = "Barracuda Backup Protocol"
	case 5133:
		description = "Policy Commander"
	case 5134:
		description = "PP ActivationServer"
	case 5135:
		description = "ERP-Scale"
	case 5136:
		description = "Minotaur SA"
	case 5137:
		description = "MyCTS server port"
	case 5145:
		description = "RMONITOR SECURE"
	case 5146:
		description = "Social Alarm Service"
	case 5150:
		description = "Ascend Tunnel Management Protocol"
	case 5151:
		description = "ESRI SDE Remote Start"
	case 5152:
		description = "ESRI SDE Instance Discovery"
	case 5154:
		description = "BZFlag game server"
	case 5155:
		description = "Oracle asControl Agent"
	case 5156:
		description = "Russian Online Game"
	case 5157:
		description = "Mediat Remote Object Exchange"
	case 5161:
		description = "SNMP over SSH Transport Model"
	case 5162:
		description = "SNMP Notification over SSH Transport Model"
	case 5163:
		description = "Shadow Backup"
	case 5164:
		description = "Virtual Protocol Adapter"
	case 5165:
		description = "ife_1corp"
	case 5166:
		description = "WinPCS Service Connection"
	case 5167:
		description = "SCTE104 Connection"
	case 5168:
		description = "SCTE30 Connection"
	case 5172:
		description = "PC over IP Endpoint Management"
	case 5190:
		description = "America-Online"
	case 5191:
		description = "AmericaOnline1"
	case 5192:
		description = "AmericaOnline2"
	case 5193:
		description = "AmericaOnline3"
	case 5194:
		description = "CipherPoint Config Service"
	case 5197:
		description = "Tunstall Lone worker device interface"
	case 5200:
		description = "TARGUS GetData"
	case 5201:
		description = "TARGUS GetData 1"
	case 5202:
		description = "TARGUS GetData 2"
	case 5203:
		description = "TARGUS GetData 3"
	case 5209:
		description = "Nomad Device Video Transfer"
	case 5215:
		description = "NOTEZA Data Safety Service"
	case 5221:
		description = "3eTI Extensible Management Protocol for OAMP"
	case 5222:
		description = "XMPP Client Connection"
	case 5223:
		description = "HP Virtual Machine Group Management"
	case 5224:
		description = "HP Virtual Machine Console Operations"
	case 5225:
		description = "HP Server"
	case 5226:
		description = "HP Status"
	case 5227:
		description = "HP System Performance Metric Service"
	case 5228:
		description = "HP Virtual Room Service"
	case 5229:
		description = "Netflow/IPFIX/sFlow Collector and Forwarder Management"
	case 5230:
		description = "JaxMP RealFlow application and protocol data"
	case 5231:
		description = "Remote Control of Scan Software for Cruse Scanners"
	case 5232:
		description = "Cruse Scanning System Service"
	case 5233:
		description = "Etinnae Network File Service"
	case 5234:
		description = "EEnet communications"
	case 5235:
		description = "Galaxy Network Service"
	case 5237:
		description = "m-net discovery"
	case 5242:
		description = "ATTUne API"
	case 5243:
		description = "xyClient Status API and rendevous point"
	case 5245:
		description = "DownTools Control Protocol"
	case 5246:
		description = "CAPWAP Control Protocol"
	case 5247:
		description = "CAPWAP Data Protocol"
	case 5248:
		description = "CA Access Control Web Service"
	case 5249:
		description = "CA AC Lang Service"
	case 5250:
		description = "soaGateway"
	case 5251:
		description = "CA eTrust VM Service"
	case 5252:
		description = "Movaz SSC"
	case 5253:
		description = "Kohler Power Device Protocol"
	case 5254:
		description = "LogCabin storage service"
	case 5264:
		description = "3Com Network Jack Port 1"
	case 5265:
		description = "3Com Network Jack Port 2"
	case 5269:
		description = "XMPP Server Connection"
	case 5270:
		description = "Cartographer XMP"
	case 5271:
		description = "StageSoft CueLink messaging"
	case 5272:
		description = "PK"
	case 5280:
		description = "Bidirectional-streams Over Synchronous HTTP (BOSH)"
	case 5281:
		description = "Undo License Manager"
	case 5282:
		description = "Marimba Transmitter Port"
	case 5298:
		description = "XMPP Link-Local Messaging"
	case 5299:
		description = "NLG Data Service"
	case 5300:
		description = "HA cluster heartbeat"
	case 5301:
		description = "HA cluster general services"
	case 5302:
		description = "HA cluster configuration"
	case 5303:
		description = "HA cluster probing"
	case 5304:
		description = "HA Cluster Commands"
	case 5305:
		description = "HA Cluster Test"
	case 5306:
		description = "Sun MC Group"
	case 5307:
		description = "SCO AIP"
	case 5308:
		description = "CFengine"
	case 5309:
		description = "J Printer"
	case 5310:
		description = "Outlaws"
	case 5312:
		description = "Permabit Client-Server"
	case 5313:
		description = "Real-time & Reliable Data"
	case 5314:
		description = "opalis-rbt-ipc"
	case 5315:
		description = "HA Cluster UDP Polling"
	case 5316:
		description = "HPBladeSystem Monitor Service"
	case 5317:
		description = "HP Device Monitor Service"
	case 5318:
		description = "PKIX Certificate Management using CMS (CMC)"
	case 5320:
		description = "Webservices-based Zn interface of BSF"
	case 5321:
		description = "Webservices-based Zn interface of BSF over SSL"
	case 5343:
		description = "Sculptor Database Server"
	case 5344:
		description = "xkoto DRCP"
	case 5349:
		description = "Session Traversal Utilities for NAT (STUN) port"
	case 5350:
		description = "Port Control Protocol Multicast"
	case 5351:
		description = "Port Control Protocol"
	case 5352:
		description = "DNS Long-Lived Queries"
	case 5353:
		description = "Multicast DNS"
	case 5354:
		description = "Multicast DNS Responder IPC"
	case 5355:
		description = "LLMNR"
	case 5356:
		description = "Microsoft Small Business"
	case 5357:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 5358:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 5359:
		description = "Microsoft Alerter"
	case 5360:
		description = "Protocol for Windows SideShow"
	case 5361:
		description = "Secure Protocol for Windows SideShow"
	case 5362:
		description = "Microsoft Windows Server WSD2 Service"
	case 5363:
		description = "Windows Network Projection"
	case 5364:
		description = "Microsoft Kernel Debugger"
	case 5397:
		description = "StressTester(tm) Injector"
	case 5398:
		description = "Elektron Administration"
	case 5399:
		description = "SecurityChase"
	case 5400:
		description = "Excerpt Search"
	case 5401:
		description = "Excerpt Search Secure"
	case 5402:
		description = "OmniCast MFTP"
	case 5403:
		description = "HPOMS-CI-LSTN"
	case 5404:
		description = "HPOMS-DPS-LSTN"
	case 5405:
		description = "NetSupport"
	case 5406:
		description = "Systemics Sox"
	case 5407:
		description = "Foresyte-Clear"
	case 5408:
		description = "Foresyte-Sec"
	case 5409:
		description = "Salient Data Server"
	case 5410:
		description = "Salient User Manager"
	case 5411:
		description = "ActNet"
	case 5412:
		description = "Continuus"
	case 5413:
		description = "WWIOTALK"
	case 5414:
		description = "StatusD"
	case 5415:
		description = "NS Server"
	case 5416:
		description = "SNS Gateway"
	case 5417:
		description = "SNS Agent"
	case 5418:
		description = "MCNTP"
	case 5419:
		description = "DJ-ICE"
	case 5420:
		description = "Cylink-C"
	case 5421:
		description = "Net Support 2"
	case 5422:
		description = "Salient MUX"
	case 5423:
		description = "VIRTUALUSER"
	case 5424:
		description = "Beyond Remote"
	case 5425:
		description = "Beyond Remote Command Channel"
	case 5426:
		description = "DEVBASIC"
	case 5427:
		description = "SCO-PEER-TTA"
	case 5428:
		description = "TELACONSOLE"
	case 5429:
		description = "Billing and Accounting System Exchange"
	case 5430:
		description = "RADEC CORP"
	case 5431:
		description = "PARK AGENT"
	case 5432:
		description = "PostgreSQL Database"
	case 5433:
		description = "Pyrrho DBMS"
	case 5434:
		description = "SGI Array Services Daemon"
	case 5435:
		description = "SCEANICS situation and action notification"
	case 5436:
		description = "pmip6-cntl"
	case 5437:
		description = "pmip6-data"
	case 5443:
		description = "Pearson HTTPS"
	case 5445:
		description = "Server Message Block over Remote Direct Memory Access"
	case 5450:
		description = "TiePie engineering data acquisition"
	case 5453:
		description = "SureBox"
	case 5454:
		description = "APC 5454"
	case 5455:
		description = "APC 5455"
	case 5456:
		description = "APC 5456"
	case 5461:
		description = "SILKMETER"
	case 5462:
		description = "TTL Publisher"
	case 5463:
		description = "TTL Price Proxy"
	case 5464:
		description = "Quail Networks Object Broker"
	case 5465:
		description = "NETOPS-BROKER"
	case 5470:
		description = "The Apsolab company's data collection protocol"
	case 5471:
		description = "The Apsolab company's secure data collection protocol"
	case 5472:
		description = "The Apsolab company's dynamic tag protocol"
	case 5473:
		description = "The Apsolab company's secure dynamic tag protocol"
	case 5474:
		description = "The Apsolab company's status query protocol"
	case 5475:
		description = "The Apsolab company's data retrieval protocol"
	case 5500:
		description = "fcp-addr-srvr1"
	case 5501:
		description = "fcp-addr-srvr2"
	case 5502:
		description = "fcp-srvr-inst1"
	case 5503:
		description = "fcp-srvr-inst2"
	case 5504:
		description = "fcp-cics-gw1"
	case 5505:
		description = "Checkout Database"
	case 5506:
		description = "Amcom Mobile Connect"
	case 5507:
		description = "PowerSysLab Electrical Management"
	case 5540:
		description = "Matter Operational Discovery and Communi"
	case 5543:
		description = "QF-Test License Server"
	case 5550:
		description = "Model Railway control using the CBUS message protocol"
	case 5553:
		description = "SGI Eventmond Port"
	case 5554:
		description = "SGI ESP HTTP"
	case 5555:
		description = "Personal Agent"
	case 5556:
		description = "Freeciv gameplay"
	case 5557:
		description = "Sandlab FARENET"
	case 5565:
		description = "Data Protector BURA"
	case 5566:
		description = "Westec Connect"
	case 5567:
		description = "DOF Protocol Stack Multicast/Secure Transport"
	case 5568:
		description = "Session Data Transport Multicast"
	case 5569:
		description = "PLASA E1.33"
	case 5573:
		description = "SAS Domain Management Messaging Protocol"
	case 5574:
		description = "SAS IO Forwarding"
	case 5575:
		description = "Oracle Access Protocol"
	case 5579:
		description = "FleetDisplay Tracking Service"
	case 5580:
		description = "T-Mobile SMS Protocol Message 0"
	case 5581:
		description = "T-Mobile SMS Protocol Message 1"
	case 5582:
		description = "T-Mobile SMS Protocol Message 3"
	case 5583:
		description = "T-Mobile SMS Protocol Message 2"
	case 5584:
		description = "BeInSync-Web"
	case 5585:
		description = "BeInSync-sync"
	case 5597:
		description = "inin secure messaging"
	case 5598:
		description = "MCT Market Data Feed"
	case 5599:
		description = "Enterprise Security Remote Install"
	case 5600:
		description = "Enterprise Security Manager"
	case 5601:
		description = "Enterprise Security Agent"
	case 5602:
		description = "A1-MSC"
	case 5603:
		description = "A1-BS"
	case 5604:
		description = "A3-SDUNode"
	case 5605:
		description = "A4-SDUNode"
	case 5618:
		description = "Fiscal Registering Protocol"
	case 5627:
		description = "Node Initiated Network Association Forma"
	case 5628:
		description = "HTrust API"
	case 5629:
		description = "Symantec Storage Foundation for Database"
	case 5630:
		description = "PreciseCommunication"
	case 5631:
		description = "pcANYWHEREdata"
	case 5632:
		description = "pcANYWHEREstat"
	case 5633:
		description = "BE Operations Request Listener"
	case 5634:
		description = "SF Message Service"
	case 5635:
		description = "SFM Authentication Subsystem"
	case 5636:
		description = "SFMdb - SFM DB server"
	case 5637:
		description = "Symantec CSSC"
	case 5638:
		description = "Symantec Fingerprint Lookup and Container Reference Service"
	case 5639:
		description = "Symantec Integrity Checking Service"
	case 5646:
		description = "Ventureforth Mobile"
	case 5666:
		description = "Nagios Remote Plugin Executor"
	case 5670:
		description = "ZeroMQ file publish-subscribe protocol"
	case 5671:
		description = "amqp protocol over TLS/SSL"
	case 5672:
		description = "AMQP"
	case 5673:
		description = "JACL Message Server"
	case 5674:
		description = "HyperSCSI Port"
	case 5675:
		description = "V5UA application port"
	case 5676:
		description = "RA Administration"
	case 5677:
		description = "Quest Central DB2 Launchr"
	case 5678:
		description = "Remote Replication Agent Connection"
	case 5679:
		description = "Direct Cable Connect Manager"
	case 5680:
		description = "Auriga Router Service"
	case 5681:
		description = "Net-coneX Control Protocol"
	case 5682:
		description = "BrightCore control & data transfer exchange"
	case 5683:
		description = "Constrained Application Protocol (CoAP)"
	case 5684:
		description = "Constrained Application Protocol (CoAP)"
	case 5687:
		description = "GOG multiplayer game protocol"
	case 5688:
		description = "GGZ Gaming Zone"
	case 5689:
		description = "QM video network management protocol"
	case 5693:
		description = "Robert Bosch Data Transfer"
	case 5696:
		description = "Key Management Interoperability Protocol"
	case 5700:
		description = "Dell SupportAssist data center management"
	case 5705:
		description = "StorageOS REST API"
	case 5713:
		description = "proshare conf audio"
	case 5714:
		description = "proshare conf video"
	case 5715:
		description = "proshare conf data"
	case 5716:
		description = "proshare conf request"
	case 5717:
		description = "proshare conf notify"
	case 5718:
		description = "DPM Communication Server"
	case 5719:
		description = "DPM Agent Coordinator"
	case 5720:
		description = "MS-Licensing"
	case 5721:
		description = "Desktop Passthru Service"
	case 5722:
		description = "Microsoft DFS Replication Service"
	case 5723:
		description = "Operations Manager - Health Service"
	case 5724:
		description = "Operations Manager - SDK Service"
	case 5725:
		description = "Microsoft Identity Lifecycle Manager"
	case 5726:
		description = "Microsoft Lifecycle Manager Secure Token Service"
	case 5727:
		description = "ASG Event Notification Framework"
	case 5728:
		description = "Dist. I/O Comm. Service Data and Control"
	case 5729:
		description = "Openmail User Agent Layer"
	case 5730:
		description = "Steltor's calendar access"
	case 5741:
		description = "IDA Discover Port 1"
	case 5742:
		description = "IDA Discover Port 2"
	case 5743:
		description = "Watchdoc NetPOD Protocol"
	case 5744:
		description = "Watchdoc Server"
	case 5745:
		description = "fcopy-server"
	case 5746:
		description = "fcopys-server"
	case 5747:
		description = "Wildbits Tunatic"
	case 5748:
		description = "Wildbits Tunalyzer"
	case 5750:
		description = "Bladelogic Agent Service"
	case 5755:
		description = "OpenMail Desk Gateway server"
	case 5757:
		description = "OpenMail X.500 Directory Server"
	case 5766:
		description = "OpenMail NewMail Server"
	case 5767:
		description = "OpenMail Suer Agent Layer (Secure)"
	case 5768:
		description = "OpenMail CMTS Server"
	case 5769:
		description = "x509solutions Internal CA"
	case 5770:
		description = "x509solutions Secure Data"
	case 5771:
		description = "NetAgent"
	case 5777:
		description = "Control commands and responses"
	case 5780:
		description = "Visual Tag System RPC"
	case 5781:
		description = "3PAR Event Reporting Service"
	case 5782:
		description = "3PAR Management Service"
	case 5783:
		description = "3PAR Management Service with SSL"
	case 5784:
		description = "Cisco Interbox Application Redundancy"
	case 5785:
		description = "3PAR Inform Remote Copy"
	case 5786:
		description = "redundancy notification"
	case 5787:
		description = "Cisco WAAS Cluster Protocol"
	case 5793:
		description = "XtreamX Supervised Peer message"
	case 5794:
		description = "Simple Peered Discovery Protocol"
	case 5798:
		description = "Proprietary Website deployment service"
	case 5813:
		description = "ICMPD"
	case 5814:
		description = "Support Automation"
	case 5820:
		description = "AutoPass licensing"
	case 5841:
		description = "Z-firm ShipRush interface for web access and bidirectional data"
	case 5842:
		description = "Reversion Backup/Restore"
	case 5859:
		description = "WHEREHOO"
	case 5863:
		description = "PlanetPress Suite Messeng"
	case 5868:
		description = "Diameter over TLS/TCP"
	case 5883:
		description = "Javascript Unit Test Environment"
	case 5900:
		description = "Remote Framebuffer"
	case 5903:
		description = "Flight & Flow Info for Collaborative Env"
	case 5904:
		description = "Air-Ground SWIM"
	case 5905:
		description = "Adv Surface Mvmnt and Guidance Cont Sys"
	case 5906:
		description = "Remotely Piloted Vehicle C&C"
	case 5907:
		description = "Distress and Safety Data App"
	case 5908:
		description = "IPS Management Application"
	case 5909:
		description = "Air-ground media advisory"
	case 5910:
		description = "Air Traffic Services applications using ATN"
	case 5911:
		description = "Air Traffic Services applications using ACARS"
	case 5912:
		description = "Applications using ACARS"
	case 5913:
		description = "Applications using ACARS"
	case 5914:
		description = "Security for Internet Protocol Suite"
	case 5963:
		description = "Indy Application Server"
	case 5968:
		description = "mppolicy-v5"
	case 5969:
		description = "mppolicy-mgr"
	case 5984:
		description = "CouchDB"
	case 5985:
		description = "WBEM WS-Management HTTP"
	case 5986:
		description = "WBEM WS-Management HTTP over TLS/SSL"
	case 5987:
		description = "WBEM RMI"
	case 5988:
		description = "WBEM CIM-XML (HTTP)"
	case 5989:
		description = "WBEM CIM-XML (HTTPS)"
	case 5990:
		description = "WBEM Export HTTPS"
	case 5991:
		description = "NUXSL"
	case 5992:
		description = "Consul InSight Security"
	case 5993:
		description = "DMTF WBEM CIM REST"
	case 5994:
		description = "RMS Agent Listening Service"
	case 5999:
		description = "CVSup"
	case 6064:
		description = "NDL-AHP-SVC"
	case 6065:
		description = "WinPharaoh"
	case 6066:
		description = "EWCTSP"
	case 6068:
		description = "GSMP/ANCP"
	case 6069:
		description = "TRIP"
	case 6070:
		description = "Messageasap"
	case 6071:
		description = "SSDTP"
	case 6072:
		description = "DIAGNOSE-PROC"
	case 6073:
		description = "DirectPlay8"
	case 6074:
		description = "Microsoft Max"
	case 6075:
		description = "Microsoft DPM Access Control Manager"
	case 6076:
		description = "Microsoft DPM WCF Certificates"
	case 6077:
		description = "iConstruct Server"
	case 6080:
		description = "Generic UDP Encapsulation"
	case 6081:
		description = "Generic Network Virtualization Encapsulation (Geneve)"
	case 6082:
		description = "APCO Project 25 Common Air Interface"
	case 6083:
		description = "telecomsoftware miami broadcast"
	case 6084:
		description = "Peer to Peer Infrastructure Configuration"
	case 6085:
		description = "konspire2b p2p network"
	case 6086:
		description = "PDTP P2P"
	case 6087:
		description = "Local Download Sharing Service"
	case 6088:
		description = "SuperDog License Manager"
	case 6099:
		description = "RAXA Management"
	case 6100:
		description = "SynchroNet-db"
	case 6101:
		description = "SynchroNet-rtc"
	case 6102:
		description = "SynchroNet-upd"
	case 6103:
		description = "RETS"
	case 6104:
		description = "DBDB"
	case 6105:
		description = "Prima Server"
	case 6106:
		description = "MPS Server"
	case 6107:
		description = "ETC Control"
	case 6108:
		description = "Sercomm-SCAdmin"
	case 6109:
		description = "GLOBECAST-ID"
	case 6110:
		description = "HP SoftBench CM"
	case 6111:
		description = "HP SoftBench Sub-Process Control"
	case 6112:
		description = "Desk-Top Sub-Process Control Daemon"
	case 6113:
		description = "Daylite Server"
	case 6114:
		description = "WRspice IPC Service"
	case 6115:
		description = "Xic IPC Service"
	case 6116:
		description = "XicTools License Manager Service"
	case 6117:
		description = "Daylite Touch Sync"
	case 6118:
		description = "Transparent Inter Process Communication"
	case 6121:
		description = "SPDY for a faster web"
	case 6122:
		description = "Backup Express Web Server"
	case 6123:
		description = "Backup Express"
	case 6124:
		description = "Phlexible Network Backup Service"
	case 6130:
		description = "The DameWare Mobile Gateway Service"
	case 6133:
		description = "New Boundary Tech WOL"
	case 6140:
		description = "Pulsonix Network License Service"
	case 6141:
		description = "Meta Corporation License Manager"
	case 6142:
		description = "Aspen Technology License Manager"
	case 6143:
		description = "Watershed License Manager"
	case 6144:
		description = "StatSci License Manager - 1"
	case 6145:
		description = "StatSci License Manager - 2"
	case 6146:
		description = "Lone Wolf Systems License Manager"
	case 6147:
		description = "Montage License Manager"
	case 6148:
		description = "Ricardo North America License Manager"
	case 6149:
		description = "tal-pod"
	case 6159:
		description = "EFB Application Control Interface"
	case 6160:
		description = "Emerson Extensible Control and Management Protocol"
	case 6161:
		description = "PATROL Internet Srv Mgr"
	case 6162:
		description = "PATROL Collector"
	case 6163:
		description = "Precision Scribe Cnx Port"
	case 6200:
		description = "LM-X License Manager by X-Formation"
	case 6201:
		description = "Management of service nodes in a processing grid for thermodynamic calculations"
	case 6209:
		description = "QMTP over TLS"
	case 6222:
		description = "Radmind Access Protocol"
	case 6241:
		description = "JEOL Network Services Data Transport Protocol 1"
	case 6242:
		description = "JEOL Network Services Data Transport Protocol 2"
	case 6243:
		description = "JEOL Network Services Data Transport Protocol 3"
	case 6244:
		description = "JEOL Network Services Data Transport Protocol 4"
	case 6251:
		description = "TL1 Raw Over SSL/TLS"
	case 6252:
		description = "TL1 over SSH"
	case 6253:
		description = "CRIP"
	case 6267:
		description = "GridLAB-D User Interface"
	case 6268:
		description = "Grid Authentication"
	case 6269:
		description = "Grid Authentication Alt"
	case 6300:
		description = "BMC GRX"
	case 6301:
		description = "BMC CONTROL-D LDAP SERVER"
	case 6306:
		description = "Unified Fabric Management Protocol"
	case 6315:
		description = "Sensor Control Unit Protocol"
	case 6316:
		description = "Ethernet Sensor Communications Protocol"
	case 6317:
		description = "Navtech Radar Sensor Data Command"
	case 6318:
		description = "IONA Measurement and control data"
	case 6320:
		description = "Double-Take Replication Service"
	case 6321:
		description = "Empress Software Connectivity Server 1"
	case 6322:
		description = "Empress Software Connectivity Server 2"
	case 6324:
		description = "HR Device Network service"
	case 6325:
		description = "Double-Take Management Service"
	case 6326:
		description = "Double-Take Virtual Recovery Assistant"
	case 6343:
		description = "sFlow traffic monitoring"
	case 6344:
		description = "Argus-Spectr security and fire-prevention systems service"
	case 6346:
		description = "gnutella-svc"
	case 6347:
		description = "gnutella-rtr"
	case 6350:
		description = "App Discovery and Access Protocol"
	case 6355:
		description = "PMCS applications"
	case 6360:
		description = "MetaEdit+ Multi-User"
	case 6363:
		description = "Named Data Networking"
	case 6370:
		description = "MetaEdit+ Server Administration"
	case 6379:
		description = "An advanced key-value cache and store"
	case 6382:
		description = "Metatude Dialogue Server"
	case 6389:
		description = "clariion-evr01"
	case 6390:
		description = "MetaEdit+ WebService API"
	case 6400:
		description = "Business Objects CMS contact port"
	case 6401:
		description = "boe-was"
	case 6402:
		description = "boe-eventsrv"
	case 6403:
		description = "boe-cachesvr"
	case 6404:
		description = "Business Objects Enterprise internal server"
	case 6405:
		description = "Business Objects Enterprise internal server"
	case 6406:
		description = "Business Objects Enterprise internal server"
	case 6407:
		description = "Business Objects Enterprise internal server"
	case 6408:
		description = "Business Objects Enterprise internal server"
	case 6409:
		description = "Business Objects Enterprise internal server"
	case 6410:
		description = "Business Objects Enterprise internal server"
	case 6417:
		description = "Faxcom Message Service"
	case 6418:
		description = "SYserver remote commands"
	case 6419:
		description = "Simple VDR Protocol"
	case 6420:
		description = "NIM_VDRShell"
	case 6421:
		description = "NIM_WAN"
	case 6432:
		description = "PgBouncer"
	case 6440:
		description = "heliosd daemon"
	case 6442:
		description = "Transitory Application Request Protocol"
	case 6443:
		description = "Service Registry Default HTTPS Domain"
	case 6444:
		description = "Grid Engine Qmaster Service"
	case 6445:
		description = "Grid Engine Execution Service"
	case 6446:
		description = "MySQL Proxy"
	case 6455:
		description = "SKIP Certificate Receive"
	case 6456:
		description = "SKIP Certificate Send"
	case 6464:
		description = "Medical devices"
	case 6471:
		description = "LVision License Manager"
	case 6480:
		description = "Service Registry Default HTTP Domain"
	case 6481:
		description = "Service Tags"
	case 6482:
		description = "Logical Domains Management Interface"
	case 6483:
		description = "SunVTS RMI"
	case 6484:
		description = "Service Registry Default JMS Domain"
	case 6485:
		description = "Service Registry Default IIOP Domain"
	case 6486:
		description = "Service Registry Default IIOPS Domain"
	case 6487:
		description = "Service Registry Default IIOPAuth Domain"
	case 6488:
		description = "Service Registry Default JMX Domain"
	case 6489:
		description = "Service Registry Default Admin Domain"
	case 6500:
		description = "BoKS Master"
	case 6501:
		description = "BoKS Servc"
	case 6502:
		description = "BoKS Servm"
	case 6503:
		description = "BoKS Clntd"
	case 6505:
		description = "BoKS Admin Private Port"
	case 6506:
		description = "BoKS Admin Public Port"
	case 6507:
		description = "BoKS Dir Server"
	case 6508:
		description = "BoKS Dir Server"
	case 6509:
		description = "MGCS-MFP Port"
	case 6510:
		description = "MCER Port"
	case 6511:
		description = "Datagram Congestion Control Protocol Encapsulation for NAT Traversal"
	case 6513:
		description = "NETCONF over TLS"
	case 6514:
		description = "Syslog over TLS"
	case 6515:
		description = "Elipse RPC Protocol"
	case 6543:
		description = "lds_distrib"
	case 6544:
		description = "LDS Dump Service"
	case 6547:
		description = "APC 6547"
	case 6548:
		description = "APC 6548"
	case 6549:
		description = "APC 6549"
	case 6550:
		description = "fg-sysupdate"
	case 6551:
		description = "Software Update Manager"
	case 6556:
		description = "Checkmk Monitoring Agent"
	case 6566:
		description = "SANE Control Port"
	case 6568:
		description = "Roaring Penguin IP Address Reputation Collection"
	case 6579:
		description = "Affiliate"
	case 6580:
		description = "Parsec Masterserver"
	case 6581:
		description = "Parsec Peer-to-Peer"
	case 6582:
		description = "Parsec Gameserver"
	case 6583:
		description = "JOA Jewel Suite"
	case 6600:
		description = "Microsoft Hyper-V Live Migration"
	case 6601:
		description = "Microsoft Threat Management Gateway SSTP"
	case 6602:
		description = "Windows WSS Communication Framework"
	case 6619:
		description = "ODETTE-FTP over TLS/SSL"
	case 6620:
		description = "Kerberos V5 FTP Data"
	case 6621:
		description = "Kerberos V5 FTP Control"
	case 6622:
		description = "Multicast FTP"
	case 6623:
		description = "Kerberos V5 Telnet"
	case 6624:
		description = "DataScaler database"
	case 6625:
		description = "DataScaler control"
	case 6626:
		description = "WAGO Service and Update"
	case 6627:
		description = "Allied Electronics NeXGen"
	case 6628:
		description = "AFE Stock Channel M/C"
	case 6632:
		description = "eGenix mxODBC Connect"
	case 6633:
		description = "Cisco vPath Services Overlay"
	case 6634:
		description = "MPLS Performance Measurement out-of-band response"
	case 6635:
		description = "Encapsulate MPLS packets in UDP tunnels."
	case 6636:
		description = "Encapsulate MPLS packets in UDP tunnels with DTLS."
	case 6640:
		description = "Open vSwitch Database protocol"
	case 6653:
		description = "OpenFlow"
	case 6655:
		description = "PC SOFT - Software factory UI/manager"
	case 6656:
		description = "Emergency Message Control Service"
	case 6657:
		description = "PalCom Discovery"
	case 6670:
		description = "Vocaltec Global Online Directory"
	case 6671:
		description = "P4P Portal Service"
	case 6672:
		description = "vision_server"
	case 6673:
		description = "vision_elmd"
	case 6678:
		description = "Viscount Freedom Bridge Protocol"
	case 6679:
		description = "Osorno Automation"
	case 6687:
		description = "CleverView for cTrace Message Service"
	case 6688:
		description = "CleverView for TCP/IP Message Service"
	case 6689:
		description = "Tofino Security Appliance"
	case 6690:
		description = "CLEVERDetect Message Service"
	case 6696:
		description = "Babel Routing Protocol"
	case 6697:
		description = "Internet Relay Chat via TLS/SSL"
	case 6699:
		description = "Babel Routing Protocol over DTLS"
	case 6701:
		description = "KTI/ICAD Nameserver"
	case 6702:
		description = "e-Design network"
	case 6703:
		description = "e-Design web"
	case 6704:
		description = "ForCES HP (High Priority) channel"
	case 6705:
		description = "ForCES MP (Medium Priority) channel"
	case 6706:
		description = "ForCES LP (Low priority) channel"
	case 6714:
		description = "Internet Backplane Protocol"
	case 6715:
		description = "Fibotrader Communications"
	case 6716:
		description = "Princity Agent"
	case 6767:
		description = "BMC PERFORM AGENT"
	case 6768:
		description = "BMC PERFORM MGRD"
	case 6769:
		description = "ADInstruments GxP Server"
	case 6770:
		description = "PolyServe http"
	case 6771:
		description = "PolyServe https"
	case 6777:
		description = "netTsunami Tracker"
	case 6778:
		description = "netTsunami p2p storage system"
	case 6784:
		description = "Bidirectional Forwarding Detection on Link Aggregation Group"
	case 6785:
		description = "DGPF Individual Exchange"
	case 6786:
		description = "Sun Java Web Console JMX"
	case 6787:
		description = "Sun Web Console Admin"
	case 6788:
		description = "SMC-HTTP"
	case 6789:
		description = "GSS-API for the Oracle Remote Administration Daemon"
	case 6790:
		description = "HNMP"
	case 6791:
		description = "Halcyon Network Manager"
	case 6801:
		description = "ACNET Control System Protocol"
	case 6817:
		description = "PenTBox Secure IM Protocol"
	case 6831:
		description = "ambit-lm"
	case 6841:
		description = "Netmo Default"
	case 6842:
		description = "Netmo HTTP"
	case 6850:
		description = "ICCRUSHMORE"
	case 6868:
		description = "Acctopus Command Channel"
	case 6888:
		description = "MUSE"
	case 6900:
		description = "R*TIME Viewer Data Interface"
	case 6901:
		description = "Novell Jetstream messaging protocol"
	case 6924:
		description = "Ping with RX/TX latency/loss split"
	case 6935:
		description = "EthoScan Service"
	case 6936:
		description = "XenSource Management Service"
	case 6946:
		description = "Biometrics Server"
	case 6951:
		description = "OTLP"
	case 6961:
		description = "JMACT3"
	case 6962:
		description = "jmevt2"
	case 6963:
		description = "swismgr1"
	case 6964:
		description = "swismgr2"
	case 6965:
		description = "swistrap"
	case 6966:
		description = "swispol"
	case 6969:
		description = "acmsoda"
	case 6970:
		description = "Conductor test coordination protocol"
	case 6980:
		description = "QoS-extended OLSR protocol"
	case 6997:
		description = "Mobility XE Protocol"
	case 6998:
		description = "IATP-highPri"
	case 6999:
		description = "IATP-normalPri"
	case 7000:
		description = "file server itself"
	case 7001:
		description = "callbacks to cache managers"
	case 7002:
		description = "users & groups database"
	case 7003:
		description = "volume location database"
	case 7004:
		description = "AFS/Kerberos authentication service"
	case 7005:
		description = "volume managment server"
	case 7006:
		description = "error interpretation service"
	case 7007:
		description = "basic overseer process"
	case 7008:
		description = "server-to-server updater"
	case 7009:
		description = "remote cache manager service"
	case 7010:
		description = "onlinet uninterruptable power supplies"
	case 7011:
		description = "Talon Discovery Port"
	case 7012:
		description = "Talon Engine"
	case 7013:
		description = "Microtalon Discovery"
	case 7014:
		description = "Microtalon Communications"
	case 7015:
		description = "Talon Webserver"
	case 7016:
		description = "SPG Controls Carrier"
	case 7017:
		description = "GeneRic Autonomic Signaling Protocol"
	case 7018:
		description = "FISA Service"
	case 7019:
		description = "doceri drawing service control"
	case 7020:
		description = "DP Serve"
	case 7021:
		description = "DP Serve Admin"
	case 7022:
		description = "CT Discovery Protocol"
	case 7023:
		description = "Comtech T2 NMCS"
	case 7024:
		description = "Vormetric service"
	case 7025:
		description = "Vormetric Service II"
	case 7026:
		description = "Loreji Webhosting Panel"
	case 7030:
		description = "ObjectPlanet probe"
	case 7031:
		description = "IPOSPLANET retailing multi devices protocol"
	case 7040:
		description = "Quest application level network service discovery"
	case 7070:
		description = "ARCP"
	case 7071:
		description = "IWGADTS Aircraft Housekeeping Message"
	case 7072:
		description = "iba Device Configuration Protocol"
	case 7073:
		description = "MarTalk protocol"
	case 7080:
		description = "EmpowerID Communication"
	case 7088:
		description = "Zixi live video transport protocol"
	case 7095:
		description = "Java Discovery Protocol"
	case 7099:
		description = "lazy-ptop"
	case 7100:
		description = "X Font Service"
	case 7101:
		description = "Embedded Light Control Network"
	case 7107:
		description = "AES-X170"
	case 7117:
		description = "Encrypted chat and file transfer service"
	case 7121:
		description = "Virtual Prototypes License Manager"
	case 7123:
		description = "End-to-end TLS Relay Control Connection"
	case 7128:
		description = "intelligent data manager"
	case 7129:
		description = "Catalog Content Search"
	case 7161:
		description = "CA BSM Comm"
	case 7162:
		description = "CA Storage Manager"
	case 7163:
		description = "CA Connection Broker"
	case 7164:
		description = "File System Repository Agent"
	case 7165:
		description = "Document WCF Server"
	case 7166:
		description = "Aruba eDiscovery Server"
	case 7167:
		description = "CA SRM Agent"
	case 7168:
		description = "cncKadServer DB & Inventory Services"
	case 7169:
		description = "Consequor Consulting Process Integration Bridge"
	case 7170:
		description = "Adaptive Name/Service Resolution"
	case 7171:
		description = "Discovery and Retention Mgt Production"
	case 7172:
		description = "Port used for MetalBend programmable interface"
	case 7173:
		description = "zSecure Server"
	case 7174:
		description = "Clutild"
	case 7181:
		description = "Janus Guidewire Enterprise Discovery Service Bus"
	case 7200:
		description = "FODMS FLIP"
	case 7201:
		description = "DLIP"
	case 7202:
		description = "ICTP"
	case 7215:
		description = "Communication ports for PaperStream Server services"
	case 7216:
		description = "PaperStream Capture Professional"
	case 7227:
		description = "Registry A & M Protocol"
	case 7228:
		description = "Citrix Universal Printing Port"
	case 7229:
		description = "Citrix UPP Gateway"
	case 7234:
		description = "Traffic forwarding for Okta cloud infra"
	case 7235:
		description = "ASP Coordination Protocol"
	case 7236:
		description = "Wi-Fi Alliance Wi-Fi Display Protocol"
	case 7237:
		description = "PADS (Public Area Display System) Server"
	case 7244:
		description = "FrontRow Calypso Human Interface Control Protocol"
	case 7262:
		description = "Calypso Network Access Protocol"
	case 7272:
		description = "WatchMe Monitoring 7272"
	case 7273:
		description = "OMA Roaming Location"
	case 7274:
		description = "OMA Roaming Location SEC"
	case 7275:
		description = "OMA UserPlane Location"
	case 7276:
		description = "OMA Internal Location Protocol"
	case 7277:
		description = "OMA Internal Location Secure Protocol"
	case 7278:
		description = "OMA Dynamic Content Delivery over CBS"
	case 7279:
		description = "Citrix Licensing"
	case 7280:
		description = "ITACTIONSERVER 1"
	case 7281:
		description = "ITACTIONSERVER 2"
	case 7282:
		description = "eventACTION/ussACTION (MZCA) server"
	case 7283:
		description = "General Statistics Rendezvous Protocol"
	case 7365:
		description = "LifeKeeper Communications"
	case 7391:
		description = "mind-file system server"
	case 7392:
		description = "mrss-rendezvous server"
	case 7393:
		description = "nFoldMan Remote Publish"
	case 7394:
		description = "File system export of backup images"
	case 7395:
		description = "winqedit"
	case 7397:
		description = "Hexarc Command Language"
	case 7400:
		description = "RTPS Discovery"
	case 7401:
		description = "RTPS Data-Distribution User-Traffic"
	case 7402:
		description = "RTPS Data-Distribution Meta-Traffic"
	case 7410:
		description = "Ionix Network Monitor"
	case 7411:
		description = "Streaming of measurement data"
	case 7420:
		description = "Multichannel real-time lighting control"
	case 7421:
		description = "Matisse Port Monitor"
	case 7426:
		description = "OpenView DM Postmaster Manager"
	case 7427:
		description = "OpenView DM Event Agent Manager"
	case 7428:
		description = "OpenView DM Log Agent Manager"
	case 7429:
		description = "OpenView DM rqt communication"
	case 7430:
		description = "OpenView DM xmpv7 api pipe"
	case 7431:
		description = "OpenView DM ovc/xmpv3 api pipe"
	case 7437:
		description = "Faximum"
	case 7443:
		description = "Oracle Application Server HTTPS"
	case 7471:
		description = "Stateless Transport Tunneling Protocol"
	case 7473:
		description = "Rise: The Vieneo Province"
	case 7474:
		description = "Neo4j Graph Database"
	case 7478:
		description = "IT Asset Management"
	case 7491:
		description = "telops-lmd"
	case 7500:
		description = "Silhouette User"
	case 7501:
		description = "HP OpenView Bus Daemon"
	case 7508:
		description = "Automation Device Configuration Protocol"
	case 7509:
		description = "ACPLT - process automation service"
	case 7510:
		description = "HP OpenView Application Server"
	case 7511:
		description = "pafec-lm"
	case 7542:
		description = "Saratoga Transfer Protocol"
	case 7543:
		description = "atul server"
	case 7544:
		description = "FlowAnalyzer DisplayServer"
	case 7545:
		description = "FlowAnalyzer UtilityServer"
	case 7546:
		description = "Cisco Fabric service"
	case 7547:
		description = "Broadband Forum CWMP"
	case 7548:
		description = "Threat Information Distribution Protocol"
	case 7549:
		description = "Network Layer Signaling Transport Layer"
	case 7550:
		description = "Cloud Signaling Service"
	case 7551:
		description = "ControlONE Console signaling"
	case 7560:
		description = "Sniffer Command Protocol"
	case 7563:
		description = "Control Framework"
	case 7566:
		description = "VSI Omega"
	case 7569:
		description = "Dell EqualLogic Host Group Management"
	case 7570:
		description = "Aries Kfinder"
	case 7574:
		description = "Oracle Coherence Cluster Service"
	case 7575:
		description = "Main access port for WTMI Panel"
	case 7588:
		description = "Sun License Manager"
	case 7606:
		description = "MIPI Alliance Debug"
	case 7624:
		description = "Instrument Neutral Distributed Interface"
	case 7626:
		description = "SImple Middlebox COnfiguration (SIMCO)"
	case 7627:
		description = "SOAP Service Port"
	case 7628:
		description = "Primary Agent Work Notification"
	case 7629:
		description = "OpenXDAS Wire Protocol"
	case 7630:
		description = "HA Web Konsole"
	case 7631:
		description = "TESLA System Messaging"
	case 7633:
		description = "PMDF Management"
	case 7648:
		description = "bonjour-cuseeme"
	case 7663:
		description = "Proprietary immutable distributed data storage"
	case 7668:
		description = "AuthorityGate Secure API communication for Orchestrated task sequencing"
	case 7672:
		description = "iMQ STOMP Server"
	case 7673:
		description = "iMQ STOMP Server over SSL"
	case 7674:
		description = "iMQ SSL tunnel"
	case 7675:
		description = "iMQ Tunnel"
	case 7676:
		description = "iMQ Broker Rendezvous"
	case 7677:
		description = "Sun App Server - HTTPS"
	case 7680:
		description = "Microsoft Delivery Optimization Peer-to-Peer"
	case 7683:
		description = "Cleondris DMT"
	case 7687:
		description = "Bolt database connection"
	case 7689:
		description = "Collaber Network Service"
	case 7690:
		description = "Service-Oriented Vehicle Diagnostics"
	case 7697:
		description = "KLIO communications"
	case 7700:
		description = "EM7 Secure Communications"
	case 7701:
		description = "SCF nFAPI defining MAC/PHY split"
	case 7707:
		description = "EM7 Dynamic Updates"
	case 7708:
		description = "scientia.net"
	case 7720:
		description = "MedImage Portal"
	case 7724:
		description = "Novell Snap-in Deep Freeze Control"
	case 7725:
		description = "Nitrogen Service"
	case 7726:
		description = "FreezeX Console Service"
	case 7727:
		description = "Trident Systems Data"
	case 7728:
		description = "Open-Source Virtual Reality"
	case 7734:
		description = "Smith Protocol over IP"
	case 7738:
		description = "HP Enterprise Discovery Agent"
	case 7741:
		description = "ScriptView Network"
	case 7742:
		description = "Mugginsoft Script Server Service"
	case 7743:
		description = "Sakura Script Transfer Protocol"
	case 7744:
		description = "RAQMON PDU"
	case 7747:
		description = "Put/Run/Get Protocol"
	case 7775:
		description = "A File System using TLS over a wide area network"
	case 7777:
		description = "cbt"
	case 7778:
		description = "Interwise"
	case 7779:
		description = "VSTAT"
	case 7781:
		description = "accu-lmgr"
	case 7784:
		description = "Seamless Bidirectional Forwarding Detection (S-BFD)"
	case 7786:
		description = "MINIVEND"
	case 7787:
		description = "Popup Reminders Receive"
	case 7789:
		description = "Office Tools Pro Receive"
	case 7794:
		description = "Q3ADE Cluster Service"
	case 7797:
		description = "Propel Connector port"
	case 7798:
		description = "Propel Encoder port"
	case 7799:
		description = "Alternate BSDP Service"
	case 7800:
		description = "Apple Software Restore"
	case 7801:
		description = "Secure Server Protocol - client"
	case 7802:
		description = "Virtualized Network Services Tunnel Protocol"
	case 7810:
		description = "Riverbed WAN Optimization Protocol"
	case 7845:
		description = "APC 7845"
	case 7846:
		description = "APC 7846"
	case 7847:
		description = "A product key authentication protocol made by CSO"
	case 7869:
		description = "MobileAnalyzer& MobileMonitor"
	case 7870:
		description = "Riverbed Steelhead Mobile Service"
	case 7871:
		description = "Mobile Device Management"
	case 7872:
		description = "TLS-based Mobile IPv6 Security"
	case 7878:
		description = "Opswise Message Service"
	case 7880:
		description = "Pearson"
	case 7887:
		description = "Universal Broker"
	case 7900:
		description = "Multicast Event"
	case 7901:
		description = "TNOS Service Protocol"
	case 7902:
		description = "TNOS shell Protocol"
	case 7903:
		description = "TNOS Secure DiaguardProtocol"
	case 7913:
		description = "QuickObjects secure port"
	case 7932:
		description = "Tier 2 Data Resource Manager"
	case 7933:
		description = "Tier 2 Business Rules Manager"
	case 7962:
		description = "Encrypted"
	case 7967:
		description = "Supercell"
	case 7979:
		description = "Micromuse-ncps"
	case 7980:
		description = "Quest Vista"
	case 7981:
		description = "Spotlight on SQL Server Desktop Collect"
	case 7982:
		description = "Spotlight on SQL Server Desktop Agent"
	case 7997:
		description = "PUSH Notification Service"
	case 7998:
		description = "USI Content Push Service"
	case 7999:
		description = "iRDMI2"
	case 8000:
		description = "iRDMI"
	case 8001:
		description = "VCOM Tunnel"
	case 8002:
		description = "Teradata ORDBMS"
	case 8003:
		description = "Mulberry Connect Reporting Service"
	case 8004:
		description = "Opensource Evolv Enterprise Platform P2P"
	case 8005:
		description = "MXI Generation II for z/OS"
	case 8006:
		description = "World Programming analytics"
	case 8007:
		description = "I/O oriented cluster computing software"
	case 8008:
		description = "HTTP Alternate"
	case 8009:
		description = "NVMe over Fabrics Discovery Service"
	case 8015:
		description = "Configuration Cloud Service"
	case 8016:
		description = "Beckhoff Automation Device Specification"
	case 8017:
		description = "Cisco Cloudsec Dataplane Port Number"
	case 8019:
		description = "QB DB Dynamic Port"
	case 8020:
		description = "Intuit Entitlement Service and Discovery"
	case 8021:
		description = "Intuit Entitlement Client"
	case 8022:
		description = "oa-system"
	case 8023:
		description = "ARCATrust vault API"
	case 8025:
		description = "CA Audit Distribution Agent"
	case 8026:
		description = "CA Audit Distribution Server"
	case 8027:
		description = "peer tracker and data relay service"
	case 8032:
		description = "ProEd"
	case 8033:
		description = "MindPrint"
	case 8034:
		description = ".vantronix Management"
	case 8040:
		description = "Ampify Messaging Protocol"
	case 8041:
		description = "Xcorpeon ASIC Carrier Ethernet Transport"
	case 8042:
		description = "FireScope Agent"
	case 8043:
		description = "FireScope Server"
	case 8044:
		description = "FireScope Management Interface"
	case 8051:
		description = "Rocrail Client Service"
	case 8052:
		description = "Senomix Timesheets Server"
	case 8060:
		description = "Asymmetric Extended Route Optimization (AERO)"
	case 8061:
		description = "Nikatron Device Protocol"
	case 8066:
		description = "Toad BI Application Server"
	case 8067:
		description = "Infinidat async replication"
	case 8070:
		description = "Oracle Unified Communication Suite's Indexed Search Converter"
	case 8074:
		description = "Gadu-Gadu"
	case 8080:
		description = "HTTP Alternate (see port 80)"
	case 8081:
		description = "Sun Proxy Admin Service"
	case 8082:
		description = "Utilistor (Client)"
	case 8083:
		description = "Utilistor (Server)"
	case 8084:
		description = "Snarl Network Protocol over HTTP"
	case 8086:
		description = "Distributed SCADA Networking Rendezvous Port"
	case 8087:
		description = "Simplify Media SPP Protocol"
	case 8088:
		description = "Radan HTTP"
	case 8090:
		description = "Vehicle to station messaging"
	case 8091:
		description = "Jam Link Framework"
	case 8097:
		description = "SAC Port Id"
	case 8100:
		description = "Xprint Server"
	case 8101:
		description = "Logical Domains Migration"
	case 8102:
		description = "Oracle Kernel zones migration server"
	case 8111:
		description = "Skynetflow network services"
	case 8115:
		description = "MTL8000 Matrix"
	case 8116:
		description = "Check Point Clustering"
	case 8117:
		description = "Purity replication clustering and remote management"
	case 8118:
		description = "Privoxy HTTP proxy"
	case 8121:
		description = "Apollo Data Port"
	case 8122:
		description = "Apollo Admin Port"
	case 8128:
		description = "PayCash Online Protocol"
	case 8129:
		description = "PayCash Wallet-Browser"
	case 8130:
		description = "INDIGO-VRMI"
	case 8131:
		description = "INDIGO-VBCP"
	case 8132:
		description = "dbabble"
	case 8140:
		description = "The Puppet master service"
	case 8148:
		description = "i-SDD file transfer"
	case 8149:
		description = "Edge of Reality game data"
	case 8153:
		description = "QuantaStor Management Interface"
	case 8160:
		description = "Patrol"
	case 8161:
		description = "Patrol SNMP"
	case 8162:
		description = "LPAR2RRD client server communication"
	case 8181:
		description = "Intermapper network management system"
	case 8182:
		description = "VMware Fault Domain Manager"
	case 8183:
		description = "ProRemote"
	case 8184:
		description = "Remote iTach Connection"
	case 8190:
		description = "Generic control plane for RPHY"
	case 8191:
		description = "Limner Pressure"
	case 8192:
		description = "SpyTech Phone Service"
	case 8194:
		description = "Bloomberg data API"
	case 8195:
		description = "Bloomberg feed"
	case 8199:
		description = "VVR DATA"
	case 8200:
		description = "TRIVNET"
	case 8201:
		description = "TRIVNET"
	case 8202:
		description = "Audio+Ethernet Standard Open Protocol"
	case 8204:
		description = "LM Perfworks"
	case 8205:
		description = "LM Instmgr"
	case 8206:
		description = "LM Dta"
	case 8207:
		description = "LM SServer"
	case 8208:
		description = "LM Webwatcher"
	case 8211:
		description = "Aruba Networks AP management"
	case 8230:
		description = "RexecJ Server"
	case 8231:
		description = "HNCP"
	case 8232:
		description = "HNCP over DTLS"
	case 8243:
		description = "Synapse Non Blocking HTTPS"
	case 8266:
		description = "ESPeasy peer-2-peer communication"
	case 8270:
		description = "Robot Framework Remote Library Interface"
	case 8276:
		description = "Microsoft Connected Cache"
	case 8280:
		description = "Synapse Non Blocking HTTP"
	case 8282:
		description = "Libelle EnterpriseBus"
	case 8292:
		description = "Bloomberg professional"
	case 8293:
		description = "Hiperscan Identification Service"
	case 8294:
		description = "Bloomberg intelligent client"
	case 8300:
		description = "Transport Management Interface"
	case 8301:
		description = "Amberon PPC/PPS"
	case 8313:
		description = "Hub Open Network"
	case 8320:
		description = "Thin(ium) Network Protocol"
	case 8321:
		description = "Thin(ium) Network Protocol"
	case 8322:
		description = "Garmin Marine"
	case 8351:
		description = "Server Find"
	case 8376:
		description = "Cruise ENUM"
	case 8377:
		description = "Cruise SWROUTE"
	case 8378:
		description = "Cruise CONFIG"
	case 8379:
		description = "Cruise DIAGS"
	case 8380:
		description = "Cruise UPDATE"
	case 8383:
		description = "M2m Services"
	case 8384:
		description = "Marathon Transport Protocol"
	case 8400:
		description = "cvd"
	case 8401:
		description = "sabarsd"
	case 8402:
		description = "abarsd"
	case 8403:
		description = "admind"
	case 8404:
		description = "SuperVault Cloud"
	case 8405:
		description = "SuperVault Backup"
	case 8415:
		description = "Delphix Session Protocol"
	case 8416:
		description = "eSpeech Session Protocol"
	case 8417:
		description = "eSpeech RTP Protocol"
	case 8423:
		description = "Aristech text-to-speech server"
	case 8432:
		description = "PostgreSQL Backup"
	case 8433:
		description = "Non Persistent Desktop and Application Streaming"
	case 8442:
		description = "CyBro A-bus Protocol"
	case 8443:
		description = "PCsync HTTPS"
	case 8444:
		description = "PCsync HTTP"
	case 8445:
		description = "Port for copy peer sync feature"
	case 8448:
		description = "Matrix Federation Protocol"
	case 8450:
		description = "npmp"
	case 8457:
		description = "Nexenta Management GUI"
	case 8470:
		description = "Cisco Address Validation Protocol"
	case 8471:
		description = "PIM over Reliable Transport"
	case 8472:
		description = "Overlay Transport Virtualization (OTV)"
	case 8473:
		description = "Virtual Point to Point"
	case 8474:
		description = "AquaMinds NoteShare"
	case 8500:
		description = "Flight Message Transfer Protocol"
	case 8501:
		description = "CYTEL Message Transfer Management"
	case 8502:
		description = "FTN Message Transfer Protocol"
	case 8503:
		description = "MPLS LSP Self-Ping"
	case 8554:
		description = "RTSP Alternate (see port 554)"
	case 8555:
		description = "SYMAX D-FENCE"
	case 8567:
		description = "DOF Tunneling Protocol"
	case 8600:
		description = "Surveillance Data"
	case 8609:
		description = "Canon Compact Printer Protocol Discovery"
	case 8610:
		description = "Canon MFNP Service"
	case 8611:
		description = "Canon BJNP Port 1"
	case 8612:
		description = "Canon BJNP Port 2"
	case 8613:
		description = "Canon BJNP Port 3"
	case 8614:
		description = "Canon BJNP Port 4"
	case 8615:
		description = "Imink Service Control"
	case 8665:
		description = "Monetra"
	case 8666:
		description = "Monetra Administrative Access"
	case 8668:
		description = "Spartan management"
	case 8675:
		description = "Motorola Solutions Customer Programming Software for Radio Management"
	case 8686:
		description = "Sun App Server - JMX/RMI"
	case 8688:
		description = "OpenRemote Controller HTTP/REST"
	case 8699:
		description = "VNYX Primary Port"
	case 8710:
		description = "gRPC for SEMI Standards implementations"
	case 8711:
		description = "Nuance Voice Control"
	case 8732:
		description = "DASGIP Net Services"
	case 8733:
		description = "iBus"
	case 8750:
		description = "DEY Storage Key Negotiation"
	case 8763:
		description = "MC-APPSERVER"
	case 8764:
		description = "OPENQUEUE"
	case 8765:
		description = "Ultraseek HTTP"
	case 8766:
		description = "Agilent Connectivity Service"
	case 8767:
		description = "Online mobile multiplayer game"
	case 8768:
		description = "Sandpolis Server"
	case 8769:
		description = "Okta MultiPlatform Access Mgmt for Cloud Svcs"
	case 8770:
		description = "Digital Photo Access Protocol (iPhoto)"
	case 8778:
		description = "Stonebranch Universal Enterprise Controller"
	case 8786:
		description = "Message Client"
	case 8787:
		description = "Message Server"
	case 8793:
		description = "Accedian Performance Measurement"
	case 8800:
		description = "Sun Web Server Admin Service"
	case 8804:
		description = "truecm"
	case 8805:
		description = "Destination Port number for PFCP"
	case 8807:
		description = "HES-CLIP Interoperability protocol"
	case 8808:
		description = "STATSports Broadcast Service"
	case 8809:
		description = "MCPTT Off-Network Protocol (MONP)"
	case 8873:
		description = "dxspider linking protocol"
	case 8880:
		description = "CDDBP"
	case 8881:
		description = "Galaxy4D Online Game Engine"
	case 8883:
		description = "Secure MQTT"
	case 8888:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 8889:
		description = "Desktop Data 1"
	case 8890:
		description = "Desktop Data 2"
	case 8891:
		description = "Desktop Data 3: NESS application"
	case 8892:
		description = "Desktop Data 4: FARM product"
	case 8893:
		description = "Desktop Data 5: NewsEDGE/Web application"
	case 8894:
		description = "Desktop Data 6: COAL application"
	case 8899:
		description = "ospf-lite"
	case 8900:
		description = "JMB-CDS 1"
	case 8901:
		description = "JMB-CDS 2"
	case 8908:
		description = "WFA Device Provisioning Protocol"
	case 8910:
		description = "manyone-http"
	case 8911:
		description = "manyone-xml"
	case 8912:
		description = "Windows Client Backup"
	case 8913:
		description = "Dragonfly System Service"
	case 8937:
		description = "Transaction Warehouse Data Service"
	case 8953:
		description = "unbound dns nameserver control"
	case 8954:
		description = "Cumulus Admin Port"
	case 8980:
		description = "Network of Devices Provider"
	case 8981:
		description = "Network of Devices Client"
	case 8989:
		description = "Sun Web Server SSL Admin Service"
	case 8990:
		description = "webmail HTTP service"
	case 8991:
		description = "webmail HTTPS service"
	case 8997:
		description = "Oracle Messaging Server Event Notification Service"
	case 8998:
		description = "Canto RoboFlow Control"
	case 8999:
		description = "Brodos Crypto Trade Protocol"
	case 9000:
		description = "CSlistener"
	case 9001:
		description = "ETL Service Manager"
	case 9002:
		description = "DynamID authentication"
	case 9005:
		description = "Golem Inter-System RPC"
	case 9006:
		description = "De-Commissioned Port"
	case 9007:
		description = "Open Grid Services Client"
	case 9008:
		description = "Open Grid Services Server"
	case 9009:
		description = "Pichat Server"
	case 9010:
		description = "Secure Data Replicator Protocol"
	case 9011:
		description = "D-Star Routing digital voice+data for amateur radio"
	case 9020:
		description = "TAMBORA"
	case 9021:
		description = "Pangolin Identification"
	case 9022:
		description = "PrivateArk Remote Agent"
	case 9023:
		description = "Secure Web Access - 1"
	case 9024:
		description = "Secure Web Access - 2"
	case 9025:
		description = "Secure Web Access - 3"
	case 9026:
		description = "Secure Web Access - 4"
	case 9050:
		description = "Versiera Agent Listener"
	case 9051:
		description = "Fusion-io Central Manager Service"
	case 9060:
		description = "CardWeb request-response I/O exchange"
	case 9080:
		description = "Groove GLRPC"
	case 9081:
		description = "Required for Adaptive Quality of Service"
	case 9082:
		description = "LCS Application Protocol"
	case 9083:
		description = "EMC PowerPath Mgmt Service"
	case 9084:
		description = "IBM AURORA Performance Visualizer"
	case 9085:
		description = "IBM Remote System Console"
	case 9086:
		description = "Vesa Net2Display"
	case 9087:
		description = "Classic Data Server"
	case 9088:
		description = "IBM Informix SQL Interface"
	case 9089:
		description = "IBM Informix SQL Interface - Encrypted"
	case 9090:
		description = "WebSM"
	case 9091:
		description = "xmltec-xmlmail"
	case 9092:
		description = "Xml-Ipc Server Reg"
	case 9093:
		description = "Copycat database replication service"
	case 9100:
		description = "PDL Data Streaming Port"
	case 9101:
		description = "Bacula Director"
	case 9102:
		description = "Bacula File Daemon"
	case 9103:
		description = "Bacula Storage Daemon"
	case 9104:
		description = "PeerWire"
	case 9105:
		description = "Xadmin Control Service"
	case 9106:
		description = "Astergate Control Service"
	case 9107:
		description = "AstergateFax Control Service"
	case 9111:
		description = "Multiple Purpose"
	case 9119:
		description = "MXit Instant Messaging"
	case 9122:
		description = "Global Relay compliant mobile instant messaging protocol"
	case 9123:
		description = "Global Relay compliant instant messaging protocol"
	case 9131:
		description = "Dynamic Device Discovery"
	case 9160:
		description = "apani1"
	case 9161:
		description = "apani2"
	case 9162:
		description = "apani3"
	case 9163:
		description = "apani4"
	case 9164:
		description = "apani5"
	case 9191:
		description = "Sun AppSvr JPDA"
	case 9200:
		description = "WAP connectionless session service"
	case 9201:
		description = "WAP session service"
	case 9202:
		description = "WAP secure connectionless session service"
	case 9203:
		description = "WAP secure session service"
	case 9204:
		description = "WAP vCard"
	case 9205:
		description = "WAP vCal"
	case 9206:
		description = "WAP vCard Secure"
	case 9207:
		description = "WAP vCal Secure"
	case 9208:
		description = "rjcdb vCard"
	case 9209:
		description = "ALMobile System Service"
	case 9210:
		description = "OMA Mobile Location Protocol"
	case 9211:
		description = "OMA Mobile Location Protocol Secure"
	case 9212:
		description = "Server View dbms access"
	case 9213:
		description = "ServerStart RemoteControl"
	case 9214:
		description = "IPDC ESG BootstrapService"
	case 9215:
		description = "Integrated Setup and Install Service"
	case 9216:
		description = "Aionex Communication Management Engine"
	case 9217:
		description = "FSC Communication Port"
	case 9222:
		description = "QSC Team Coherence"
	case 9255:
		description = "Manager On Network"
	case 9277:
		description = "GPS Data transmitted from train to ground network"
	case 9278:
		description = "Pegasus GPS Platform"
	case 9279:
		description = "Pegaus GPS System Control Interface"
	case 9280:
		description = "Predicted GPS"
	case 9281:
		description = "SofaWare transport port 1"
	case 9282:
		description = "SofaWare transport port 2"
	case 9283:
		description = "CallWaveIAM"
	case 9284:
		description = "VERITAS Information Serve"
	case 9285:
		description = "N2H2 Filter Service Port"
	case 9286:
		description = "n2 monitoring receiver"
	case 9287:
		description = "Cumulus"
	case 9292:
		description = "ArmTech Daemon"
	case 9293:
		description = "StorView Client"
	case 9294:
		description = "ARMCenter http Service"
	case 9295:
		description = "ARMCenter https Service"
	case 9300:
		description = "Virtual Racing Service"
	case 9306:
		description = "Sphinx search server (MySQL listener)"
	case 9310:
		description = "SAP Message Server"
	case 9312:
		description = "Sphinx search server"
	case 9318:
		description = "PKIX TimeStamp over TLS"
	case 9321:
		description = "guibase"
	case 9339:
		description = "gRPC Network Mgmt/Operations Interface"
	case 9340:
		description = "gRPC Routing Information Base Interface"
	case 9343:
		description = "MpIdcMgr"
	case 9344:
		description = "Mphlpdmc"
	case 9345:
		description = "Rancher Agent"
	case 9346:
		description = "C Tech Licensing"
	case 9374:
		description = "fjdmimgr"
	case 9380:
		description = "Brivs! Open Extensible Protocol"
	case 9387:
		description = "D2D Configuration Service"
	case 9388:
		description = "D2D Data Transfer Service"
	case 9389:
		description = "Active Directory Web Services"
	case 9390:
		description = "OpenVAS Transfer Protocol"
	case 9396:
		description = "fjinvmgr"
	case 9397:
		description = "MpIdcAgt"
	case 9400:
		description = "Samsung Twain for Network Server"
	case 9401:
		description = "Samsung Twain for Network Client"
	case 9402:
		description = "Samsung PC2FAX for Network Server"
	case 9418:
		description = "git pack transfer service"
	case 9443:
		description = "WSO2 Tungsten HTTPS"
	case 9444:
		description = "WSO2 ESB Administration Console HTTPS"
	case 9445:
		description = "MindArray Systems Console Agent"
	case 9450:
		description = "Sentinel Keys Server"
	case 9500:
		description = "ismserver"
	case 9522:
		description = "SMA Speedwire"
	case 9535:
		description = "Management Suite Remote Control"
	case 9536:
		description = "Surveillance buffering function"
	case 9555:
		description = "Trispen Secure Remote Access"
	case 9559:
		description = "P4Runtime gRPC Service"
	case 9592:
		description = "LANDesk Gateway"
	case 9593:
		description = "LANDesk Management Agent (cba8)"
	case 9594:
		description = "Message System"
	case 9595:
		description = "Ping Discovery Service"
	case 9596:
		description = "Mercury Discovery"
	case 9597:
		description = "PD Administration"
	case 9598:
		description = "Very Simple Ctrl Protocol"
	case 9599:
		description = "Robix"
	case 9600:
		description = "MICROMUSE-NCPW"
	case 9612:
		description = "StreamComm User Directory"
	case 9614:
		description = "iADT Protocol over TLS"
	case 9616:
		description = "eRunbook Agent"
	case 9617:
		description = "eRunbook Server"
	case 9618:
		description = "Condor Collector Service"
	case 9628:
		description = "ODBC Pathway Service"
	case 9629:
		description = "UniPort SSO Controller"
	case 9630:
		description = "Peovica Controller"
	case 9631:
		description = "Peovica Collector"
	case 9632:
		description = "Mobile-C Communications"
	case 9640:
		description = "ProQueSys Flows Service"
	case 9666:
		description = "Zoom Control Panel Game Server Management"
	case 9667:
		description = "Cross-platform Music Multiplexing System"
	case 9668:
		description = "tec5 Spectral Device Control Protocol"
	case 9694:
		description = "T-Mobile Client Wakeup Message"
	case 9695:
		description = "Content Centric Networking"
	case 9700:
		description = "Board M.I.T. Service"
	case 9747:
		description = "L5NAS Parallel Channel"
	case 9750:
		description = "Board M.I.T. Synchronous Collaboration"
	case 9753:
		description = "rasadv"
	case 9762:
		description = "WSO2 Tungsten HTTP"
	case 9800:
		description = "WebDav Source Port"
	case 9801:
		description = "Sakura Script Transfer Protocol-2"
	case 9802:
		description = "WebDAV Source TLS/SSL"
	case 9875:
		description = "Session Announcement v1"
	case 9876:
		description = "Session Director"
	case 9877:
		description = "The X.510 wrapper protocol"
	case 9878:
		description = "The KX509 Kerberized Certificate Issuance Protocol"
	case 9888:
		description = "CYBORG Systems"
	case 9889:
		description = "Port for Cable network related data proxy or repeater"
	case 9898:
		description = "MonkeyCom"
	case 9899:
		description = "SCTP TUNNELING"
	case 9900:
		description = "IUA"
	case 9901:
		description = "enrp server channel"
	case 9902:
		description = "enrp/tls server channel"
	case 9903:
		description = "Multicast Ping Protocol"
	case 9909:
		description = "domaintime"
	case 9911:
		description = "SYPECom Transport Protocol"
	case 9925:
		description = "XYBRID Cloud"
	case 9950:
		description = "APC 9950"
	case 9951:
		description = "APC 9951"
	case 9952:
		description = "APC 9952"
	case 9954:
		description = "HaloteC Instrument Network Protocol"
	case 9955:
		description = "Contact Port for AllJoyn standard messaging"
	case 9956:
		description = "Alljoyn Name Service"
	case 9966:
		description = "OKI Data Network Setting Protocol"
	case 9978:
		description = "XYBRID RT Server"
	case 9979:
		description = "Valley Information Systems Weather station data"
	case 9981:
		description = "Event sourcing database engine with a built-in programming language"
	case 9987:
		description = "DSM/SCM Target Interface"
	case 9988:
		description = "Software Essentials Secure HTTP server"
	case 9990:
		description = "OSM Applet Server"
	case 9991:
		description = "OSM Event Server"
	case 9992:
		description = "OnLive-1"
	case 9993:
		description = "OnLive-2"
	case 9994:
		description = "OnLive-3"
	case 9995:
		description = "Palace-4"
	case 9996:
		description = "Palace-5"
	case 9997:
		description = "Palace-6"
	case 9998:
		description = "Distinct32"
	case 9999:
		description = "distinct"
	case 10000:
		description = "Network Data Management Protocol"
	case 10001:
		description = "SCP Configuration"
	case 10002:
		description = "EMC-Documentum Content Server Product"
	case 10003:
		description = "EMC-Documentum Content Server Product"
	case 10004:
		description = "dasHost.exe / EMC Replication Manager Client"
	case 10005:
		description = "EMC Replication Manager Server"
	case 10006:
		description = "Sync replication protocol among different NetApp platforms"
	case 10007:
		description = "MVS Capacity"
	case 10008:
		description = "Octopus Multiplexer"
	case 10009:
		description = "Systemwalker Desktop Patrol"
	case 10010:
		description = "ooRexx rxapi services"
	case 10020:
		description = "Hardware configuration and maintenance"
	case 10023:
		description = "Comtech EF-Data's Vipersat Management Protocol"
	case 10050:
		description = "Zabbix Agent"
	case 10051:
		description = "Zabbix Trapper"
	case 10055:
		description = "Quantapoint FLEXlm Licensing Service"
	case 10080:
		description = "Amanda"
	case 10081:
		description = "FAM Archive Server"
	case 10100:
		description = "VERITAS ITAP DDTP"
	case 10101:
		description = "eZmeeting"
	case 10102:
		description = "eZproxy"
	case 10103:
		description = "eZrelay"
	case 10104:
		description = "Systemwalker Desktop Patrol"
	case 10107:
		description = "VERITAS BCTP"
	case 10110:
		description = "NMEA-0183 Navigational Data"
	case 10111:
		description = "NMEA OneNet multicast messaging"
	case 10113:
		description = "NetIQ Endpoint"
	case 10114:
		description = "NetIQ Qcheck"
	case 10115:
		description = "NetIQ Endpoint"
	case 10116:
		description = "NetIQ VoIP Assessor"
	case 10117:
		description = "NetIQ IQCResource Managament Svc"
	case 10125:
		description = "HotLink CIMple REST API"
	case 10128:
		description = "BMC-PERFORM-SERVICE DAEMON"
	case 10129:
		description = "BMC General Manager Server"
	case 10160:
		description = "QB Database Server"
	case 10161:
		description = "SNMP-TLS"
	case 10162:
		description = "SNMP-Trap-TLS"
	case 10200:
		description = "Trigence AE Soap Service"
	case 10201:
		description = "Remote Server Management Service"
	case 10252:
		description = "Apollo Relay Port"
	case 10253:
		description = "Relay of EAPOL frames"
	case 10260:
		description = "Axis WIMP Port"
	case 10261:
		description = "Tile remote machine learning"
	case 10288:
		description = "Blocks"
	case 10439:
		description = "BalanceNG session table synchronization protocol"
	case 10443:
		description = "CirrosSP Workstation Communication"
	case 10500:
		description = "HIP NAT-Traversal"
	case 10540:
		description = "MOS Media Object Metadata Port"
	case 10541:
		description = "MOS Running Order Port"
	case 10542:
		description = "MOS Low Priority Port"
	case 10543:
		description = "MOS SOAP Default Port"
	case 10544:
		description = "MOS SOAP Optional Port"
	case 10548:
		description = "Apple Document Sharing Service"
	case 10631:
		description = "Printopia Serve"
	case 10800:
		description = "Gestor de Acaparamiento para Pocket PCs"
	case 10805:
		description = "LUCIA Pareja Data Group"
	case 10809:
		description = "Network Block Device"
	case 10810:
		description = "Nuance Mobile Care Discovery"
	case 10860:
		description = "Helix Client/Server"
	case 10880:
		description = "BVEssentials HTTP API"
	case 10933:
		description = "Listen port used by the Octopus Deploy Tentacle deployment agent"
	case 10990:
		description = "Auxiliary RMI Port"
	case 11000:
		description = "IRISA"
	case 11001:
		description = "Metasys"
	case 11095:
		description = "Nest device-to-device and device-to-service application protocol"
	case 11103:
		description = "OrigoDB Server Sync Interface"
	case 11104:
		description = "NetApp Intercluster Management"
	case 11105:
		description = "NetApp Intercluster Data"
	case 11106:
		description = "SGI LK Licensing service"
	case 11108:
		description = "Hardware Terminals Discovery and Low-Level Communication Protocol"
	case 11109:
		description = "Data migration facility Manager"
	case 11110:
		description = "Data migration facility (DMF) SOAP"
	case 11111:
		description = "Viral Computing Environment (VCE)"
	case 11112:
		description = "DICOM"
	case 11161:
		description = "sun cacao snmp access point"
	case 11162:
		description = "sun cacao JMX-remoting access point"
	case 11163:
		description = "sun cacao rmi registry access point"
	case 11164:
		description = "sun cacao command-streaming access point"
	case 11165:
		description = "sun cacao web service access point"
	case 11171:
		description = "Surgical Notes Security Service Discovery (SNSS)"
	case 11172:
		description = "OEM cacao JMX-remoting access point"
	case 11173:
		description = "Straton Runtime Programing"
	case 11174:
		description = "OEM cacao rmi registry access point"
	case 11175:
		description = "OEM cacao web service access point"
	case 11201:
		description = "smsqp"
	case 11202:
		description = "DCSL Network Backup Services"
	case 11208:
		description = "WiFree Service"
	case 11211:
		description = "Memory cache service"
	case 11235:
		description = "numerical systems messaging"
	case 11319:
		description = "IMIP"
	case 11320:
		description = "IMIP Channels Port"
	case 11321:
		description = "Arena Server Listen"
	case 11367:
		description = "ATM UHAS"
	case 11371:
		description = "OpenPGP HTTP Keyserver"
	case 11430:
		description = "Lenbrook Service Discovery Protocol"
	case 11489:
		description = "ASG Cypress Secure Only"
	case 11600:
		description = "Tempest Protocol Port"
	case 11623:
		description = "EMC XtremSW distributed config"
	case 11720:
		description = "H.323 Call Control Signalling Alternate"
	case 11723:
		description = "EMC XtremSW distributed cache"
	case 11751:
		description = "Intrepid SSL"
	case 11796:
		description = "LanSchool"
	case 11876:
		description = "X2E Xoraya Multichannel protocol"
	case 11877:
		description = "X2E service discovery protocol"
	case 11967:
		description = "SysInfo Service Protocol"
	case 11971:
		description = "TiBS Service"
	case 11997:
		description = "WorldMailExpress"
	case 11998:
		description = "WorldMailExpress"
	case 11999:
		description = "WorldMailExpress"
	case 12000:
		description = "IBM Enterprise Extender SNA XID Exchange"
	case 12001:
		description = "IBM Enterprise Extender SNA COS Network Priority"
	case 12002:
		description = "IBM Enterprise Extender SNA COS High Priority"
	case 12003:
		description = "IBM Enterprise Extender SNA COS Medium Priority"
	case 12004:
		description = "IBM Enterprise Extender SNA COS Low Priority"
	case 12005:
		description = "DBISAM Database Server - Regular"
	case 12006:
		description = "DBISAM Database Server - Admin"
	case 12007:
		description = "Accuracer Database System Server"
	case 12008:
		description = "Accuracer Database System Admin"
	case 12009:
		description = "Green Hills VPN"
	case 12010:
		description = "ElevateDB Server"
	case 12012:
		description = "Vipera Messaging Service"
	case 12013:
		description = "Vipera Messaging Service over SSL Communication"
	case 12109:
		description = "RETS over SSL"
	case 12121:
		description = "NuPaper Session Service"
	case 12168:
		description = "CA Web Access Service"
	case 12172:
		description = "HiveP"
	case 12300:
		description = "LinoGrid Engine"
	case 12302:
		description = "Remote Administration Daemon"
	case 12321:
		description = "Warehouse Monitoring Syst SSS"
	case 12322:
		description = "Warehouse Monitoring Syst"
	case 12345:
		description = "Italk Chat System"
	case 12546:
		description = "Carbonite Server Replication Control"
	case 12753:
		description = "tsaf port"
	case 12865:
		description = "control port for the netperf benchmark"
	case 13160:
		description = "I-ZIPQD"
	case 13216:
		description = "Black Crow Software application logging"
	case 13217:
		description = "R&S Proxy Installation Assistant Service"
	case 13218:
		description = "EMC Virtual CAS Service"
	case 13223:
		description = "PowWow Client"
	case 13224:
		description = "PowWow Server"
	case 13400:
		description = "DoIP Data"
	case 13720:
		description = "BPRD Protocol (VERITAS NetBackup)"
	case 13722:
		description = "BP Java MSVC Protocol"
	case 13724:
		description = "Veritas Network Utility"
	case 13782:
		description = "VERITAS NetBackup"
	case 13783:
		description = "VOPIED Protocol"
	case 13785:
		description = "NetBackup Database"
	case 13786:
		description = "Veritas-nomdb"
	case 13818:
		description = "DSMCC Config"
	case 13819:
		description = "DSMCC Session Messages"
	case 13820:
		description = "DSMCC Pass-Thru Messages"
	case 13821:
		description = "DSMCC Download Protocol"
	case 13822:
		description = "DSMCC Channel Change Protocol"
	case 13823:
		description = "Blackmagic Design Streaming Server"
	case 13832:
		description = "Certificate Management and Issuing"
	case 13894:
		description = "Ultimate Control communication protocol"
	case 13929:
		description = "D-TA SYSTEMS"
	case 13930:
		description = "MedEvolve Port Requester"
	case 14000:
		description = "SCOTTY High-Speed Filetransfer"
	case 14001:
		description = "SUA"
	case 14002:
		description = "Discovery of a SCOTTY hardware codec board"
	case 14033:
		description = "sage Best! Config Server 1"
	case 14034:
		description = "sage Best! Config Server 2"
	case 14141:
		description = "VCS Application"
	case 14142:
		description = "IceWall Cert Protocol"
	case 14143:
		description = "IceWall Cert Protocol over TLS"
	case 14145:
		description = "GCM Application"
	case 14149:
		description = "Veritas Traffic Director"
	case 14150:
		description = "Veritas Cluster Server Command Server"
	case 14154:
		description = "Veritas Application Director"
	case 14250:
		description = "Fencing Server"
	case 14414:
		description = "CA eTrust Web Update Service"
	case 14500:
		description = "xpra network protocol"
	case 14936:
		description = "hde-lcesrvr-1"
	case 14937:
		description = "hde-lcesrvr-2"
	case 15000:
		description = "Hypack Data Aquisition"
	case 15002:
		description = "Open Network Environment TLS"
	case 15118:
		description = "v2g Supply Equipment Communication Controller"
	case 15345:
		description = "XPilot Contact Port"
	case 15363:
		description = "3Link Negotiation"
	case 15555:
		description = "Cisco Stateful NAT"
	case 15660:
		description = "Backup Express Restore Server"
	case 15740:
		description = "Picture Transfer Protocol"
	case 15998:
		description = "2ping Bi-Directional Ping Service"
	case 15999:
		description = "ProGrammar Enterprise"
	case 16000:
		description = "Administration Server Access"
	case 16001:
		description = "Administration Server Connector"
	case 16002:
		description = "GoodSync Mediation Service"
	case 16003:
		description = "Automation and Control by REGULACE.ORG"
	case 16020:
		description = "Filemaker Java Web Publishing Core"
	case 16021:
		description = "Filemaker Java Web Publishing Core Binary"
	case 16161:
		description = "Solaris SEA Port"
	case 16162:
		description = "Solaris Audit - secure remote audit log"
	case 16309:
		description = "etb4j"
	case 16310:
		description = "Policy Distribute"
	case 16311:
		description = "Policy definition and update management"
	case 16360:
		description = "Network Serial Extension Ports One"
	case 16361:
		description = "Network Serial Extension Ports Two"
	case 16367:
		description = "Network Serial Extension Ports Three"
	case 16368:
		description = "Network Serial Extension Ports Four"
	case 16384:
		description = "Connected Corp"
	case 16385:
		description = "Reliable Datagram Sockets"
	case 16619:
		description = "X509 Objects Management Service"
	case 16665:
		description = "Reliable multipath data transport for high latencies"
	case 16666:
		description = "Vidder Tunnel Protocol"
	case 16900:
		description = "Newbay Mobile Client Update Service"
	case 16950:
		description = "Simple Generic Client Interface Protocol"
	case 16991:
		description = "INTEL-RCI-MP"
	case 16992:
		description = "Intel(R) AMT SOAP/HTTP"
	case 16993:
		description = "Intel(R) AMT SOAP/HTTPS"
	case 16994:
		description = "Intel(R) AMT Redirection/TCP"
	case 16995:
		description = "Intel(R) AMT Redirection/TLS"
	case 17010:
		description = "Plan 9 cpu port"
	case 17184:
		description = "Vestas Data Layer Protocol"
	case 17185:
		description = "Sounds Virtual"
	case 17219:
		description = "Chipper"
	case 17220:
		description = "IEEE 1722 Transport Protocol for Time Sensitive Applications"
	case 17221:
		description = "IEEE 1722.1 AVB Discovery"
	case 17222:
		description = "Control Plane Synchronization Protocol (SPSP)"
	case 17223:
		description = "ISA100 GCI"
	case 17224:
		description = "Train Realtime Data Protocol (TRDP) Process Data"
	case 17225:
		description = "Train Realtime Data Protocol (TRDP) Message Data"
	case 17234:
		description = "Integrius Secure Tunnel Protocol"
	case 17235:
		description = "SSH Tectia Manager"
	case 17500:
		description = "Dropbox LanSync Protocol"
	case 17555:
		description = "Ailith management of routers"
	case 17729:
		description = "Eclipse Aviation"
	case 17754:
		description = "Encap. ZigBee Packets"
	case 17755:
		description = "ZigBee IP Transport Service"
	case 17756:
		description = "ZigBee IP Transport Secure Service"
	case 17777:
		description = "SolarWinds Orion"
	case 18000:
		description = "Beckman Instruments"
	case 18104:
		description = "RAD PDF Service"
	case 18136:
		description = "z/OS Resource Access Control Facility"
	case 18181:
		description = "OPSEC CVP"
	case 18182:
		description = "OPSEC UFP"
	case 18183:
		description = "OPSEC SAM"
	case 18184:
		description = "OPSEC LEA"
	case 18185:
		description = "OPSEC OMI"
	case 18186:
		description = "Occupational Health SC"
	case 18187:
		description = "OPSEC ELA"
	case 18242:
		description = "Checkpoint router monitoring"
	case 18243:
		description = "Checkpoint router state backup"
	case 18262:
		description = "GV NetConfig Service"
	case 18463:
		description = "AC Cluster"
	case 18516:
		description = "HeyThings Device communicate service"
	case 18634:
		description = "Reliable Datagram Service"
	case 18635:
		description = "Reliable Datagram Service over IP"
	case 18668:
		description = "Manufacturing Execution Systems Mesh Communication"
	case 18769:
		description = "IQue Protocol"
	case 18881:
		description = "Infotos"
	case 18888:
		description = "APCNECMP"
	case 19000:
		description = "iGrid Server"
	case 19007:
		description = "Scintilla protocol for device services"
	case 19020:
		description = "J-Link TCP/IP Protocol"
	case 19191:
		description = "OPSEC UAA"
	case 19194:
		description = "UserAuthority SecureAgent"
	case 19220:
		description = "Client Connection Management and Data Exchange Service"
	case 19283:
		description = "Key Server for SASSAFRAS"
	case 19315:
		description = "Key Shadow for SASSAFRAS"
	case 19398:
		description = "mtrgtrans"
	case 19410:
		description = "hp-sco"
	case 19411:
		description = "hp-sca"
	case 19412:
		description = "HP-SESSMON"
	case 19539:
		description = "FXUPTP"
	case 19540:
		description = "SXUPTP"
	case 19541:
		description = "JCP Client"
	case 19788:
		description = "Mesh Link Establishment"
	case 19790:
		description = "FairCom Database"
	case 19998:
		description = "IEC 60870-5-104 process control - secure"
	case 19999:
		description = "Distributed Network Protocol - Secure"
	case 20000:
		description = "Distributed Network Protocol"
	case 20001:
		description = "MicroSAN"
	case 20002:
		description = "Commtact HTTP"
	case 20003:
		description = "Commtact HTTPS"
	case 20005:
		description = "OpenWebNet protocol for electric network"
	case 20012:
		description = "Samsung Interdevice Interaction discovery"
	case 20013:
		description = "Samsung Interdevice Interaction"
	case 20014:
		description = "OpenDeploy Listener"
	case 20034:
		description = "NetBurner ID Port"
	case 20046:
		description = "TMOP HL7 Message Transfer Service"
	case 20048:
		description = "NFS mount protocol"
	case 20049:
		description = "Network File System (NFS) over RDMA"
	case 20057:
		description = "AvesTerra Hypergraph Transfer Protocol (HGTP)"
	case 20167:
		description = "TOLfab Data Change"
	case 20202:
		description = "IPD Tunneling Port"
	case 20222:
		description = "iPulse-ICS"
	case 20480:
		description = "emWave Message Service"
	case 20670:
		description = "Track"
	case 20810:
		description = "CRTech NLM"
	case 20999:
		description = "AT Hand MMP"
	case 21000:
		description = "IRTrans Control"
	case 21010:
		description = "Notezilla.Lan Server"
	case 21212:
		description = "Distributed artificial intelligence"
	case 21213:
		description = "Cohesity backup agents"
	case 21221:
		description = "Services for Air Server"
	case 21553:
		description = "Raima RDM TFS"
	case 21554:
		description = "MineScape Design File Server"
	case 21590:
		description = "VoFR Gateway"
	case 21800:
		description = "TVNC Pro Multiplexing"
	case 21801:
		description = "Safe AutoLogon"
	case 21845:
		description = "webphone"
	case 21846:
		description = "NetSpeak Corp. Directory Services"
	case 21847:
		description = "NetSpeak Corp. Connection Services"
	case 21848:
		description = "NetSpeak Corp. Automatic Call Distribution"
	case 21849:
		description = "NetSpeak Corp. Credit Processing System"
	case 22000:
		description = "SNAPenetIO"
	case 22001:
		description = "OptoControl"
	case 22002:
		description = "Opto Host Port 2"
	case 22003:
		description = "Opto Host Port 3"
	case 22004:
		description = "Opto Host Port 4"
	case 22005:
		description = "Opto Host Port 5"
	case 22125:
		description = "dCache Access Protocol"
	case 22128:
		description = "GSI dCache Access Protocol"
	case 22222:
		description = "dasHost.exe / EasyEngine"
	case 22273:
		description = "wnn6"
	case 22305:
		description = "CompactIS Tunnel"
	case 22333:
		description = "ShowCockpit Networking"
	case 22335:
		description = "Initium Labs Security and Automation Control"
	case 22343:
		description = "CompactIS Secure Tunnel"
	case 22347:
		description = "WibuKey Standard WkLan"
	case 22350:
		description = "CodeMeter Standard"
	case 22351:
		description = "TPC/IP requests of copy protection software to a server"
	case 22537:
		description = "CaldSoft Backup server file transfer"
	case 22555:
		description = "Vocaltec Internet Phone"
	case 22763:
		description = "Talika Main Server"
	case 22800:
		description = "Telerate Information Platform LAN"
	case 22951:
		description = "Telerate Information Platform WAN"
	case 23000:
		description = "Inova LightLink Server Type 1"
	case 23001:
		description = "Inova LightLink Server Type 2"
	case 23002:
		description = "Inova LightLink Server Type 3"
	case 23003:
		description = "Inova LightLink Server Type 4"
	case 23004:
		description = "Inova LightLink Server Type 5"
	case 23005:
		description = "Inova LightLink Server Type 6"
	case 23053:
		description = "Generic Notification Transport Protocol"
	case 23272:
		description = "S102 application"
	case 23294:
		description = "5AFE SDN Directory"
	case 23333:
		description = "Emulex HBAnyware Remote Management"
	case 23400:
		description = "Novar Data"
	case 23401:
		description = "Novar Alarm"
	case 23402:
		description = "Novar Global"
	case 23456:
		description = "Aequus Service"
	case 23457:
		description = "Aequus Service Mgmt"
	case 23546:
		description = "AreaGuard Neo - WebServer"
	case 24000:
		description = "med-ltp"
	case 24001:
		description = "med-fsp-rx"
	case 24002:
		description = "med-fsp-tx"
	case 24003:
		description = "med-supp"
	case 24004:
		description = "med-ovw"
	case 24005:
		description = "med-ci"
	case 24006:
		description = "med-net-svc"
	case 24242:
		description = "fileSphere"
	case 24249:
		description = "Vista 4GL"
	case 24321:
		description = "Isolv Local Directory"
	case 24322:
		description = "Transport of Human Interface Device data streams"
	case 24323:
		description = "Verimag mobile class protocol over TCP"
	case 24386:
		description = "Intel RCI"
	case 24465:
		description = "Tonido Domain Server"
	case 24554:
		description = "BINKP"
	case 24577:
		description = "bilobit Service"
	case 24666:
		description = "SmarDTV"
	case 24676:
		description = "Canditv Message Service"
	case 24677:
		description = "FlashFiler"
	case 24678:
		description = "Turbopower Proactivate"
	case 24680:
		description = "TCC User HTTP Service"
	case 24754:
		description = "Citrix StorageLink Gateway"
	case 24850:
		description = "Device Association Discovery"
	case 24922:
		description = "Find Identification of Network Devices"
	case 25000:
		description = "icl-twobase1"
	case 25001:
		description = "icl-twobase2"
	case 25002:
		description = "icl-twobase3"
	case 25003:
		description = "icl-twobase4"
	case 25004:
		description = "icl-twobase5"
	case 25005:
		description = "icl-twobase6"
	case 25006:
		description = "icl-twobase7"
	case 25007:
		description = "icl-twobase8"
	case 25008:
		description = "icl-twobase9"
	case 25009:
		description = "icl-twobase10"
	case 25100:
		description = "IBM Db2 Client Interface - Encrypted"
	case 25471:
		description = "RNSAP User Adaptation for Iurh"
	case 25576:
		description = "Sauter Dongle"
	case 25604:
		description = "Identifier Tracing Protocol"
	case 25793:
		description = "Vocaltec Address Server"
	case 25900:
		description = "TASP Network Comm"
	case 25901:
		description = "NIObserver"
	case 25902:
		description = "NILinkAnalyst"
	case 25903:
		description = "NIProbe"
	case 25954:
		description = "Bitfighter game server"
	case 25955:
		description = "Bitfighter master server"
	case 26000:
		description = "quake"
	case 26133:
		description = "Symbolic Computation Software Composability Protocol"
	case 26208:
		description = "wnn6-ds"
	case 26257:
		description = "CockroachDB"
	case 26260:
		description = "eZproxy"
	case 26261:
		description = "eZmeeting"
	case 26262:
		description = "K3 Software-Server"
	case 26263:
		description = "K3 Software-Client"
	case 26264:
		description = "De-registered"
	case 26486:
		description = "EXOline"
	case 26487:
		description = "EXOconfig"
	case 26489:
		description = "EXOnet"
	case 27010:
		description = "A protocol for managing license services"
	case 27016:
		description = "Cloud hosting environment network"
	case 27017:
		description = "Mongo database system"
	case 27345:
		description = "ImagePump"
	case 27442:
		description = "Job controller service"
	case 27504:
		description = "Kopek HTTP Head Port"
	case 27782:
		description = "ARS VISTA Application"
	case 27876:
		description = "Astrolink Protocol"
	case 27999:
		description = "Attribute Certificate Services"
	case 28000:
		description = "NX License Manager"
	case 28001:
		description = "PQ Service"
	case 28010:
		description = "Gruber cash registry protocol"
	case 28080:
		description = "thor/server - ML engine"
	case 28119:
		description = "A27 cdma2000 RAN Management"
	case 28200:
		description = "VoxelStorm game server"
	case 28240:
		description = "Siemens GSM"
	case 29000:
		description = "Siemens Licensing Server"
	case 29118:
		description = "SGsAP in 3GPP"
	case 29167:
		description = "ObTools Message Protocol"
	case 29168:
		description = "SBcAP in 3GPP"
	case 29169:
		description = "HNBAP and RUA Common Association"
	case 29999:
		description = "IEC61850 - wind power plants"
	case 30000:
		description = "Secure Network Data Management Protocol"
	case 30001:
		description = "Pago Services 1"
	case 30002:
		description = "Pago Services 2"
	case 30003:
		description = "Amicon FPSU-IP Remote Administration"
	case 30004:
		description = "Amicon FPSU-IP VPN"
	case 30100:
		description = "Remote Window Protocol"
	case 30260:
		description = "Kingdoms Online (CraigAvenue)"
	case 30400:
		description = "GroundStar RealTime System"
	case 30832:
		description = "Samsung Convergence Discovery Protocol"
	case 30999:
		description = "OpenView Service Desk Client"
	case 31016:
		description = "Kollective Agent Kollective Delivery Protocol"
	case 31020:
		description = "Autotrac ACP 245"
	case 31029:
		description = "YaWN - Yet Another Windows Notifier"
	case 31337:
		description = "eldim is a secure file upload proxy"
	case 31400:
		description = "PACE license server"
	case 31416:
		description = "XQoS network monitor"
	case 31457:
		description = "TetriNET Protocol"
	case 31620:
		description = "lm mon"
	case 31685:
		description = "DS Expert Monitor"
	case 31765:
		description = "GameSmith Port"
	case 31948:
		description = "Embedded Device Configuration Protocol TX"
	case 31949:
		description = "Embedded Device Configuration Protocol RX"
	case 32034:
		description = "iRacing helper service"
	case 32249:
		description = "T1 Distributed Processor"
	case 32400:
		description = "Plex multimedia"
	case 32483:
		description = "Access Point Manager Link"
	case 32635:
		description = "SecureNotebook-CLNT"
	case 32636:
		description = "DMExpress"
	case 32767:
		description = "FileNet BPM WS-ReliableMessaging Client"
	case 32768:
		description = "Filenet TMS"
	case 32769:
		description = "Filenet RPC"
	case 32770:
		description = "Filenet NCH"
	case 32771:
		description = "FileNet RMI"
	case 32772:
		description = "FileNET Process Analyzer"
	case 32773:
		description = "FileNET Component Manager"
	case 32774:
		description = "FileNET Rules Engine"
	case 32775:
		description = "Performance Clearinghouse"
	case 32776:
		description = "FileNET BPM IOR"
	case 32777:
		description = "FileNet BPM CORBA"
	case 32801:
		description = "Multiple Listing Service Network"
	case 32811:
		description = "Real Estate Transport Protocol"
	case 32896:
		description = "Attachmate ID Manager"
	case 33000:
		description = "WatchGuard Endpoint Communications"
	case 33060:
		description = "MySQL Database Extended Interface"
	case 33123:
		description = "Aurora (Balaena Ltd)"
	case 33331:
		description = "DiamondCentral Interface"
	case 33333:
		description = "Digital Gaslight Service"
	case 33334:
		description = "SpeedTrace TraceAgent"
	case 33434:
		description = "traceroute use"
	case 33435:
		description = "IP Multicast Traceroute"
	case 33656:
		description = "SNIP Slave"
	case 33890:
		description = "Adept IP protocol"
	case 34249:
		description = "TurboNote Relay Server Default Port"
	case 34378:
		description = "P-Net on IP local"
	case 34379:
		description = "P-Net on IP remote"
	case 34567:
		description = "dhanalakshmi.org EDI Service"
	case 34962:
		description = "PROFInet RT Unicast"
	case 34963:
		description = "PROFInet RT Multicast"
	case 34964:
		description = "PROFInet Context Manager"
	case 34980:
		description = "EtherCAT Port"
	case 35000:
		description = "HeathView"
	case 35001:
		description = "ReadyTech Viewer"
	case 35002:
		description = "ReadyTech Sound Server"
	case 35003:
		description = "ReadyTech DeviceMapper Server"
	case 35004:
		description = "ReadyTech ClassManager"
	case 35005:
		description = "ReadyTech LabTracker"
	case 35006:
		description = "ReadyTech Helper Service"
	case 35100:
		description = "Axiomatic discovery protocol"
	case 35354:
		description = "KIT Messenger"
	case 35355:
		description = "Altova License Management"
	case 35356:
		description = "Gutters Note Exchange"
	case 35357:
		description = "OpenStack ID Service"
	case 36001:
		description = "AllPeers Network"
	case 36411:
		description = "Wireless LAN Control plane Protocol (WLCP)"
	case 36412:
		description = "S1-Control Plane (3GPP)"
	case 36422:
		description = "X2-Control Plane (3GPP)"
	case 36423:
		description = "SLm Interface Application Protocol"
	case 36424:
		description = "Nq and Nq' Application Protocol"
	case 36443:
		description = "M2 Application Part"
	case 36444:
		description = "M3 Application Part"
	case 36462:
		description = "Xw-Control Plane (3GPP)"
	case 36524:
		description = "Febooti Automation Workshop"
	case 36602:
		description = "Observium statistics collection agent"
	case 36700:
		description = "MapX communication"
	case 36865:
		description = "KastenX Pipe"
	case 37472:
		description = "W1 signalling transport"
	case 37475:
		description = "science + computing's Venus Administration Port"
	case 37483:
		description = "Google Drive Sync"
	case 37601:
		description = "Epipole File Transfer Protocol"
	case 37654:
		description = "Unisys ClearPath ePortal"
	case 38000:
		description = "InfoVista Server Database"
	case 38001:
		description = "InfoVista Server Insertion"
	case 38002:
		description = "Cresco Controller"
	case 38201:
		description = "Galaxy7 Data Tunnel"
	case 38202:
		description = "Fairview Message Service"
	case 38203:
		description = "AppGate Policy Server"
	case 38412:
		description = "NG Control Plane (3GPP)"
	case 38422:
		description = "Xn Control Plane (3GPP)"
	case 38462:
		description = "E1 signalling transport (3GPP)"
	case 38472:
		description = "F1 Control Plane (3GPP)"
	case 38638:
		description = "Premier SQL Middleware Server"
	case 38865:
		description = "secRMM SafeCopy"
	case 39063:
		description = "Children's hearing test/Telemedicine"
	case 39681:
		description = "TurboNote Default Port"
	case 40000:
		description = "SafetyNET p"
	case 40023:
		description = "K-PatentsSensorInformation"
	case 40404:
		description = "Simplify Printing TX"
	case 40841:
		description = "CSCP"
	case 40842:
		description = "CSCCREDIR"
	case 40843:
		description = "CSCCFIREWALL"
	case 40853:
		description = "ORTEC Service Discovery"
	case 41111:
		description = "Foursticks QoS Protocol"
	case 41121:
		description = "Tentacle Server"
	case 41230:
		description = "Z-Wave Protocol over SSL/TLS"
	case 41794:
		description = "Crestron Control Port"
	case 41795:
		description = "Crestron Terminal Port"
	case 41796:
		description = "Crestron Secure Control Port"
	case 41797:
		description = "Crestron Secure Terminal Port"
	case 42508:
		description = "Computer Associates network discovery protocol"
	case 42509:
		description = "CA discovery response"
	case 42510:
		description = "CA eTrust RPC"
	case 42999:
		description = "API endpoint for search application"
	case 43000:
		description = "Receiver Remote Control"
	case 43188:
		description = "REACHOUT"
	case 43189:
		description = "NDM-AGENT-PORT"
	case 43190:
		description = "IP-PROVISION"
	case 43191:
		description = "Reconnoiter Agent Data Transport"
	case 43210:
		description = "Shaper Automation Server Management"
	case 43438:
		description = "HmIP LAN Routing"
	case 43439:
		description = "EQ3 firmware update"
	case 43440:
		description = "Cisco EnergyWise Management"
	case 43441:
		description = "Cisco NetMgmt DB Ports"
	case 44123:
		description = "Z-Wave Secure Tunnel"
	case 44321:
		description = "PCP server (pmcd)"
	case 44322:
		description = "PCP server (pmcd) proxy"
	case 44323:
		description = "HTTP binding for Performance Co-Pilot client API"
	case 44444:
		description = "Cognex DataMan Management Protocol"
	case 44445:
		description = "Acronis Backup Gateway service port"
	case 44544:
		description = "DOMIQ Building Automation"
	case 44553:
		description = "REALbasic Remote Debug"
	case 44600:
		description = "AudioScience HPI"
	case 44818:
		description = "EtherNet/IP messaging"
	case 44900:
		description = "M3DA is used for efficient machine-to-machine communications"
	case 45000:
		description = "Nuance AutoStore Status Monitoring Protocol (data transfer)"
	case 45001:
		description = "Nuance AutoStore Status Monitoring Protocol (secure data transfer)"
	case 45002:
		description = "Redspeed Status Monitor"
	case 45045:
		description = "Remote application control protocol"
	case 45054:
		description = "InVision AG"
	case 45185:
		description = "Wire and Wireless transfer on synchroniz"
	case 45514:
		description = "ASSIA CloudCheck WiFi Management System"
	case 45678:
		description = "EBA PRISE"
	case 45824:
		description = "Server for the DAI family of client-server products"
	case 45825:
		description = "Qpuncture Data Access Service"
	case 45966:
		description = "SSRServerMgr"
	case 46336:
		description = "Listen port used for Inedo agent communication"
	case 46999:
		description = "MediaBox Server"
	case 47000:
		description = "Message Bus"
	case 47001:
		description = "Windows Remote Management Service"
	case 47100:
		description = "Configuration of motors connected to Industrial Ethernet"
	case 47557:
		description = "Databeam Corporation"
	case 47624:
		description = "Direct Play Server"
	case 47806:
		description = "ALC Protocol"
	case 47808:
		description = "Building Automation and Control Networks"
	case 47809:
		description = "PreSonus Universal Control Network Protocol"
	case 48000:
		description = "Nimbus Controller"
	case 48001:
		description = "Nimbus Spooler"
	case 48002:
		description = "Nimbus Hub"
	case 48003:
		description = "Nimbus Gateway"
	case 48004:
		description = "NimbusDB Connector"
	case 48005:
		description = "NimbusDB Control"
	case 48048:
		description = "Juliar Programming Language Protocol"
	case 48049:
		description = "3GPP Cell Broadcast Service Protocol"
	case 48050:
		description = "WeFi Access Network Discovery and Selection Function"
	case 48128:
		description = "Image Systems Network Services"
	case 48129:
		description = "Bloomberg locator"
	case 48556:
		description = "com-bardac-dw"
	case 48619:
		description = "iqobject"
	case 48653:
		description = "Robot Raconteur transport"
	case 49000:
		description = "Matahari Broker"
	case 49001:
		description = "Nuance Unity Service Request Protocol"
	case 49150:
		description = "InSpider System"
	default:
		description = ""
	}

	// Return the description in parentheses unless it's empty
	if description == "" {
		return description
	}

	return "(" + description + ")"

}
