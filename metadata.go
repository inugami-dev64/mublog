package mublog

var Metadata map[string]string

func FindMetadataDecl(loc int, md []byte) int {
	sepCount := 0
	for i := loc; i < len(md); i++ {
		if md[i] == '-' {
			sepCount++
		} else {
			sepCount = 0
		}

		if sepCount == 3 {
			return i + 1
		}
	}

	return -1
}

func ParseMetadata(md []byte) []byte {
	Metadata = make(map[string]string)
	var beg int = FindMetadataDecl(0, md)
	if beg == -1 {
		return md
	}

	var end int = FindMetadataDecl(beg, md)
	if end == -1 {
		return md
	}

	var otherMd []byte = md[end:]

	readKeyword := true
	var keyword []byte
	var value []byte
	for i := beg + 1; i < end-3; i++ {
		if md[i] == '\n' {
			Metadata[string(keyword)] = string(value)
			keyword = nil
			value = nil
			readKeyword = true
		} else if readKeyword == true && md[i] != ' ' {
			// check if trailing ':' has been reached
			if md[i] == ':' {
				readKeyword = false
				i++
				continue
			}

			// read the keyword
			keyword = append(keyword, md[i])
		} else if readKeyword == false && md[i] != '\n' {
			// read the value
			value = append(value, md[i])
		}
	}

	return otherMd
}
