BUILD_DIR := builder
OCB := ocb
OCB_CONF := builder_config.yaml
OCB_VER := 0.81.0
OCB_OUT := otelcol-custom

error:
	@echo "Please choose one of the following target: ocb_build, ocb_clean, clean"
	@exit 2

.PHONY: ocb_build
ocb_build: download_ocb
	@echo "building ocb"
	cd $(BUILD_DIR) && ./$(OCB) --config $(OCB_CONF)

.PHONY: ocb_clean
ocb_clean:
	@echo "removing ocb and collector"
	rm -rfv $(OCB_OUT)

download_ocb:
	@if [ -f $(BUILD_DIR)/$(OCB) ]; then \
		echo "ocb file exists"; \
	else \
		echo "downloading ocb file..."; \
		wget "https://github.com/open-telemetry/opentelemetry-collector/releases/download/cmd%2Fbuilder%2Fv$(OCB_VER)/ocb_$(OCB_VER)_linux_amd64" -O $(OCB); \
		chmod +x $(OCB); \
		mv $(OCB) $(BUILD_DIR); \
	fi

.PHONY: clean
clean: ocb_clean
	rm -rfv $(BUILD_DIR)/$(OCB)

