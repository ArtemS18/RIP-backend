package usecase

import (
	"failiverCheck/internal/app/ds"

	log "github.com/sirupsen/logrus"
)

func (uc *UseCase) GetComponent(componetId int) (ds.Component, error) {
	component, err := uc.Postgres.GetComponentById(componetId)
	if err != nil {
		return ds.Component{}, err
	}

	return component, nil
}

func (uc *UseCase) GetComponents(searchQuery string) ([]ds.Component, error) {
	var components []ds.Component
	var err error
	log.Info(searchQuery)
	if searchQuery == "" {
		components, err = uc.Postgres.GetComponents()
		if err != nil {
			return nil, err
		}
	} else {
		components, err = uc.Postgres.GetComponentsByTitle(searchQuery)
		if err != nil {
			return nil, err
		}
	}
	return components, nil
}

// func (uc *UseCase) UpdateComponent(ctx *gin.Context) {
// 	id := uc.getIntParam(ctx, "id")
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	var update dto.UpdateComponentDTO
// 	if err := ctx.BindJSON(&update); err != nil {
// 		uc.errorHandler(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	component, err := uc.Postgres.UpdateComponentById(uint(id), update)
// 	if err != nil {
// 		uc.errorHandler(ctx, http.StatusNotFound, err)
// 		return
// 	}

// 	uc.successHandler(ctx, http.StatusOK, component)
// }

// func (uc *UseCase) CreateComponent(ctx *gin.Context) {
// 	var create dto.CreateComponentDTO
// 	uc.validateFields(ctx, &create)
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	component, err := uc.Postgres.CreateComponent(create)
// 	if err != nil {
// 		uc.errorHandler(ctx, http.StatusInternalServerError, err)
// 		return
// 	}

// 	uc.successHandler(ctx, http.StatusOK, component)
// }

// func (uc *UseCase) DeleteComponent(ctx *gin.Context) {
// 	id := uc.getIntParam(ctx, "id")
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	component, err := uc.Postgres.GetComponentById(id)
// 	if err != nil {
// 		uc.errorHandler(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	imgUrl := component.Img

// 	if err := uc.Postgres.DeletedComponentById(uint(id)); err != nil {
// 		uc.errorHandler(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	if err := uc.Minio.DeleteComponentImg(ctx, &imgUrl); err != nil {
// 		uc.errorHandler(ctx, http.StatusBadRequest, err)
// 		return
// 	}
// 	ctx.Status(http.StatusNoContent)
// }

// func (uc *UseCase) AddComponentInSystemCalc(ctx *gin.Context) {
// 	var err error
// 	var userId uint = uc.GetUserID(ctx)
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	id := uc.getIntParam(ctx, "id")
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	err = uc.Postgres.AddComponentInSystemCalc(uint(id), userId)
// 	if err != nil {
// 		uc.errorHandler(ctx, http.StatusNotFound, err)
// 		return
// 	}
// 	ctx.Status(http.StatusNoContent)
// }

// func (uc *UseCase) UpdateComponentImg(ctx *gin.Context) {
// 	id := uc.getIntParam(ctx, "id")
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	contentType, fileSize := uc.getFileHeaders(ctx)
// 	if ctx.IsAborted() {
// 		return
// 	}
// 	location, err := uc.Minio.UploadComponentImg(ctx, dto.ComponentImgCreateDTO{
// 		File:        ctx.Request.Body,
// 		FilePath:    "img/",
// 		FileSize:    fileSize,
// 		ContentType: contentType,
// 	})
// 	log.Info(location)
// 	if err != nil {
// 		uc.errorHandler(ctx, 404, err)
// 		return
// 	}
// 	component, err := uc.Postgres.GetComponentById(id)
// 	if err != nil {
// 		uc.errorHandler(ctx, 404, err)
// 		return
// 	}
// 	if component.Img != "" {
// 		if err = uc.Minio.DeleteComponentImg(ctx, &component.Img); err != nil {
// 			log.Error(err)
// 		}
// 	}
// 	uc.Postgres.UpdateComponentById(uint(id), dto.UpdateComponentDTO{Img: &location})
// 	uc.successHandler(ctx, 200, map[string]string{"img_url": location})
// }
