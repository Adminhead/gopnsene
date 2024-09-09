package pkg

import "fmt"

type CoreBackupController struct {
	APIClient
}

type CoreBackupBackupsItem struct {
	Time        string `json:"time"`
	TimeISO     string `json:"time_iso"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Filesize    int    `json:"filesize"`
	ID          string `json:"id"`
}

type CoreBackupBackupsResponse struct {
	Items []CoreBackupBackupsItem `json:"items"`
}

type CoreBackupDiffResponse struct {
	Items []string `json:"items"`
}

func (c *CoreModule) Backup() *CoreBackupController {
	nc := c.APIClient
	nc.Controller = "backup"
	nc.Host = "this"
	return &CoreBackupController{
		nc,
	}

}

func (c *CoreBackupController) SetHost(host string) *CoreBackupController {
	c.Host = host
	return c
}

func (c *CoreBackupController) Backups() (*CoreBackupBackupsResponse, error) {
	c.Command = "backups"
	result := &CoreBackupBackupsResponse{}
	_, err := c.client.R().SetResult(result).Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host))
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	//fmt.Println(resp.Result())
	return result, nil
}

func (c *CoreBackupController) DeleteBackup(backup string) error {
	c.Command = "deleteBackup"
	_, err := c.client.R().Get(fmt.Sprintf("/%s/%s/%s/%s", c.Module, c.Controller, c.Command, backup))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

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

func (c *CoreBackupController) Download(backup string) error {
	c.Command = "download"
	resp, err := c.client.R().Get(fmt.Sprintf("/%s/%s/%s/%s/%s", c.Module, c.Controller, c.Command, c.Host, backup))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	return nil
}
