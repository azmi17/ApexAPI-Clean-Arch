package apexrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/repository/constant"
	"database/sql"
	"errors"
	"fmt"
)

func newApexMysqlImpl(conn1, conn2 *sql.DB) ApexRepo {
	return &ApexMysqlImpl{
		db1: conn1,
		db2: conn2,
	}
}

type ApexMysqlImpl struct {
	db1, db2 *sql.DB
}

func (e *ApexMysqlImpl) CreateNasabah(newNasabah entities.Nasabah) (nasabah entities.Nasabah, er error) {

	stmt, er := e.db1.Prepare(`INSERT INTO nasabah(
		nasabah_id, 
		nama_nasabah, 
		alamat, 
		telpon, 
		jenis_kelamin, 
		tempatlahir, 
		tgllahir, 
		jenis_id, 
		no_id, 
		kode_group1, 
		kode_group2, 
		kode_group3, 
		kode_agama, 
		desa, 
		kecamatan, 
		kota_kab, 
		propinsi, 
		verifikasi, 
		hp, 
		tgl_register, 
		nama_ibu_kandung, 
		kodepos, 
		kode_kantor, 
		userid, 
		nama_alias, 
		status_gelar, 
		jenis_debitur, 
		kode_area, 
		negara_domisili, 
		gol_debitur, 
		langgar_bmpk, 
		lampaui_bmpk, 
		nama_nasabah_sid, 
		alamat2, 
		flag_masa_berlaku,
		status_marital, 
		hp1, 
		hp2, 
		status_tempat_tinggal, 
		masa_berlaku_ktp
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return nasabah, errors.New(fmt.Sprint("error while prepare add nasabah : ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Exec..
	if _, er := stmt.Exec(
		newNasabah.Nasabah_Id,
		newNasabah.Nama_Nasabah,
		newNasabah.Alamat,
		newNasabah.Telepon,
		newNasabah.Jenis_Kelamin,
		newNasabah.TempatLahir,
		newNasabah.TglLahir,
		newNasabah.Jenis_Id,
		newNasabah.No_Id,
		newNasabah.Kode_Group1,
		newNasabah.Kode_Group2,
		newNasabah.Kode_Group3,
		newNasabah.Kode_Agama,
		newNasabah.Desa,
		newNasabah.Kecamatan,
		newNasabah.Kota_Kab,
		newNasabah.Provinsi,
		newNasabah.Verifikasi,
		newNasabah.Hp,
		newNasabah.Tgl_Register,
		newNasabah.Nama_Ibu_Kandung,
		newNasabah.Kodepos,
		newNasabah.Kode_Kantor,
		newNasabah.UserId,
		newNasabah.Nama_Alias,
		newNasabah.Status_Gelar,
		newNasabah.Jenis_Debitur,
		newNasabah.Kode_Area,
		newNasabah.Negara_Domisili,
		newNasabah.Gol_Debitur,
		newNasabah.Langgar_Bmpk,
		newNasabah.Lampaui_Bmpk,
		newNasabah.Nama_Nasabah_Sid,
		newNasabah.Alamat2,
		newNasabah.Flag_Masa_Berlaku,
		newNasabah.Status_Marital,
		newNasabah.Hp1,
		newNasabah.Hp2,
		newNasabah.Status_Tempat_Tinggal,
		newNasabah.Masa_Berlaku_Ktp,
	); er != nil {
		return nasabah, errors.New(fmt.Sprint("error while add nasabah : ", er.Error()))
	} else {
		return newNasabah, nil
	}

}

func (e *ApexMysqlImpl) CreateTabung(newTabung entities.Tabung) (tabung entities.Tabung, er error) {

	stmt, er := e.db1.Prepare(`INSERT INTO tabung(
		no_rekening,
		nasabah_id,
		kode_bi_pemilik,
		suku_bunga,
		persen_pph,
		tgl_register,
		saldo_akhir,
		kode_group1,
		kode_group2,
		verifikasi,
		status,
		kode_kantor,
		kode_integrasi,
		kode_produk,
		userid,
		kode_group3,
		minimum,
		setoran_minimum,
		jkw,
		abp,
		setoran_wajib,
		adm_per_bln,
		target_nominal,
		saldo_akhir_titipan_bunga,
		kode_bi_lokasi,
		saldo_titipan_pokok,
		saldo_titipan_bunga_ks,
		saldo_blokir,
		premi,
		kode_keterkaitan,
		kode_kantor_kas,
		no_rekening_virtual
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return tabung, errors.New(fmt.Sprint("error while prepare add tabung : ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Exec..
	if _, er := stmt.Exec(
		newTabung.No_Rekening,
		newTabung.Nasabah_Id,
		newTabung.Kode_Bi_Pemilik,
		newTabung.Suku_Bunga,
		newTabung.Persen_Pph,
		newTabung.Tgl_Register,
		newTabung.Saldo_Akhir,
		newTabung.Kode_Group1,
		newTabung.Kode_Group2,
		newTabung.Verifikasi,
		newTabung.Status,
		newTabung.Kode_Kantor,
		newTabung.Kode_Integrasi,
		newTabung.Kode_Produk,
		newTabung.UserId,
		newTabung.Kode_Group3,
		newTabung.Minimum,
		newTabung.Setoran_Minimum,
		newTabung.Jkw,
		newTabung.Abp,
		newTabung.Setoran_Wajib,
		newTabung.Adm_Per_Bln,
		newTabung.Target_Nominal,
		newTabung.Saldo_Akhir_Titipan_bunga,
		newTabung.Kode_Bi_Lokasi,
		newTabung.Saldo_Titipan_Pokok,
		newTabung.Saldo_Titipan_Bunga_Ks,
		newTabung.Saldo_Blokir,
		newTabung.Premi,
		newTabung.Kode_Keterkaitan,
		newTabung.Kode_Kantor_Kas,
		newTabung.No_Rekening_Virtual); er != nil {
		return tabung, errors.New(fmt.Sprint("error while add tabung : ", er.Error()))
	} else {
		return newTabung, nil
	}

}

func (e *ApexMysqlImpl) CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	stmt, er := e.db2.Prepare(`INSERT INTO sys_daftar_user(
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

func (e *ApexMysqlImpl) UpdateNasabah(updNasabah entities.Nasabah) (nasabah entities.Nasabah, er error) {

	stmt, er := e.db1.Prepare("UPDATE nasabah SET nama_nasabah = ?, nama_nasabah_sid = ?, nama_alias = ?, alamat = ?, alamat2 = ?, telpon = ?, hp = ?, hp1 = ?,  hp2 = ?, userid = ?  WHERE nasabah_id = ?")
	if er != nil {
		return nasabah, errors.New(fmt.Sprint("error while prepare update nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		updNasabah.Nama_Nasabah,
		updNasabah.Nama_Nasabah_Sid,
		updNasabah.Nama_Alias,
		updNasabah.Alamat,
		updNasabah.Alamat2,
		updNasabah.Telepon,
		updNasabah.Hp,
		updNasabah.Hp1,
		updNasabah.Hp2,
		updNasabah.UserId,
		updNasabah.Nasabah_Id); er != nil {
		return nasabah, errors.New(fmt.Sprint("error while update nasabah : ", er.Error()))
	}

	return updNasabah, nil

}

func (e *ApexMysqlImpl) UpdateSysDaftarUser(updSysuser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	stmt, er := e.db2.Prepare("UPDATE sys_daftar_user SET nama_lengkap = ? WHERE user_name = ?")
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

func (e *ApexMysqlImpl) HardDeleteNasabah(kodeLkm string) (er error) {

	stmt, er := e.db1.Prepare("DELETE FROM nasabah WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete nasabah : ", er.Error()))
	}

	return nil
}

func (e *ApexMysqlImpl) HardDeleteTabung(kodeLkm string) (er error) {

	stmt, er := e.db1.Prepare("DELETE FROM tabung WHERE no_rekening = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabung : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete tabung : ", er.Error()))
	}

	return nil
}

func (e *ApexMysqlImpl) HardDeleteSysDaftarUser(kodeLkm string) (er error) {

	stmt, er := e.db2.Prepare("DELETE FROM sys_daftar_user WHERE user_name = ?")
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

func (e *ApexMysqlImpl) DeleteNasabah(kodeLkm string) (er error) {

	stmt, er := e.db1.Prepare("UPDATE nasabah SET nama_nasabah='XXX' WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete nasabah : ", er.Error()))
	}

	return nil
}

func (e *ApexMysqlImpl) DeleteTabung(kodeLkm string) (er error) {

	stmt, er := e.db1.Prepare("UPDATE tabung SET status = 0 WHERE no_rekening = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabung : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete tabung : ", er.Error()))
	}

	return nil
}

func (e *ApexMysqlImpl) DeleteSysDaftarUser(kodeLkm string) (er error) {

	stmt, er := e.db2.Prepare("UPDATE sys_daftar_user SET status_aktif = 0 WHERE user_name = ?")
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

func (e *ApexMysqlImpl) GetScGroup() (list []web.SCGroup, er error) {
	rows, er := e.db1.Query("SELECT kode_group2, deskripsi_group2 FROM tab_kode_group2")
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var scGroup web.SCGroup
		if er = rows.Scan(&scGroup.KodeGroup, &scGroup.DeskripsiGroup); er != nil {
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

func (e *ApexMysqlImpl) GetLkmDetailInfo(KodeLkm string) (detail web.GetDetailLKMInfo, er error) {

	row := e.db1.QueryRow(`SELECT
		t.no_rekening, 
		g.deskripsi_group2,
		n.nama_nasabah,
		n.alamat,
		n.telpon,
		t.no_rekening,
		t.saldo_akhir,
		t.minimum,
		t.status
	FROM tabung AS t 
	INNER JOIN nasabah AS n ON(t.nasabah_id=n.nasabah_id) 
	INNER JOIN tab_kode_group2 AS g ON(t.kode_group2 = g.kode_group2) 
	WHERE t.status=1 AND t.no_rekening = ?`, KodeLkm)

	er = row.Scan(
		&detail.KodeLembaga,
		&constant.SQLVendor,
		&detail.NamaLembaga,
		&constant.SQLAlamat,
		&constant.SQLKontak,
		&detail.NoRekening,
		&detail.Saldo,
		&constant.SQLPlafond,
		&detail.StatusTab,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return detail, err.NoRecord
		} else {
			return detail, errors.New(fmt.Sprint("error while get instution detail: ", er.Error()))
		}
	}
	detail.Vendor = constant.SQLVendor.String
	detail.Alamat = constant.SQLAlamat.String
	detail.Kontak = constant.SQLKontak.String
	detail.Plafond = constant.SQLPlafond.Float64

	// constant.ConvertSQLDataType()
	return
}

func (e *ApexMysqlImpl) GetLkmInfoList(limitOffset web.LimitOffsetLkmUri) (list []web.GetDetailLKMInfo, er error) {

	args := []interface{}{}
	limit := ""
	if limitOffset.Limit > 0 {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, limitOffset.Limit, limitOffset.Offset)
	} else {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, -1, limitOffset.Offset)
	}
	rows, er := e.db1.Query(`SELECT 
		t.no_rekening, 
		g.deskripsi_group2,
		n.nama_nasabah,
		n.alamat,
		n.telpon,
		t.no_rekening,
		t.saldo_akhir,
		t.minimum,
		t.status
		FROM tabung AS t 
	INNER JOIN nasabah AS n ON(t.nasabah_id=n.nasabah_id) 
	INNER JOIN tab_kode_group2 AS g ON(t.kode_group2 = g.kode_group2) `+limit+``, args...)
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var lkmList web.GetDetailLKMInfo
		if er = rows.Scan(
			&lkmList.KodeLembaga,
			&constant.SQLVendor,
			&lkmList.NamaLembaga,
			&constant.SQLAlamat,
			&constant.SQLKontak,
			&lkmList.NoRekening,
			&lkmList.Saldo,
			&constant.SQLPlafond,
			&lkmList.StatusTab,
		); er != nil {
			return list, er
		}

		lkmList.Vendor = constant.SQLVendor.String
		lkmList.Alamat = constant.SQLAlamat.String
		lkmList.Kontak = constant.SQLKontak.String
		lkmList.Plafond = constant.SQLPlafond.Float64
		list = append(list, lkmList)
	}

	if len(list) == 0 {
		return list, nil // no.record
	} else {
		return
	}
}

func (e *ApexMysqlImpl) ResetApexPassword(user entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {
	stmt, er := e.db2.Prepare("UPDATE sys_daftar_user SET user_web_password = ? WHERE user_name = ?")
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare apex password: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(user.User_Web_Password, user.User_Name); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while update apex password: ", er.Error()))
	}

	return user, nil
}
