package sysuserrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"database/sql"
	"errors"
	"fmt"
)

func newSysUserMysqlImpl(apexConn *sql.DB) SysUserRepo {
	return &SysUserMysqlImpl{
		apexDb: apexConn,
	}
}

type SysUserMysqlImpl struct {
	apexDb *sql.DB
}

func (s *SysUserMysqlImpl) CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	stmt, er := s.apexDb.Prepare(`INSERT INTO sys_daftar_user(
		user_name,
		user_password,
		nama_lengkap,
		unit_kerja,
		jabatan,
		user_code,
		tgl_expired,
		user_web_password,
		flag,
		status_aktif,
		penerimaan,
		pengeluaran
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare add sys user : ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Exec..
	if _, er := stmt.Exec(
		newSysUser.User_Name,
		newSysUser.User_Password,
		newSysUser.Nama_Lengkap,
		newSysUser.Unit_Kerja,
		newSysUser.Jabatan,
		newSysUser.User_Code,
		newSysUser.Tgl_Expired,
		newSysUser.User_Web_Password,
		newSysUser.Flag,
		newSysUser.Status_Aktif,
		newSysUser.Penerimaan,
		newSysUser.Pengeluaran); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while add tabung : ", er.Error()))
	} else {
		return newSysUser, nil
	}

}

func (s *SysUserMysqlImpl) UpdateSysDaftarUser(updSysuser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET nama_lengkap = ? WHERE user_name = ?")
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare update user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		updSysuser.Nama_Lengkap,
		updSysuser.User_Name); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while update user : ", er.Error()))
	}

	return updSysuser, nil

}

func (s *SysUserMysqlImpl) HardDeleteSysDaftarUser(kodeLkm string) (er error) {

	stmt, er := s.apexDb.Prepare("DELETE FROM sys_daftar_user WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
	}

	return nil
}

func (s *SysUserMysqlImpl) DeleteSysDaftarUser(kodeLkm string) (er error) {

	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET status_aktif = 0 WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
	}

	return nil
}

func (s *SysUserMysqlImpl) ResetUserPassword(user entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {
	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET user_web_password = ? WHERE user_name = ?")
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare update apex password: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(user.User_Web_Password, user.User_Name); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while update apex password: ", er.Error()))
	}

	return user, nil
}

func (s *SysUserMysqlImpl) FindByUserName(userName string) (user entities.SysDaftarUser, er error) {
	row := s.apexDb.QueryRow(`SELECT
		user_id, 
		user_name,
		nama_lengkap,
		tgl_expired,
		user_web_password
	FROM sys_daftar_user WHERE user_name = ? LIMIT 1`, userName)
	er = row.Scan(
		&user.User_Id,
		&user.User_Name,
		&user.Nama_Lengkap,
		&user.Tgl_Expired,
		&user.User_Web_Password_Hash,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return user, err.NoRecord
		} else {
			return user, errors.New(fmt.Sprint("error while get user name: ", er.Error()))
		}
	}
	return
}

func (s *SysUserMysqlImpl) GetListSysApexRoutingRekInduk() (list []web.RoutingRekIndukData, er error) {
	rows, er := s.apexDb.Query("SELECT norek, norek_induk FROM routing_rek_induk")
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var scGroup web.RoutingRekIndukData
		if er = rows.Scan(&scGroup.KodeLkm, &scGroup.NorekInduk); er != nil {
			return list, er
		}

		list = append(list, scGroup)
	}

	if len(list) == 0 {
		return list, err.NoRecord
	} else {
		return
	}
}

func (s *SysUserMysqlImpl) CreateSysApexRoutingRekInduk(bankCode, norekInduk string) (routing web.RoutingRekIndukData, er error) {

	stmt, er := s.apexDb.Prepare(`INSERT INTO routing_rek_induk(
		norek,
		norek_induk
	) VALUES(?,?)`)
	if er != nil {
		return routing, errors.New(fmt.Sprint("error while prepare add routing rek induk: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		bankCode,
		norekInduk,
	); er != nil {
		return routing, errors.New(fmt.Sprint("error while add routing rek induk: ", er.Error()))
	} else {
		routing.KodeLkm = bankCode
		routing.NorekInduk = norekInduk
		return routing, nil
	}
}

func (s *SysUserMysqlImpl) UpdateSysApexRoutingRekInduk(newBankCode, norekInduk, currentBankCode string) (routing web.RoutingRekIndukData, er error) {
	stmt, er := s.apexDb.Prepare("UPDATE routing_rek_induk SET norek = ?, norek_induk = ? WHERE norek = ?")
	if er != nil {
		return routing, errors.New(fmt.Sprint("error while prepare update routing rek induk: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(newBankCode, norekInduk, currentBankCode); er != nil {
		return routing, errors.New(fmt.Sprint("error while update routing rek induk: ", er.Error()))
	}

	routing.KodeLkm = newBankCode
	routing.NorekInduk = norekInduk
	return routing, nil
}

func (s *SysUserMysqlImpl) DeleteSysApexRoutingRekInduk(kodeLkm string) (er error) {

	stmt, er := s.apexDb.Prepare("DELETE FROM routing_rek_induk WHERE norek = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete routing rek induk: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete routing rek induk: ", er.Error()))
	}

	return nil
}
