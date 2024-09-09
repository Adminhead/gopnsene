package pkg

import "fmt"

type CoreBackupController struct {
	APIClient
}

type CoreBackupBackupsResponse struct {
	Items []struct {
		Time        string `json:"time"`
		TimeISO     string `json:"time_iso"`
		Description string `json:"description"`
		Username    string `json:"username"`
		Filesize    int    `json:"filesize"`
		ID          string `json:"id"`
	} `json:"items"`
}

type CoreBackupDiffResponse struct {
	Items []string `json:"items"`
}

type CoreBackupProviderResponse struct {
	Items struct {
		This struct {
			Description string `json:"description"`
			DirName     string `json:"dirname"`
		} `json:"this"`
	} `json:"items"`
}

// Backup returns a new backup controller
func (c *CoreModule) Backup() *CoreBackupController {
	nc := c.APIClient
	nc.Controller = "backup"
	nc.Host = "this"
	return &CoreBackupController{
		nc,
	}

}

// SetHost sets the host for the backup controller (defaults to "this")
func (c *CoreBackupController) SetHost(host string) *CoreBackupController {
	c.Host = host
	return c
}

// Backups returns a list of backups
func (c *CoreBackupController) Backups() (*CoreBackupBackupsResponse, error) {
	c.Command = "backups"
	result := &CoreBackupBackupsResponse{}
	_, err := c.client.R().SetResult(result).Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host))
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}

// DeleteBackup deletes the backup with the given ID
func (c *CoreBackupController) DeleteBackup(backup string) error {
	c.Command = "deleteBackup"
	_, err := c.client.R().Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, backup))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Diff returns the difference between two backups
func (c *CoreBackupController) Diff(backupOne, backupTwo string) (CoreBackupDiffResponse, error) {
	c.Command = "diff"
	var results CoreBackupDiffResponse
	_, err := c.client.R().SetResult(&results).Get(fmt.Sprintf("/%s/%s/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host, backupOne, backupTwo))
	if err != nil {
		fmt.Println(err)
		return results, err
	}
	return results, nil
}

// Download returns the text value of a backup with the given ID
func (c *CoreBackupController) Download(backup string) (string, error) {
	c.Command = "download"
	resp, err := c.client.R().Get(fmt.Sprintf("/%s/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host, backup))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(resp.Body()), nil
}

// Providers returns the available backup providers
func (c *CoreBackupController) Providers() (CoreBackupProviderResponse, error) {
	c.Command = "providers"
	var results CoreBackupProviderResponse
	resp, err := c.client.R().SetResult(&results).Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host))
	if err != nil {
		fmt.Println(err)
		return results, err
	}
	fmt.Println(resp)
	return results, nil
}

// RevertBackup reverts the system to the given backup ID
func (c *CoreBackupController) RevertBackup(backup string) error {
	c.Command = "restore"
	_, err := c.client.R().Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, backup))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
