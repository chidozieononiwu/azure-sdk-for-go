package customimagesearchapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/customimagesearch"
)

// CustomInstanceClientAPI contains the set of methods on the CustomInstanceClient type.
type CustomInstanceClientAPI interface {
	ImageSearch(ctx context.Context, customConfig string, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, aspect customimagesearch.ImageAspect, colorParameter customimagesearch.ImageColor, countryCode string, count *int32, freshness customimagesearch.Freshness, height *int32, ID string, imageContent customimagesearch.ImageContent, imageType customimagesearch.ImageType, license customimagesearch.ImageLicense, market string, maxFileSize *int64, maxHeight *int64, maxWidth *int64, minFileSize *int64, minHeight *int64, minWidth *int64, offset *int64, safeSearch customimagesearch.SafeSearch, size customimagesearch.ImageSize, setLang string, width *int32) (result customimagesearch.Images, err error)
}

var _ CustomInstanceClientAPI = (*customimagesearch.CustomInstanceClient)(nil)
