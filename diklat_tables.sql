CREATE TABLE IF NOT EXISTS `train_category` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`code` VARCHAR(20) NOT NULL DEFAULT 'teknis',
	`category_name` VARCHAR(200) NOT NULL,
	`description` VARCHAR(256) DEFAULT ''
) ENGINE=MyISAM;

INSERT IGNORE INTO `train_category` (id, code, category_name, description) VALUES 
	(1, 'teknis', 'Diklat Teknis', 'Kegiatan Diklat Teknis'),
	(2, 'teknis', 'PKP', 'Kegiatan Peningkatan Kapasitas dan Profesionalisme'),
	(3, 'teknis', 'Bimtek', 'Kegiatan Bimbingan Teknis'),
	(4, 'teknis', 'Workshop', 'Kegiatan Workshop'),
	(5, 'teknis', 'Training', 'Kegiatan Training'),
	(6, 'teknis', 'Seminar', 'Kegiatan Seminar'),
	(7, 'fungsional', 'Diklat Fungsional', 'Kegiatan Diklat Fungsional'),
	(8, 'pim', 'Diklat Kepemimpinan Tk. I', 'Kegiatan Kepemimpinan Tingkat I'),
	(9, 'pim', 'Diklat Kepemimpinan Tk. II', 'Kegiatan Kepemimpinan Tingkat II'),
	(10, 'pim', 'Diklat Kepemimpinan Tk. III', 'Kegiatan Kepemimpinan Tingkat III'),
	(11, 'pim', 'Diklat Kepemimpinan Tk. IV', 'Kegiatan Kepemimpinan Tingkat IV'),
	(12, 'pim', 'Diklat Kepemimpinan Tk. V', 'Kegiatan Kepemimpinan V'),
	(13, 'prajabatan', 'Prajabatan', 'Kegiatan Diklat Prajabatan'),
	(14, 'struktural', 'Struktural', 'Kegiatan Diklat Struktural selain Diklatpim'),
	(15, 'undefined', 'Undefined', 'Kegiatan Undefined');
	
CREATE TABLE IF NOT EXISTS `train_penyelenggara` (
	`id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`short_name` VARCHAR(256) NOT NULL UNIQUE,
	`long_name` VARCHAR(512) NOT NULL,
	`alamat` VARCHAR(512) NOT NULL,
	`kota` CHAR(50)
) ENGINE=MyISAM;

INSERT IGNORE INTO `train_penyelenggara` (id, short_name, long_name, alamat, kota) VALUES 
	(1, 'BPSDMD Provinsi Jateng', 'Badan Pengembangan Sumber Daya Manusia Daerah 
	    Provinsi Jawa Tengah', 'Jl. Setiabudi No. 201 Srondol Semarang', 'Semarang'),
	(2, 'BKD Provinsi Jateng', 'Badan Kepegawaian Daerah Provinsi Jawa Tengah', 'Jl. Stadion Selatan No. 1, Semarang', 'Semarang'),
	(3, 'Disarsip Perpus Provinsi Jateng', 'Dinas Kearsipan dan Perpustakaan Provinsi Jawa Tengah', 'Jl. Setiabudi No. 101-103 Semarang', 'Semarang'),
	(4, 'Inspekorat Provinsi Jateng', 'Inspekorat Provinsi Jawa Tengah', 'JL. Pemuda, No. 127-133, Semarang', 'Semarang'),
	(5, 'BKPPD Kabupaten Kebumen', 'Badan Kepegawaian Pendidikan dan Pelatihan Daerah Kabupaten Kebumen',
	    'Jl. Veteran No. 2 Kebumen', 'Kebumen'),
	(6, 'Bapelkes Gombong','Balai Pelatihan Teknis Profesi Kesehatan Provinsi Jawa Tengah Gombong', 'Jl. Wilis No. 1 Wero Gombong', 'Gombong Kebumen'),
	(7, 'Pusdiklat Regional Yogyakarta', 'Pusdiklat Kemendagri Regional Yogakarta',
	    'Jl. Melati Kulon No. 1 Baciro Yogyakarta', 'Yogakarta'),
	(8, 'Balai Diklat VIII Yogyakarta', 'Balai Diklat PU Wilayah III Yogyakarta', 'Jalan Ngeksigondo No. 1-2, Kotagede, Prenggan, Kotagede, Kota Yogyakarta, Daerah Istimewa Yogyakarta', 'train_penyelenggaratrain_categoryYogyakarta'),
	(9, 'Badan Diklat Yogakarta', 'Badan Diklat Daerah Istimewa Yogyakarta', 'Gunungsempu, Desa Tamantirto, Kecamatan Kasihan, Kabupaten Bantul, Daerah Istimewa Yogyakarta', 'Yogyakarta'),
	(10, 'Pusdiklatwas BPKP Bogor', 'Pusdiklatwas BPKP Bogor', 'Jl. Beringin II, Pandansari, Ciawi, Bogor ', 'Bogor'),
	(11, 'Pusdiklat Reskrim POLRI', 'Pusdiklat Reskrim POLRI Megamendung Bogor', 'Megamendung, Bogor, Jawa Barat', 'Bogor'),
	(12, 'PKP2A I LAN Bandung', 'Pusat Kajian Pendidikan dan Pelatihan Aparatur I LAN RI', 'Jl. Kiara Payung km. 4,7 Jatinangor Sumedang, Jawa Barat', 'Sumedang'),
	(13, 'LAN RI', 'Lembaga Administrasi Negara (LAN) RI', 'l. Veteran No.10, Jakarta', 'Jakarta'),
	(14, 'BPSDM Kemendagri RI', 'Badan Pengembangan Sumber Daya Manusia Kemendagri RI', 'Jalan Kompleks Taman Makam Pahlawan No.8, Kalibata, Pancoran, RT.6/RW.4, Duren Tiga, Pancoran, RT.6/RW.4, Duren Tiga, Pancoran, Kota Jakarta Selatan', 'Jakarta'),
	(15, 'Undefined', 'Undefined', 'Undefined', 'Undefined');
	

CREATE TABLE IF NOT EXISTS `train_list` (
	`id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(200) NOT NULL,
	`description` VARCHAR(200) NOT NULL,
	`begin` VARCHAR(50) NOT NULL,
	`end` VARCHAR(50) NOT NULL,
	`location` VARCHAR(100),
	`address` VARCHAR(200),
	`city` VARCHAR(50),
	`jp` INT NOT NULL DEFAULT 0,
	`category` INT NOT NULL,
	`penyelenggara` INT NOT NULL
) ENGINE=MyISAM;


CREATE TABLE IF NOT EXISTS `train_history` (
	`id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`pns` INT NOT NULL,
	`diklat` INT NOT NULL,
	`cert_number` VARCHAR(100),
	`cert_date` DATE
) ENGINE=MyISAM;
	
CREATE OR REPLACE VIEW lof_diklat_view AS 
	SELECT tl.name as diklat, tp.short_name as penyelenggara, COUNT(th.pns) AS peserta, YEAR(tl.begin) as year
	FROM train_history th
	RIGHT JOIN train_list tl
    	ON tl.id = th.diklat
	LEFT JOIN train_penyelenggara tp
    	ON tp.id = tl.penyelenggara
	GROUP BY tl.id
	ORDER BY year ASC;