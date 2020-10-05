REPO := gcr.io/triggermesh/bridge-examples
TARGETS := slack zendesk

# produce goals to build images for each target ("target1.image target2.image ...")
images := $(foreach t,$(TARGETS),$(t).image)
# produce goals to release images for each target ("target1.image.release target2.image.release ...")
releases := $(foreach i,$(images),$(i).release)
# produce goals to clean images for each target ("target1.image.clean target2.image.clean ...")
imgclean := $(foreach i,$(images),$(i).clean)

.PHONY: images release $(images) $(releases)

images: $(images)
release: $(releases)
clean: $(imgclean)

$(images): %.image:
	$(eval $@_img_name := $(REPO)/oracle-$*)
	@echo "+ Building image $($@_img_name)"
	docker image build --build-arg TARGET=$* -t $($@_img_name) ./

$(releases): %.release: %
	$(eval $@_img_name := $(REPO)/oracle-$(subst .image,,$*))
	@echo "+ Pushing image $($@_img_name)"
	docker image push $($@_img_name)

$(imgclean): %.clean:
	$(eval $@_img_name := $(REPO)/oracle-$(subst .image,,$*))
	@echo "+ Removing image $($@_img_name)"
	docker image rm $($@_img_name) 2>/dev/null || true
