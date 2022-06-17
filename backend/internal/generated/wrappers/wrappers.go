package wrappers

//go:generate gowrap gen -p ../../services/advantage_photo -i AdvantagePhoto -t logrus -o ../../services/advantage_photo/advantage_photo_with_logger.go
//go:generate gowrap gen -p ../../services/advantages -i Advantages -t logrus -o ../../services/advantages/advantages_with_logger.go
//go:generate gowrap gen -p ../../services/auth -i Auth -t logrus -o ../../services/auth/auth_with_logger.go
//go:generate gowrap gen -p ../../services/docs -i Docs -t logrus -o ../../services/docs/docs_with_logger.go
//go:generate gowrap gen -p ../../services/editable -i Editable -t logrus -o ../../services/editable/editable_with_logger.go
//go:generate gowrap gen -p ../../services/hod_photos -i Photos -t logrus -o ../../services/hod_photos/hod_photos_with_logger.go
//go:generate gowrap gen -p ../../services/leads -i Leads -t logrus -o ../../services/leads/leads_with_logger.go
//go:generate gowrap gen -p ../../services/months -i Months -t logrus -o ../../services/months/months_with_logger.go
//go:generate gowrap gen -p ../../services/plans -i Plans -t logrus -o ../../services/plans/plans_with_logger.go
//go:generate gowrap gen -p ../../services/posters -i Posters -t logrus -o ../../services/posters/posters_with_logger.go
//go:generate gowrap gen -p ../../services/project_info -i ProjectInfo -t logrus -o ../../services/project_info/project_info_with_logger.go
//go:generate gowrap gen -p ../../services/users -i Users -t logrus -o ../../services/users/users_with_logger.go
//go:generate gowrap gen -p ../../services/video -i Video -t logrus -o ../../services/video/video_with_logger.go
//go:generate gowrap gen -p ../../services/years -i Years -t logrus -o ../../services/years/years_with_logger.go

//gowrap gen -p ../../services/advantages -i Advantages -t ./cache -o ../../services/advantages/advantages_with_cache.go
//gowrap gen -p ../../services/docs -i Docs -t ./cache -o ../../services/docs/docs_with_cache.go
//gowrap gen -p ../../services/editable -i Editable -t ./cache -o ../../services/editable/editable_with_cache.go
//gowrap gen -p ../../services/plans -i Plans -t ./cache -o ../../services/plans/plans_with_cache.go
//gowrap gen -p ../../services/posters -i Posters -t ./cache -o ../../services/posters/posters_with_cache.go
//gowrap gen -p ../../services/video -i Video -t ./cache -o ../../services/video/video_with_cache.go
//gowrap gen -p ../../services/project_info -i ProjectInfo -t ./cache -o ../../services/project_info/project_info_with_cache.go
