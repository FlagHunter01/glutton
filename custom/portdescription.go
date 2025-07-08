// [CUSTOM]
package custom

// Returns a string describing the port (empty if the port is unknown)
func PortDescription(port uint16) string {

	// Describes the port
	var description string

	// Assigns a value to the description for some intresting ports
	switch port {
	case 20:
		description = "FTP"
	case 21:
		description = "FTP (command)"
	case 22:
		description = "SSH"
	case 23:
		description = "Telnet"
	case 25:
		description = "SMTP"
	case 107:
		description = "RTelnet"
	case 115:
		description = "SFTP"
	case 135:
		description = "Microsoft EPMAP"
	case 137:
		description = "SAMBA"
	case 143:
		description = "IMAP"
	case 156:
		description = "SQL"
	case 209:
		description = "Quick Mail Transfer Protocol"
	case 220:
		description = "IMAP"
	case 443:
		description = "HTTPS"
	case 445:
		description = "Microsoft-DS"
	case 666:
		description = "DOOM"
	case 1433:
		description = "MSSQL"
	case 1755:
		description = "Microsoft Media Services"
	case 3306:
		description = "MySQL"
	case 3389:
		description = "RDP"
	case 5000:
		description = "This honeypot's TCP redirect port"
	case 5001:
		description = "This honeypot's UDP redirect port"
	case 5357:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 5358:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 5722:
		description = "Microsoft RPC, DFSR (SYSVOL) Replication Service"
	case 5938:
		description = "TeamViewer"
	case 5985:
		description = "WinRM-HTTP"
	case 5986:
		description = "WinRM-HTTPS"
	case 6571:
		description = "Windows Live FolderShare client "
	case 8530:
		description = "Windows Server Update Services over HTTP"
	case 8531:
		description = "Windows Server Update Services over HTTPS"
	case 8888:
		description = "⚠️ WARNING: This port is used for UPnP. It should NOT be intercepted by Glutton"
	case 9389:
		description = "adws, Microsoft AD DS Web Services, Powershell uses this port"
	default:
		description = ""
	}

	// Return the description in parentheses unless it's empty
	if description == "" {
		return description
	}
	
	return "(" + description + ")"

}
